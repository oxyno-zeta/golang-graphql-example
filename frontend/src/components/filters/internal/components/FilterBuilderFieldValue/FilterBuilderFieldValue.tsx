import React, { memo } from 'react';
import MenuItem from '@mui/material/MenuItem';
import Chip from '@mui/material/Chip';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import Autocomplete from '@mui/material/Autocomplete';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import Typography from '@mui/material/Typography';
import { useTranslation } from 'react-i18next';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import dayjs, { Dayjs } from 'dayjs';
import { DateTimePicker } from '@mui/x-date-pickers';
import { FilterOperationMetadataModel, FilterDefinitionEnumObjectModel } from '../../../../../models/general';

/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react/no-array-index-key */
export interface Props {
  readonly id: string;
  readonly value: any;
  readonly onChange: (v: any) => void;
  readonly operation: FilterOperationMetadataModel<any>;
  readonly errorMsg: string | null | undefined;
}

function FilterBuilderFieldValue({ value, onChange, operation, errorMsg, id }: Props) {
  // Setup translate
  const { t } = useTranslation();

  // Check if it is an enum and multiple value is set
  if (operation.enumValues && operation.multipleValues) {
    // Transform value into enum object
    let vEnum: FilterDefinitionEnumObjectModel<any>[] = [];
    // Check if value is set
    if (value) {
      // Loop over value
      vEnum = (value as any[]).reduce((arr, v) => {
        const found = (operation.enumValues as FilterDefinitionEnumObjectModel<any>[]).find((it) => it.value === v);

        if (found) {
          arr.push(found);
        }

        return arr;
      }, []);
    }

    return (
      <Autocomplete<FilterDefinitionEnumObjectModel<any>, true, false, false>
        clearText={t('common.clearAction')}
        closeText={t('common.closeAction')}
        fullWidth
        getOptionLabel={(option) => t(option.display)}
        id={id}
        multiple
        noOptionsText={t('common.filter.noOptions')}
        onChange={(event, newValue) => {
          // Reformat data
          const res = newValue.map((it) => it.value);
          // Save values
          onChange(res);
        }}
        openText={t('common.openAction')}
        options={operation.enumValues}
        renderInput={(params) => (
          <TextField
            {...params}
            error={!!errorMsg}
            helperText={errorMsg ? t(errorMsg) : null}
            label={t('common.filter.value')}
            placeholder={t('common.filter.value')}
            type={operation.inputType}
          />
        )}
        renderOption={(props, data: FilterDefinitionEnumObjectModel<any>, { inputValue }) => {
          const displayedOption = t(data.display);
          const matches = match(displayedOption, inputValue, { insideWords: true, findAllOccurrences: true });
          const parts = parse(displayedOption, matches);

          return (
            <MenuItem {...props}>
              <Box sx={{ display: 'block' }}>
                <Typography>
                  {parts.map((part: any, index: number) => (
                    <span
                      key={index}
                      style={{
                        fontWeight: part.highlight ? 700 : 400,
                      }}
                    >
                      {part.text}
                    </span>
                  ))}
                </Typography>
                {data.description ? (
                  <Typography sx={{ fontStyle: 'italic', overflowWrap: 'break-word', whiteSpace: 'normal' }}>
                    {t(data.description)}
                  </Typography>
                ) : null}
              </Box>
            </MenuItem>
          );
        }}
        renderTags={(tagValue, getTagProps) =>
          tagValue.map((option, index) => <Chip label={t(option.display)} size="small" {...getTagProps({ index })} />)
        }
        size="small"
        value={vEnum}
      />
    );
  }
  // Check if it is an enum
  if (operation.enumValues) {
    // Transform value into enum object
    let vEnum: FilterDefinitionEnumObjectModel<any> | undefined | null = operation.enumValues.find(
      (it) => it.value === value,
    );
    // Check if vEnum is undefined to force it to null
    // This is done to force a controlled component in Autocomplete
    if (vEnum === undefined) {
      vEnum = null;
    }
    return (
      <Autocomplete<FilterDefinitionEnumObjectModel<any>, false, false, false>
        clearText={t('common.clearAction')}
        closeText={t('common.closeAction')}
        fullWidth
        getOptionLabel={(option: FilterDefinitionEnumObjectModel<any> | string) => {
          // Check if option is empty
          if (option === '') {
            return '';
          }

          // Normal case
          return t((option as FilterDefinitionEnumObjectModel<any>).display);
        }}
        id={id}
        noOptionsText={t('common.filter.noOptions')}
        onChange={(input, newValue) => {
          // Check if new value is a string
          if (newValue === null) {
            onChange(null);
            return;
          }

          onChange(newValue.value);
        }}
        openText={t('common.openAction')}
        options={operation.enumValues}
        renderInput={(params) => (
          <TextField
            {...params}
            error={!!errorMsg}
            helperText={errorMsg ? t(errorMsg) : null}
            label={t('common.filter.value')}
            placeholder={t('common.filter.value')}
          />
        )}
        renderOption={(props, data: FilterDefinitionEnumObjectModel<any>, { inputValue }) => {
          const displayedOption = t(data.display);
          const matches = match(displayedOption, inputValue, { insideWords: true, findAllOccurrences: true });
          const parts = parse(displayedOption, matches);

          return (
            <MenuItem {...props}>
              <Box sx={{ display: 'block' }}>
                <Typography>
                  {parts.map((part: any, index: number) => (
                    <span
                      key={index}
                      style={{
                        fontWeight: part.highlight ? 700 : 400,
                      }}
                    >
                      {part.text}
                    </span>
                  ))}
                </Typography>
                {data.description ? (
                  <Typography sx={{ fontStyle: 'italic', overflowWrap: 'break-word', whiteSpace: 'normal' }}>
                    {t(data.description)}
                  </Typography>
                ) : null}
              </Box>
            </MenuItem>
          );
        }}
        size="small"
        value={vEnum}
      />
    );
  }

  // Check if it is a multi value
  if (operation.input && operation.multipleValues) {
    return (
      <Autocomplete
        clearText={t('common.clearAction')}
        closeText={t('common.closeAction')}
        filterOptions={(options, params) => {
          // Open params
          const { inputValue } = params;

          // Check if input value is set
          if (inputValue === '') {
            // Return no option and force no option value and text
            return [];
          }

          // Save value
          let optionValue: string | number = inputValue;
          // Check if input is a number to parse input value
          if (operation.inputType === 'number') {
            optionValue = parseFloat(inputValue);
          }

          // Return new option
          return [
            {
              value: optionValue,
              display: t('common.filter.addOption', { option: inputValue }),
            },
          ];
        }}
        fullWidth
        getOptionLabel={(option) => option.display}
        id={id}
        multiple
        noOptionsText={t('common.filter.noOptions')}
        onChange={(event, newValue) => {
          // Reformat data
          const res = newValue.map((it) => {
            // Check if it is a new option
            if (typeof it === 'object') {
              // Get value from option
              return it.value;
            }

            // Already selected value
            return it;
          });
          // Save values
          onChange(res);
        }}
        openText={t('common.openAction')}
        options={[]}
        renderInput={(params) => (
          <TextField
            {...params}
            error={!!errorMsg}
            helperText={errorMsg ? t(errorMsg) : null}
            label={t('common.filter.value')}
            placeholder={t('common.filter.value')}
            type={operation.inputType}
          />
        )}
        renderTags={(tagValue, getTagProps) =>
          tagValue.map((option, index) => <Chip label={option} size="small" {...getTagProps({ index })} />)
        }
        size="small"
        value={value}
      />
    );
  }

  // Check if it is a simple input
  if (operation.input) {
    // Check if type is a date to include a date picker
    if (operation.inputType === 'date') {
      let val: Dayjs | null = null;
      // Parse if date exists
      if (value !== null && value !== '') {
        val = dayjs(value).tz();
      }

      return (
        <LocalizationProvider dateAdapter={AdapterDayjs} dateLibInstance={dayjs}>
          <DateTimePicker
            ampm={false}
            ampmInClock={false}
            label={t('common.filter.value')}
            localeText={{
              openPreviousView: t('common.date.previousMonthAction'),
              previousMonth: t('common.date.previousMonthAction'),
              openNextView: t('common.date.nextMonthAction'),
              nextMonth: t('common.date.nextMonthAction'),
              toolbarTitle: t('common.date.dateTimePickerToolbarTitle'),
            }}
            onChange={(newValue) => {
              // Check if date is null
              if (newValue === null) {
                onChange(null);
                return;
              }

              // Save formatted to ISO8601 value and it is compatible with RFC3339 due to removal of ns
              // !! Note: This is using the default timezone
              onChange(newValue.tz().format());
            }}
            slotProps={{ textField: { size: 'small', fullWidth: true } }}
            value={val}
            // See formatting here: https://moment.github.io/luxon/#/parsing
            format="YYYY-MM-DD HH:mm:ss"
          />
        </LocalizationProvider>
      );
    }

    return (
      <TextField
        error={!!errorMsg}
        fullWidth
        helperText={errorMsg ? t(errorMsg) : null}
        id={id}
        label={t('common.filter.value')}
        onChange={(event) => {
          onChange(event.target.value);
        }}
        placeholder={t('common.filter.value')}
        size="small"
        type={operation.inputType}
        value={value}
      />
    );
  }

  // Default
  return null;
}
/* eslint-enable @typescript-eslint/no-explicit-any */

export default memo(FilterBuilderFieldValue);

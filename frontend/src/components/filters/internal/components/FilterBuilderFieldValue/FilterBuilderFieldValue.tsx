import React, { memo } from 'react';
import MenuItem from '@mui/material/MenuItem';
import Chip from '@mui/material/Chip';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import Autocomplete from '@mui/material/Autocomplete';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { MobileDateTimePicker } from '@mui/x-date-pickers/MobileDateTimePicker';
import Typography from '@mui/material/Typography';
import { useTranslation } from 'react-i18next';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import dayjs, { Dayjs } from 'dayjs';
import { FilterOperationMetadataModel, FilterDefinitionEnumObjectModel } from '../../../../../models/general';
import { AdapterDayjsTZ } from './AdapterDayjsTZ';

/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react/no-array-index-key */
export interface Props {
  id: string;
  value: any;
  setValue: (v: any) => void;
  operation: FilterOperationMetadataModel<any>;
  errorMsg: string | null | undefined;
}

function FilterBuilderFieldValue({ value, setValue, operation, errorMsg, id }: Props) {
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
        multiple
        fullWidth
        value={vEnum}
        size="small"
        id={id}
        onChange={(event, newValue) => {
          // Reformat data
          const res = newValue.map((it) => it.value);
          // Save values
          setValue(res);
        }}
        options={operation.enumValues}
        getOptionLabel={(option) => t(option.display)}
        renderTags={(tagValue, getTagProps) =>
          tagValue.map((option, index) => <Chip size="small" label={t(option.display)} {...getTagProps({ index })} />)
        }
        renderInput={(params) => (
          <TextField
            {...params}
            label={t('common.filter.value')}
            placeholder={t('common.filter.value')}
            error={!!errorMsg}
            helperText={errorMsg && t(errorMsg)}
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
                {data.description && (
                  <Typography sx={{ fontStyle: 'italic', overflowWrap: 'break-word', whiteSpace: 'normal' }}>
                    {t(data.description)}
                  </Typography>
                )}
              </Box>
            </MenuItem>
          );
        }}
        noOptionsText={t('common.filter.noOptions')}
        openText={t('common.openAction')}
        clearText={t('common.clearAction')}
        closeText={t('common.closeAction')}
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
        fullWidth
        noOptionsText={t('common.filter.noOptions')}
        openText={t('common.openAction')}
        clearText={t('common.clearAction')}
        closeText={t('common.closeAction')}
        size="small"
        id={id}
        value={vEnum}
        options={operation.enumValues}
        getOptionLabel={(option: FilterDefinitionEnumObjectModel<any> | string) => {
          // Check if option is empty
          if (option === '') {
            return '';
          }

          // Normal case
          return t((option as FilterDefinitionEnumObjectModel<any>).display);
        }}
        renderInput={(params) => (
          <TextField
            {...params}
            error={!!errorMsg}
            helperText={errorMsg && t(errorMsg)}
            label={t('common.filter.value')}
            placeholder={t('common.filter.value')}
          />
        )}
        onChange={(input, newValue) => {
          // Check if new value is a string
          if (newValue === null) {
            setValue(null);
            return;
          }

          setValue((newValue as FilterDefinitionEnumObjectModel<any>).value);
        }}
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
                {data.description && (
                  <Typography sx={{ fontStyle: 'italic', overflowWrap: 'break-word', whiteSpace: 'normal' }}>
                    {t(data.description)}
                  </Typography>
                )}
              </Box>
            </MenuItem>
          );
        }}
      />
    );
  }

  // Check if it is a multi value
  if (operation.input && operation.multipleValues) {
    return (
      <Autocomplete
        multiple
        fullWidth
        value={value}
        size="small"
        id={id}
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
          setValue(res);
        }}
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
        options={[]}
        getOptionLabel={(option) => option.display}
        renderTags={(tagValue, getTagProps) =>
          tagValue.map((option, index) => <Chip size="small" label={option} {...getTagProps({ index })} />)
        }
        renderInput={(params) => (
          <TextField
            {...params}
            label={t('common.filter.value')}
            placeholder={t('common.filter.value')}
            error={!!errorMsg}
            helperText={errorMsg && t(errorMsg)}
            type={operation.inputType}
          />
        )}
        noOptionsText={t('common.filter.noOptions')}
        openText={t('common.openAction')}
        clearText={t('common.clearAction')}
        closeText={t('common.closeAction')}
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
        <LocalizationProvider dateAdapter={AdapterDayjsTZ} dateLibInstance={dayjs}>
          <MobileDateTimePicker
            label={t('common.filter.value')}
            value={val}
            onChange={(newValue) => {
              // Check if date is null
              if (newValue === null) {
                setValue(null);
                return;
              }

              // Save formatted to ISO8601 value and it is compatible with RFC3339 due to removal of ns
              // !! Note: This is using the default timezone
              setValue(newValue.tz().format());
            }}
            localeText={{
              openPreviousView: t('common.date.previousMonthAction'),
              previousMonth: t('common.date.previousMonthAction'),
              openNextView: t('common.date.nextMonthAction'),
              nextMonth: t('common.date.nextMonthAction'),
              toolbarTitle: t('common.date.dateTimePickerToolbarTitle'),
            }}
            slotProps={{ textField: { size: 'small', fullWidth: true } }}
          />
        </LocalizationProvider>
      );
    }

    return (
      <TextField
        fullWidth
        size="small"
        id={id}
        type={operation.inputType}
        label={t('common.filter.value')}
        placeholder={t('common.filter.value')}
        error={!!errorMsg}
        helperText={errorMsg && t(errorMsg)}
        value={value}
        onChange={(event) => {
          setValue(event.target.value);
        }}
      />
    );
  }

  // Default
  return null;
}
/* eslint-enable @typescript-eslint/no-explicit-any */

export default memo(FilterBuilderFieldValue);

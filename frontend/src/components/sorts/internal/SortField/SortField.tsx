import React, { memo } from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import MenuItem from '@mui/material/MenuItem';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import { useTranslation } from 'react-i18next';
import { type SortOrderModel, type SortOrderFieldModel } from '../../../../models/general';

const valueOptions: { value: SortOrderModel; display: string }[] = [
  { value: 'ASC', display: 'common.sort.asc' },
  { value: 'DESC', display: 'common.sort.desc' },
];

/* eslint-disable react/no-array-index-key */
export interface Props {
  readonly value: Record<string, SortOrderModel>;
  readonly sortFields: SortOrderFieldModel[];
  readonly availableFields: SortOrderFieldModel[];
  readonly onChange: (v: Record<string, SortOrderModel>) => void;
}

function SortField({ value, sortFields, availableFields, onChange }: Props) {
  // Setup translate
  const { t } = useTranslation();

  // Get key
  const key = Object.keys(value)[0];
  // Get field value
  const fieldValue = value[key];

  return (
    <>
      <Grid
        size={{
          xl: 6,
          lg: 6,
          md: 6,
          sm: 6,
          xs: 12,
        }}
      >
        <Autocomplete
          closeText={t('common.closeAction')}
          disableClearable
          fullWidth
          getOptionLabel={(option) => t(option.display)}
          isOptionEqualToValue={(a, b) => a.field === b.field}
          noOptionsText={t('common.filter.noOptions')}
          onChange={(input, newValue) => {
            onChange({ [newValue.field]: fieldValue });
          }}
          openText={t('common.openAction')}
          options={availableFields}
          renderInput={(params) => <TextField {...params} />}
          renderOption={(props, option, { inputValue }) => {
            const displayedOption = t(option.display);
            const matches = match(displayedOption, inputValue, { insideWords: true, findAllOccurrences: true });
            const parts = parse(displayedOption, matches);

            return (
              <MenuItem {...props}>
                <Typography>
                  {parts.map((part: { highlight: boolean; text: string }, index: number) => (
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
              </MenuItem>
            );
          }}
          size="small"
          value={sortFields.find((it) => it.field === key)}
        />
      </Grid>
      <Grid
        size={{
          xl: 6,
          lg: 6,
          md: 6,
          sm: 6,
          xs: 12,
        }}
      >
        <Autocomplete
          closeText={t('common.closeAction')}
          disableClearable
          fullWidth
          getOptionLabel={(option) => t(option.display)}
          noOptionsText={t('common.filter.noOptions')}
          onChange={(input, newValue) => {
            onChange({ [key]: newValue.value });
          }}
          openText={t('common.openAction')}
          options={valueOptions}
          renderInput={(params) => <TextField {...params} />}
          renderOption={(props, option, { inputValue }) => {
            const displayedOption = t(option.display);
            const matches = match(displayedOption, inputValue, { insideWords: true, findAllOccurrences: true });
            const parts = parse(displayedOption, matches);

            return (
              <MenuItem {...props}>
                <Typography>
                  {parts.map((part: { highlight: boolean; text: string }, index: number) => (
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
              </MenuItem>
            );
          }}
          size="small"
          value={valueOptions.find((it) => it.value === fieldValue)}
        />
      </Grid>
    </>
  );
}

export default memo(SortField);

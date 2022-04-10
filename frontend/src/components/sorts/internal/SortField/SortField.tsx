import React, { memo } from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete from '@mui/material/Autocomplete';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import MenuItem from '@mui/material/MenuItem';
import Typography from '@mui/material/Typography';
import { useTranslation } from 'react-i18next';
import { SortOrderModel, SortOrderFieldModel } from '../../../../models/general';

const valueOptions: { value: SortOrderModel; display: string }[] = [
  { value: 'ASC', display: 'common.sort.asc' },
  { value: 'DESC', display: 'common.sort.desc' },
  { value: undefined, display: 'common.sort.undefined' },
];

/* eslint-disable react/no-array-index-key */
interface Props {
  value: SortOrderModel;
  onChange: (v: SortOrderModel) => void;
  fieldDeclaration: SortOrderFieldModel;
}

function SortField({ value, fieldDeclaration, onChange }: Props) {
  // Setup translate
  const { t } = useTranslation();

  return (
    <Autocomplete
      fullWidth
      noOptionsText={t('common.filter.noOptions')}
      openText={t('common.openAction')}
      closeText={t('common.closeAction')}
      disableClearable
      size="small"
      value={valueOptions.find((it) => it.value === value)}
      options={valueOptions}
      renderInput={(params) => (
        <TextField {...params} label={t(fieldDeclaration.display)} placeholder={t(fieldDeclaration.display)} />
      )}
      onChange={(input, newValue) => {
        onChange(newValue ? newValue.value : undefined);
      }}
      getOptionLabel={(option) => t(option.display)}
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
    />
  );
}

export default memo(SortField);

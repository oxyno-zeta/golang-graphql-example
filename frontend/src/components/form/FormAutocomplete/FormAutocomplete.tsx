import React from 'react';
import { useTranslation } from 'react-i18next';
import Autocomplete, { type AutocompleteProps } from '@mui/material/Autocomplete';
import MenuItem from '@mui/material/MenuItem';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import TextField, { type TextFieldProps } from '@mui/material/TextField';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import { type Control, useController, type Path, type FieldValues } from 'react-hook-form';
import type { YupTranslateErrorModel } from '../../../models/general';

interface ValueModel {
  display: string;
  value: string;
}

interface Props<T extends FieldValues> {
  readonly control: Control<T>;
  readonly name: Path<T>;
  readonly values: ValueModel[];
  readonly autocompleteProps?: Partial<AutocompleteProps<ValueModel, false, false, false>>;
  readonly textFieldProps?: Partial<TextFieldProps>;
}

/* eslint-disable react/no-array-index-key */
function FormAutocomplete<T extends FieldValues>({
  control,
  name,
  values,
  autocompleteProps = {},
  textFieldProps = {},
}: Props<T>) {
  // Setup translate
  const { t } = useTranslation();
  // Use controller
  const { field, fieldState } = useController({ control, name });

  const errorProps: Partial<TextFieldProps> = {};
  // Check if error is set
  if (fieldState.error) {
    // Set error boolean
    errorProps.error = !!fieldState.error;

    // Check type of error message
    if (typeof fieldState.error.message === 'object') {
      const mess = fieldState.error.message as YupTranslateErrorModel;

      // Add helper text
      errorProps.helperText = t(mess.key, mess.values) as string;
    } else if (typeof fieldState.error.message === 'string') {
      // Add helper text
      errorProps.helperText = t(fieldState.error.message);
    } else if (Array.isArray(fieldState.error)) {
      errorProps.helperText = fieldState.error.reduce((acc, v) => {
        const mess = v.message as YupTranslateErrorModel;

        return `${acc} ${t(mess.key, mess.values)}`;
      }, '');
    }
  }

  // Transform value into value object
  // eslint-disable-next-line react-hooks/refs
  let value: ValueModel | undefined | null = values.find((it) => it.value === field.value);
  // Check if value is undefined to force it to null
  // This is done to force a controlled component in Autocomplete
  if (value === undefined) {
    value = null;
  }

  return (
    <Autocomplete<ValueModel>
      clearText={t('common.clearAction')}
      closeText={t('common.closeAction')}
      getOptionLabel={(option: ValueModel | string) => {
        // Check if option is empty
        if (option === '') {
          return '';
        }

        // Normal case
        return (option as ValueModel).display;
      }}
      noOptionsText={t('common.filter.noOptions')}
      // eslint-disable-next-line react-hooks/refs
      onBlur={field.onBlur}
      onChange={(input, newValue) => {
        if (newValue === null) {
          field.onChange(null);

          return;
        }

        field.onChange(newValue.value);
      }}
      openText={t('common.openAction')}
      options={values}
      // eslint-disable-next-line react-hooks/refs
      ref={field.ref}
      renderInput={(params) => <TextField {...params} {...textFieldProps} {...errorProps} />}
      renderOption={(props, data: ValueModel, { inputValue }) => {
        const displayedOption = t(data.display);
        const matches = match(displayedOption, inputValue, { insideWords: true, findAllOccurrences: true });
        const parts = parse(displayedOption, matches);

        return (
          <MenuItem {...props}>
            <Box sx={{ display: 'block' }}>
              <Typography>
                {parts.map((part: { text: string; highlight: boolean }, index: number) => (
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
            </Box>
          </MenuItem>
        );
      }}
      value={value}
      {...autocompleteProps}
    />
  );
}

export default FormAutocomplete;

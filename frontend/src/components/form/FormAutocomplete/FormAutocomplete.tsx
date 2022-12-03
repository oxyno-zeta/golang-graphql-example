import React from 'react';
import { useTranslation } from 'react-i18next';
import Autocomplete, { AutocompleteProps } from '@mui/material/Autocomplete';
import MenuItem from '@mui/material/MenuItem';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import TextField, { TextFieldProps } from '@mui/material/TextField';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import { Control, useController, Path, FieldValues } from 'react-hook-form';
import type { YupTranslateErrorModel } from '../../../models/general';

type ValueModel = { display: string; value: string };

type Props<T extends FieldValues> = {
  control: Control<T>;
  name: Path<T>;
  values: ValueModel[];
  autocompleteProps?: Partial<AutocompleteProps<ValueModel, false, false, false>>;
  textFieldProps?: Partial<TextFieldProps>;
};

/* eslint-disable react/no-array-index-key */
function FormAutocomplete<T extends FieldValues>({
  control,
  name,
  values,
  autocompleteProps,
  textFieldProps,
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

    // Check type <of error message
    if (typeof fieldState.error.message === 'object') {
      const mess = fieldState.error.message as YupTranslateErrorModel;

      // Add helper text
      errorProps.helperText = <>{t(mess.key, mess.values)}</>;
    } else if (typeof fieldState.error.message === 'string') {
      // Add helper text
      errorProps.helperText = t(fieldState.error.message);
    }
  }

  // Transform value into value object
  let value: ValueModel | undefined | null = values.find((it) => it.value === field.value);
  // Check if value is undefined to force it to null
  // This is done to force a controlled component in Autocomplete
  if (value === undefined) {
    value = null;
  }

  return (
    <Autocomplete<ValueModel, false, false, false>
      noOptionsText={t('common.filter.noOptions')}
      openText={t('common.openAction')}
      clearText={t('common.clearAction')}
      closeText={t('common.closeAction')}
      onBlur={field.onBlur}
      value={value}
      ref={field.ref}
      options={values}
      getOptionLabel={(option: ValueModel | string) => {
        // Check if option is empty
        if (option === '') {
          return '';
        }

        // Normal case
        return (option as ValueModel).display;
      }}
      renderInput={(params) => <TextField {...params} {...textFieldProps} {...errorProps} />}
      onChange={(input, newValue) => {
        // field.onChange();
        if (newValue === null) {
          field.onChange(null);

          return;
        }

        field.onChange((newValue as ValueModel).value);
      }}
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
      {...autocompleteProps}
    />
  );
}

FormAutocomplete.defaultProps = {
  autocompleteProps: {},
  textFieldProps: {},
};

export default FormAutocomplete;

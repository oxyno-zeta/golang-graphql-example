import React from 'react';
import { useTranslation } from 'react-i18next';
import TextField, { TextFieldProps } from '@mui/material/TextField';
import { Control, useController, Path, FieldValues } from 'react-hook-form';
import type { YupTranslateErrorModel } from '../../../models/general';

type Props<T extends FieldValues> = {
  readonly control: Control<T>;
  readonly name: Path<T>;
  readonly textFieldProps?: TextFieldProps;
};

function FormInput<T extends FieldValues>({ control, name, textFieldProps = {} }: Props<T>) {
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

        return `${acc} ${t(mess.key, mess.values)}` as string;
      }, '');
    }
  }

  return <TextField {...errorProps} {...field} {...textFieldProps} />;
}

export default FormInput;

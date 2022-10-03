import React from 'react';
import { useTranslation } from 'react-i18next';
import TextField, { TextFieldProps } from '@mui/material/TextField';
import { Control, useController, Path, FieldValues } from 'react-hook-form';
import { YupTranslateErrorModel } from '../../../models/general';

type Props<T extends FieldValues> = {
  control: Control<T>;
  name: Path<T>;
  textFieldProps?: TextFieldProps;
};

function FormInput<T extends FieldValues>({ control, name, textFieldProps }: Props<T>) {
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
      errorProps.helperText = t(mess.key, mess.values);
    } else if (typeof fieldState.error.message === 'string') {
      // Add helper text
      errorProps.helperText = t(fieldState.error.message);
    }
  }

  return <TextField {...errorProps} {...field} {...textFieldProps} />;
}

FormInput.defaultProps = {
  textFieldProps: {},
};

export default FormInput;

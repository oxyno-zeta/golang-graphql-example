import React, { useEffect, useState } from 'react';
import MenuItem from '@mui/material/MenuItem';
import Autocomplete from '@mui/material/Autocomplete';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid2';
import Typography from '@mui/material/Typography';
import { useTranslation } from 'react-i18next';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import type {
  FilterDefinitionFieldsModel,
  FilterDefinitionFieldObjectMetadataModel,
  FilterOperationMetadataModel,
} from '../../../../../models/general';
import FilterBuilderFieldValue from '../FilterBuilderFieldValue';
import { requiredInputValidate } from '../../utils';
import type { FieldInitialValueObject, FieldOperationValueObject, FilterValueObject } from '../../types';

/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react/no-array-index-key */
export interface Props {
  id: string;
  filterDefinitionModel: FilterDefinitionFieldsModel;
  initialValue: FieldInitialValueObject;
  onChange: (fo: null | FilterValueObject) => void;
}

function FilterBuilderField({ filterDefinitionModel, onChange, initialValue, id }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // States
  const [selectedField, setSelectedField] = useState(initialValue.field);
  const [selectedOperation, setSelectedOperation] = useState(initialValue.operation);
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const [value, setValue] = useState<any>(initialValue.value);
  // Data
  const fieldKeys = Object.keys(filterDefinitionModel);
  const selectedFieldData = selectedField !== '' ? filterDefinitionModel[selectedField] : null;
  const operations = selectedFieldData ? Object.keys(selectedFieldData.operations) : [];
  const operationData = selectedFieldData
    ? (selectedFieldData as FilterDefinitionFieldObjectMetadataModel<any>).operations[selectedOperation]
    : null;
  // Validation
  const fieldErrorMsg = requiredInputValidate(selectedField);
  const operationErrorMsg = requiredInputValidate(selectedOperation);
  let valueErrorMsg: string | null | undefined = null;
  // Check if operationData is set
  if (operationData) {
    valueErrorMsg = operationData.inputValidate ? operationData.inputValidate(value) : null;
  }

  // Watch initialValue
  useEffect(() => {
    // Set initial value data
    setSelectedField(initialValue.field);
    setSelectedOperation(initialValue.operation);
    setValue(initialValue.value);
  }, [initialValue]);

  // UseEffect to build line filter
  useEffect(() => {
    // Check if there is an error
    if (fieldErrorMsg || operationErrorMsg || valueErrorMsg) {
      onChange(null);
      return;
    }

    // Build inner object
    const innerObj: FieldOperationValueObject = { [selectedOperation]: value };
    if (operationData?.caseInsensitiveEnabled) {
      innerObj.caseInsensitive = true;
    }
    // Save filter object
    onChange({ [selectedField]: innerObj });
  }, [
    fieldErrorMsg,
    onChange,
    operationData?.caseInsensitiveEnabled,
    operationErrorMsg,
    selectedField,
    selectedOperation,
    value,
    valueErrorMsg,
  ]);

  return (
    <>
      <Grid
        sx={{ display: 'flex' }}
        size={{
          xl: 4,
          lg: 4,
          md: 6,
          sm: 6,
          xs: 12,
        }}
      >
        <Autocomplete
          fullWidth
          noOptionsText={t('common.filter.noOptions')}
          openText={t('common.openAction')}
          clearText={t('common.clearAction')}
          closeText={t('common.closeAction')}
          size="small"
          id={`${id}-field`}
          value={t(filterDefinitionModel[selectedField]?.display)}
          options={fieldKeys}
          renderInput={(params) => (
            <TextField
              {...params}
              error={!!fieldErrorMsg}
              helperText={fieldErrorMsg && t(fieldErrorMsg)}
              label={t('common.filter.field')}
              placeholder={t('common.filter.field')}
            />
          )}
          isOptionEqualToValue={(option, v) => {
            // Check if value is set
            if (v === '') {
              return false;
            }

            // Get option
            const optDisplay = filterDefinitionModel[option].display;

            // Check if displayed option is the same as value
            return t(optDisplay) === v;
          }}
          onChange={(input, newValue) => {
            // Check if value isn't the same as actual
            if (newValue !== selectedField) {
              // Set new value
              setSelectedField(newValue || '');
              // Force flush operation
              setSelectedOperation('');
              // Flush value
              setValue(undefined);
            }
          }}
          renderOption={(props, fieldKey: string, { inputValue }) => {
            const fieldData = filterDefinitionModel[fieldKey];
            const displayedOption = t(fieldData.display);
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
                  {fieldData.description && (
                    <Typography sx={{ fontStyle: 'italic', overflowWrap: 'break-word', whiteSpace: 'normal' }}>
                      {t(fieldData.description)}
                    </Typography>
                  )}
                </Box>
              </MenuItem>
            );
          }}
        />
      </Grid>
      <Grid
        size={{
          xl: 4,
          lg: 4,
          md: 6,
          sm: 6,
          xs: 12,
        }}
      >
        {selectedFieldData && (
          <Autocomplete
            fullWidth
            noOptionsText={t('common.filter.noOptions')}
            openText={t('common.openAction')}
            clearText={t('common.clearAction')}
            closeText={t('common.closeAction')}
            size="small"
            id={`${id}-operation`}
            value={t(selectedFieldData.operations[selectedOperation]?.display)}
            options={operations}
            getOptionLabel={(option: string) => t(option)}
            renderInput={(params) => (
              <TextField
                {...params}
                error={!!operationErrorMsg}
                helperText={operationErrorMsg && t(operationErrorMsg)}
                label={t('common.filter.operation')}
                placeholder={t('common.filter.operation')}
              />
            )}
            isOptionEqualToValue={(option, v) => {
              // Check if value is set
              if (v === '') {
                return false;
              }

              // Get option
              const optDisplay = (selectedFieldData as FilterDefinitionFieldObjectMetadataModel<any>).operations[option]
                .display;

              // Check if displayed option is the same as value
              return t(optDisplay) === v;
            }}
            onChange={(input, newValue) => {
              // Check if value isn't the same as actual
              if (newValue !== selectedOperation) {
                // Set new value
                setSelectedOperation(newValue || '');

                if (newValue !== null) {
                  // Reset value
                  setValue(
                    (selectedFieldData as FilterDefinitionFieldObjectMetadataModel<any>).operations[newValue]
                      .initialValue,
                  );
                }
              }
            }}
            renderOption={(props, key: string, { inputValue }) => {
              const opData = (selectedFieldData as FilterDefinitionFieldObjectMetadataModel<any>).operations[key];
              const displayedOption = t(opData.display);
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
                    {opData.description && (
                      <Typography sx={{ fontStyle: 'italic', overflowWrap: 'break-word', whiteSpace: 'normal' }}>
                        {t(opData.description)}
                      </Typography>
                    )}
                  </Box>
                </MenuItem>
              );
            }}
          />
        )}
      </Grid>
      <Grid
        size={{
          xl: 4,
          lg: 4,
          md: 6,
          sm: 6,
          xs: 12,
        }}
      >
        {operationData && (
          <FilterBuilderFieldValue
            value={value}
            id={`${id}-value`}
            onChange={setValue}
            operation={operationData as FilterOperationMetadataModel<any>}
            errorMsg={valueErrorMsg}
          />
        )}
      </Grid>
    </>
  );
}
/* eslint-enable @typescript-eslint/no-explicit-any */

export default FilterBuilderField;

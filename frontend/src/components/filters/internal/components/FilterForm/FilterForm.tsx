import React, { useEffect, useState, memo } from 'react';
import MenuItem from '@mui/material/MenuItem';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import { useTranslation } from 'react-i18next';
import Autocomplete from '@mui/material/Autocomplete';
import TextField from '@mui/material/TextField';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import { FilterDefinitionFieldsModel } from '../../../../../models/general';
import FilterBuilder from '../FilterBuilder';
import { buildFilterBuilderInitialItems } from '../../utils';
import { BuilderInitialValueObject, FilterValueObject, PredefinedFilter } from '../../types';

/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react/no-array-index-key */
export interface Props {
  filterDefinitionModel: FilterDefinitionFieldsModel;
  onChange: (filter: FilterValueObject | null) => void;
  predefinedFilterObjects?: PredefinedFilter[];
  initialFilter?: undefined | null | FilterValueObject;
}

function FilterForm({ filterDefinitionModel, predefinedFilterObjects, initialFilter, onChange }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // State
  // Create counter state to force refresh
  const [init, setInit] = useState<BuilderInitialValueObject>(buildFilterBuilderInitialItems(initialFilter));
  const [predefinedFilter, setPredefinedFilter] = useState<PredefinedFilter | null>(null);

  // Watch initialFilter
  useEffect(() => {
    // Build new value
    const nV = buildFilterBuilderInitialItems(initialFilter);
    // Check if objects are different
    if (JSON.stringify(nV) !== JSON.stringify(init)) {
      // Set init
      setInit(nV);
    }
  }, [initialFilter]);

  return (
    <>
      {predefinedFilterObjects && (
        <Box sx={{ display: 'flex', margin: '7px 0' }}>
          <Autocomplete
            fullWidth
            freeSolo
            id="predefined-filters"
            sx={{ maxWidth: 300 }}
            noOptionsText={t('common.filter.noOptions')}
            openText={t('common.openAction')}
            clearText={t('common.clearAction')}
            closeText={t('common.closeAction')}
            size="small"
            value={predefinedFilter}
            options={predefinedFilterObjects}
            getOptionLabel={(option: PredefinedFilter | string) => {
              // Handle empty case
              if (option === '') {
                return '';
              }

              // Normal case
              return t((option as PredefinedFilter).display);
            }}
            renderInput={(params) => (
              <TextField
                {...params}
                label={t('common.filter.selectPredefinedFilter')}
                placeholder={t('common.filter.selectPredefinedFilter')}
              />
            )}
            onChange={(input, newValue) => {
              // Handle empty case
              if (newValue === '') {
                setPredefinedFilter(null);
              }
              // Normal case
              setPredefinedFilter(newValue as PredefinedFilter);
            }}
            renderOption={(props, option: PredefinedFilter, { inputValue }) => {
              const displayedOption = t(option.display);
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
                    {option.description && (
                      <Typography sx={{ fontStyle: 'italic', overflowWrap: 'break-word', whiteSpace: 'normal' }}>
                        {t(option.description)}
                      </Typography>
                    )}
                  </Box>
                </MenuItem>
              );
            }}
          />
          <Button
            size="small"
            sx={{ marginLeft: '5px' }}
            disabled={predefinedFilter === null}
            onClick={() => {
              // Load new init object
              setInit(buildFilterBuilderInitialItems(predefinedFilter?.filter));
            }}
          >
            {t('common.loadAction')}
          </Button>
        </Box>
      )}
      <FilterBuilder
        filterDefinitionModel={filterDefinitionModel}
        onChange={onChange}
        initialValue={init}
        acceptEmptyLines
      />
    </>
  );
}
/* eslint-enable react/no-array-index-key */

FilterForm.defaultProps = {
  predefinedFilterObjects: undefined,
  initialFilter: undefined,
};

export default memo(FilterForm);

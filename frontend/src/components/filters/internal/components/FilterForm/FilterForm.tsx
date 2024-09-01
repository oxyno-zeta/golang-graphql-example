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
  readonly filterDefinitionModel: FilterDefinitionFieldsModel;
  readonly onChange: (filter: FilterValueObject | null) => void;
  readonly predefinedFilterObjects?: PredefinedFilter[];
  readonly initialFilter?: undefined | null | FilterValueObject;
}

function FilterForm({
  filterDefinitionModel,
  predefinedFilterObjects = undefined,
  initialFilter = undefined,
  onChange,
}: Props) {
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
    // Set init
    setInit((innerInit) => {
      // Check if objects are different
      if (JSON.stringify(nV) !== JSON.stringify(innerInit)) {
        return nV;
      }

      // Keep like that
      return innerInit;
    });
  }, [initialFilter]);

  return (
    <>
      {predefinedFilterObjects ? (
        <Box sx={{ display: 'flex', margin: '7px 0' }}>
          <Autocomplete
            clearText={t('common.clearAction')}
            closeText={t('common.closeAction')}
            freeSolo
            fullWidth
            getOptionLabel={(option: PredefinedFilter | string) => {
              // Handle empty case
              if (option === '') {
                return '';
              }

              // Normal case
              return t((option as PredefinedFilter).display);
            }}
            id="predefined-filters"
            noOptionsText={t('common.filter.noOptions')}
            onChange={(input, newValue) => {
              // Handle empty case
              if (newValue === '') {
                setPredefinedFilter(null);
              }
              // Normal case
              setPredefinedFilter(newValue as PredefinedFilter);
            }}
            openText={t('common.openAction')}
            options={predefinedFilterObjects}
            renderInput={(params) => (
              <TextField
                {...params}
                label={t('common.filter.selectPredefinedFilter')}
                placeholder={t('common.filter.selectPredefinedFilter')}
              />
            )}
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
                    {option.description ? (
                      <Typography sx={{ fontStyle: 'italic', overflowWrap: 'break-word', whiteSpace: 'normal' }}>
                        {t(option.description)}
                      </Typography>
                    ) : null}
                  </Box>
                </MenuItem>
              );
            }}
            size="small"
            sx={{ maxWidth: 300 }}
            value={predefinedFilter}
          />
          <Button
            disabled={predefinedFilter === null}
            onClick={() => {
              // Load new init object
              setInit(buildFilterBuilderInitialItems(predefinedFilter?.filter));
            }}
            size="small"
            sx={{ marginLeft: '5px' }}
          >
            {t('common.loadAction')}
          </Button>
        </Box>
      ) : null}
      <FilterBuilder
        acceptEmptyLines
        filterDefinitionModel={filterDefinitionModel}
        initialValue={init}
        onChange={onChange}
      />
    </>
  );
}
/* eslint-enable react/no-array-index-key */

export default memo(FilterForm);

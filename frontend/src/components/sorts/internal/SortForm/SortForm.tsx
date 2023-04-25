import React, { useState } from 'react';
import IconButton from '@mui/material/IconButton';
import Box from '@mui/material/Box';
import { mdiPlus, mdiDelete, mdiChevronDown, mdiChevronUp } from '@mdi/js';
import SvgIcon from '@mui/material/SvgIcon';
import Button from '@mui/material/Button';
import ButtonGroup from '@mui/material/ButtonGroup';
import Tooltip from '@mui/material/Tooltip';
import DialogActions from '@mui/material/DialogActions';
import Grid from '@mui/material/Grid';
import DialogContent from '@mui/material/DialogContent';
import { useTranslation } from 'react-i18next';
import SortField from '../SortField';
import { SortOrderModel, SortOrderFieldModel, SortOrderAsc } from '../../../../models/general';
import { arrayMoveItem } from '../../../../utils/array';

export interface Props<T extends Record<string, SortOrderModel>> {
  onSubmit: (sort: T[]) => void;
  onReset: () => void;
  initialSorts: null | undefined | T[];
  sortFields: SortOrderFieldModel[];
}

/* eslint-disable react/no-array-index-key */
function SortForm<T extends Record<string, SortOrderModel>>({ onSubmit, onReset, initialSorts, sortFields }: Props<T>) {
  // Setup translate
  const { t } = useTranslation();
  // State
  const [result, setResult] = useState<T[]>(initialSorts || []);

  // Flatten all keys in sorts
  const sortsKeys = result.reduce((accu, sortItem) => {
    // Get keys of sort item
    const keys = Object.keys(sortItem);
    // Check if keys are found
    if (!keys || keys.length === 0) {
      // No keys detected
      // Ignoring case
      return accu;
    }

    // As the api is managing only 1 key, only take the first one
    accu.push(keys[0]);

    return accu;
  }, [] as string[]);

  // Compute available fields
  const availableFields = sortFields.filter((value) => !sortsKeys.find((key) => key === value.field));

  // Add new line handler
  const onAddLine = () => {
    const tmp = [...result];
    // Add new item with first available field
    tmp.push({ [availableFields[0].field]: SortOrderAsc } as T);
    // Save
    setResult(tmp);
  };

  return (
    <>
      <DialogContent>
        {result.map((item, index) => {
          // Build id
          const id = `sort-${index}-${Object.keys(item)[0]}`;

          return (
            <Box key={id} data-testid={id} sx={{ display: 'flex', margin: '10px 0 20px 0' }}>
              <Box sx={{ display: 'flex', alignItems: 'center', marginRight: '5px' }}>
                {result.length > 1 && (
                  <ButtonGroup orientation="vertical">
                    {index !== 0 && (
                      <Tooltip title={<>{t('common.upAction')}</>}>
                        <IconButton
                          size="small"
                          sx={{ height: '24px', width: '24px' }}
                          onClick={() => {
                            // Move
                            arrayMoveItem(result, index, index - 1);
                            // Save
                            setResult([...result]);
                          }}
                        >
                          <SvgIcon sx={{ height: '16px', width: '16px' }}>
                            <path d={mdiChevronUp} />
                          </SvgIcon>
                        </IconButton>
                      </Tooltip>
                    )}
                    {index !== result.length - 1 && (
                      <Tooltip title={<>{t('common.downAction')}</>}>
                        <IconButton
                          size="small"
                          sx={{ height: '24px', width: '24px' }}
                          onClick={() => {
                            // Move
                            arrayMoveItem(result, index, index + 1);
                            // Save
                            setResult([...result]);
                          }}
                        >
                          <SvgIcon sx={{ height: '16px', width: '16px' }}>
                            <path d={mdiChevronDown} />
                          </SvgIcon>
                        </IconButton>
                      </Tooltip>
                    )}
                  </ButtonGroup>
                )}
                <Tooltip title={<>{t('common.sort.deleteField')}</>}>
                  <IconButton
                    onClick={() => {
                      const res = [...result];
                      // Delete item
                      res.splice(index, 1);
                      // Save
                      setResult(res);
                    }}
                  >
                    <SvgIcon>
                      <path d={mdiDelete} />
                    </SvgIcon>
                  </IconButton>
                </Tooltip>
              </Box>
              <Grid container spacing={1}>
                <SortField
                  availableFields={availableFields}
                  sortFields={sortFields}
                  value={item}
                  onChange={(v) => {
                    const res = [...result];

                    // Change item at index
                    res.splice(index, 1, v as T);

                    // Save
                    setResult(res);
                  }}
                />
              </Grid>
            </Box>
          );
        })}
        {availableFields.length !== 0 && (
          <Tooltip title={<>{t('common.sort.addNewField')}</>}>
            <IconButton onClick={onAddLine} sx={{ margin: '0 5px' }}>
              <SvgIcon>
                <path d={mdiPlus} />
              </SvgIcon>
            </IconButton>
          </Tooltip>
        )}
      </DialogContent>
      <DialogActions>
        <Button
          onClick={() => {
            onReset();
          }}
          sx={{ marginLeft: 'auto', marginRight: '5px' }}
        >
          {t('common.resetAction')}
        </Button>
        <Button
          variant="contained"
          onClick={() => {
            onSubmit(result as T[]);
          }}
          autoFocus
        >
          {t('common.applyAction')}
        </Button>
      </DialogActions>
    </>
  );
}

export default SortForm;

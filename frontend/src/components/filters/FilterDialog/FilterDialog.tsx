import React, { useCallback, useEffect, useRef, useState } from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { useTranslation } from 'react-i18next';
import FilterForm from '../internal/components/FilterForm';
import { FilterDefinitionFieldsModel } from '../../../models/general';
import { FilterValueObject, PredefinedFilter } from '../internal/types';

/* eslint-disable @typescript-eslint/no-explicit-any */
export type Props<T extends FilterValueObject> = {
  onSubmit: (filter: T) => void;
  onReset: () => void;
  onClose: () => void;
  open: boolean;
  filterDefinitionModel: FilterDefinitionFieldsModel;
  predefinedFilterObjects?: PredefinedFilter[];
  initialFilter?: undefined | null | T;
};

function FilterDialog<T extends FilterValueObject>({
  filterDefinitionModel,
  predefinedFilterObjects = undefined,
  onSubmit,
  onReset,
  onClose,
  initialFilter = undefined,
  open,
}: Props<T>) {
  // Setup translate
  const { t } = useTranslation();
  // State
  const resultRef = useRef<T | undefined | null>(null);
  const [error, setError] = useState(false);

  const localOnChange = useCallback((filter: null | FilterValueObject) => {
    resultRef.current = filter as T | null;

    if (filter === null || filter === undefined) {
      setError(true);
    } else {
      setError(false);
    }
  }, []);

  // Watch initial filter
  useEffect(() => {
    setError(false);
    resultRef.current = initialFilter;
  }, [initialFilter]);

  return (
    <Dialog
      open={open}
      onClose={onClose}
      fullWidth
      maxWidth="lg"
      aria-labelledby="filter-dialog-title"
      aria-describedby="filter-dialog-description"
    >
      <DialogTitle id="filter-dialog-title">{t('common.filter.dialogTitle')}</DialogTitle>
      <DialogContent>
        <FilterForm
          filterDefinitionModel={filterDefinitionModel}
          onChange={localOnChange}
          initialFilter={initialFilter}
          predefinedFilterObjects={predefinedFilterObjects}
        />
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
          disabled={error}
          onClick={() => {
            onSubmit(resultRef.current as T);
          }}
          autoFocus
        >
          {t('common.applyAction')}
        </Button>
      </DialogActions>
    </Dialog>
  );
}

export default FilterDialog;

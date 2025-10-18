import React, { useCallback, useEffect, useRef, useState } from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { useTranslation } from 'react-i18next';
import FilterForm from '../internal/components/FilterForm';
import { type FilterDefinitionFieldsModel } from '../../../models/general';
import { type FilterValueObject, type PredefinedFilter } from '../internal/types';

export interface Props<T extends FilterValueObject> {
  readonly onSubmit: (filter: T) => void;
  readonly onReset: () => void;
  readonly onClose: () => void;
  readonly open: boolean;
  readonly filterDefinitionModel: FilterDefinitionFieldsModel;
  readonly predefinedFilterObjects?: PredefinedFilter[];
  readonly initialFilter?: undefined | null | T;
}

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
    setError(() => {
      resultRef.current = initialFilter;

      return false;
    });
  }, [initialFilter]);

  return (
    <Dialog
      aria-describedby="filter-dialog-description"
      aria-labelledby="filter-dialog-title"
      fullWidth
      maxWidth="lg"
      onClose={onClose}
      open={open}
    >
      <DialogTitle id="filter-dialog-title">{t('common.filter.dialogTitle')}</DialogTitle>
      <DialogContent>
        <FilterForm
          filterDefinitionModel={filterDefinitionModel}
          initialFilter={initialFilter}
          onChange={localOnChange}
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
          autoFocus
          disabled={error}
          onClick={() => {
            onSubmit(resultRef.current as T);
          }}
          variant="contained"
        >
          {t('common.applyAction')}
        </Button>
      </DialogActions>
    </Dialog>
  );
}

export default FilterDialog;

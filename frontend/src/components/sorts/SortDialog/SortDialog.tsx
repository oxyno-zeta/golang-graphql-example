import React, { useEffect, useState } from 'react';
import Button from '@mui/material/Button';
import DialogActions from '@mui/material/DialogActions';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { useTranslation } from 'react-i18next';
import { SortOrderModel, SortOrderFieldModel } from '../../../models/general';
import SortForm from '../internal/SortForm';

type Props<T extends Record<string, SortOrderModel>> = {
  onSubmit: (sort: T) => void;
  onReset: () => void;
  onClose: () => void;
  open: boolean;
  initialSort: null | undefined | T;
  sortFields: SortOrderFieldModel[];
};

function SortDialog<T extends Record<string, SortOrderModel>>({
  onSubmit,
  onReset,
  onClose,
  open,
  initialSort,
  sortFields,
}: Props<T>) {
  // Setup translate
  const { t } = useTranslation();
  // State
  const [result, setResult] = useState<T | undefined | null>(initialSort);

  // Watch initial filter
  useEffect(() => {
    setResult(initialSort ? { ...initialSort } : initialSort);
  }, [initialSort, open]);
  return (
    <Dialog
      open={open}
      onClose={onClose}
      fullWidth
      maxWidth="lg"
      aria-labelledby="sort-dialog-title"
      aria-describedby="sort-dialog-description"
    >
      <DialogTitle id="sort-dialog-title">{t('common.sort.dialogTitle')}</DialogTitle>
      <DialogContent>
        <SortForm
          sort={result}
          sortFields={sortFields}
          onChange={(f, v) => {
            // Initialize
            const res = result || ({} as Record<string, SortOrderModel>);
            // Update
            res[f] = v;

            // Save
            setResult({ ...res } as T);
          }}
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
          onClick={() => {
            onSubmit(result as T);
          }}
          autoFocus
        >
          {t('common.applyAction')}
        </Button>
      </DialogActions>
    </Dialog>
  );
}

export default SortDialog;

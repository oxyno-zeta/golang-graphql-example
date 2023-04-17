import React from 'react';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import { useTranslation } from 'react-i18next';
import { SortOrderModel, SortOrderFieldModel } from '../../../models/general';
import SortForm from '../internal/SortForm';

export type Props<T extends Record<string, SortOrderModel>> = {
  onSubmit: (sort: T[]) => void;
  onReset: () => void;
  onClose: () => void;
  open: boolean;
  initialSorts: null | undefined | T[];
  sortFields: SortOrderFieldModel[];
};

function SortDialog<T extends Record<string, SortOrderModel>>({
  onSubmit,
  onReset,
  onClose,
  open,
  initialSorts,
  sortFields,
}: Props<T>) {
  // Setup translate
  const { t } = useTranslation();

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
      {open && <SortForm initialSorts={initialSorts} onReset={onReset} onSubmit={onSubmit} sortFields={sortFields} />}
    </Dialog>
  );
}

export default SortDialog;

import React from 'react';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import { useTranslation } from 'react-i18next';
import { type SortOrderModel, type SortOrderFieldModel } from '../../../models/general';
import SortForm from '../internal/SortForm';

export interface Props<T extends Record<string, SortOrderModel>> {
  readonly onSubmit: (sort: T[]) => void;
  readonly onReset: () => void;
  readonly onClose: () => void;
  readonly open: boolean;
  readonly initialSorts: null | undefined | T[];
  readonly sortFields: SortOrderFieldModel[];
}

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
      aria-describedby="sort-dialog-description"
      aria-labelledby="sort-dialog-title"
      fullWidth
      maxWidth="lg"
      onClose={onClose}
      open={open}
    >
      <DialogTitle id="sort-dialog-title">{t('common.sort.dialogTitle')}</DialogTitle>
      {open ? (
        <SortForm initialSorts={initialSorts} onReset={onReset} onSubmit={onSubmit} sortFields={sortFields} />
      ) : null}
    </Dialog>
  );
}

export default SortDialog;

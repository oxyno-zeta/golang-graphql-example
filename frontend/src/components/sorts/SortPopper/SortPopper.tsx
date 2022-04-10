import React, { useEffect, useState } from 'react';
import Button from '@mui/material/Button';
import DialogActions from '@mui/material/DialogActions';
import Popper from '@mui/material/Popper';
import Paper from '@mui/material/Paper';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import { useTranslation } from 'react-i18next';
import { SortOrderModel, SortOrderFieldModel } from '../../../models/general';
import SortForm from '../internal/SortForm';

type Props<T extends Record<string, SortOrderModel>> = {
  onSubmit: (sort: T) => void;
  onReset: () => void;
  onClose: () => void;
  open: boolean;
  anchorElement: HTMLButtonElement | null;
  initialSort: null | undefined | T;
  sortFields: SortOrderFieldModel[];
};

function SortPopper<T extends Record<string, SortOrderModel>>({
  onSubmit,
  onReset,
  onClose,
  open,
  anchorElement,
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
    <Popper
      open={open}
      anchorEl={anchorElement}
      placement="bottom-end"
      disablePortal={false}
      modifiers={[]}
      sx={{
        minWidth: {
          xs: '300px',
          sm: '550px',
          md: '700px',
          lg: '950px',
          xl: '1100px',
        },
        maxHeight: '500px',
        overflowY: 'auto',
      }}
    >
      <ClickAwayListener onClickAway={onClose}>
        <Paper
          variant="outlined"
          sx={{
            borderTopLeftRadius: 0,
            borderTopRightRadius: 0,
            borderTop: '0px',
          }}
        >
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
        </Paper>
      </ClickAwayListener>
    </Popper>
  );
}

export default SortPopper;

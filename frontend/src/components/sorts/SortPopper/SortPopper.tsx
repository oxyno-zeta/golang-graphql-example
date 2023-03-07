import React from 'react';
import Popper from '@mui/material/Popper';
import Paper from '@mui/material/Paper';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import { SortOrderModel, SortOrderFieldModel } from '../../../models/general';
import SortForm from '../internal/SortForm';

type Props<T extends Record<string, SortOrderModel>> = {
  onSubmit: (sort: T[]) => void;
  onReset: () => void;
  onClose: () => void;
  open: boolean;
  anchorElement: HTMLButtonElement | null;
  initialSorts: null | undefined | T[];
  sortFields: SortOrderFieldModel[];
};

function SortPopper<T extends Record<string, SortOrderModel>>({
  onSubmit,
  onReset,
  onClose,
  open,
  anchorElement,
  initialSorts,
  sortFields,
}: Props<T>) {
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
          {open && (
            <SortForm initialSorts={initialSorts} onReset={onReset} onSubmit={onSubmit} sortFields={sortFields} />
          )}
        </Paper>
      </ClickAwayListener>
    </Popper>
  );
}

export default SortPopper;

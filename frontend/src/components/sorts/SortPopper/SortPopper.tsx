import React from 'react';
import Popper from '@mui/material/Popper';
import Paper from '@mui/material/Paper';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import { type SortOrderModel, type SortOrderFieldModel } from '../../../models/general';
import SortForm from '../internal/SortForm';

export interface Props<T extends Record<string, SortOrderModel>> {
  readonly onSubmit: (sort: T[]) => void;
  readonly onReset: () => void;
  readonly onClose: () => void;
  readonly open: boolean;
  readonly anchorElement: HTMLButtonElement | null;
  readonly initialSorts: null | undefined | T[];
  readonly sortFields: SortOrderFieldModel[];
}

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
      anchorEl={anchorElement}
      disablePortal={false}
      modifiers={[]}
      open={open}
      placement="bottom-end"
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
          sx={{
            borderTopLeftRadius: 0,
            borderTopRightRadius: 0,
            borderTop: '0px',
          }}
          variant="outlined"
        >
          {open ? (
            <SortForm initialSorts={initialSorts} onReset={onReset} onSubmit={onSubmit} sortFields={sortFields} />
          ) : null}
        </Paper>
      </ClickAwayListener>
    </Popper>
  );
}

export default SortPopper;

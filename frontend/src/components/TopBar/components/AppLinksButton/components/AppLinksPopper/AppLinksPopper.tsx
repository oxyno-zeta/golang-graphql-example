import React, { type ReactNode } from 'react';
import Popper from '@mui/material/Popper';
import Paper from '@mui/material/Paper';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import type { SxProps } from '@mui/material';

export interface Props {
  readonly onClose: () => void;
  readonly open: boolean;
  readonly anchorElement: HTMLFormElement | HTMLButtonElement | null;
  readonly children: ReactNode;
  readonly popperSx?: SxProps;
}

function AppLinksPopper({ onClose, open, anchorElement, children, popperSx = {} }: Props) {
  return (
    <Popper
      anchorEl={anchorElement}
      disablePortal={false}
      modifiers={[]}
      open={open}
      placement="bottom-end"
      sx={{
        minWidth: '300px',
        maxHeight: '500px',
        overflowY: 'auto',
        zIndex: 1200, // 1300 is the topbar, 0 won't overlap the left drawer
        ...popperSx,
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
          {children}
        </Paper>
      </ClickAwayListener>
    </Popper>
  );
}

export default AppLinksPopper;

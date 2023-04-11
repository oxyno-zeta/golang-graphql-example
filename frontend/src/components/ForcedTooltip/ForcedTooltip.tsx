import React, { ReactElement, useState } from 'react';
import MuiTooltip, { TooltipProps } from '@mui/material/Tooltip';
import ClickAwayListener from '@mui/material/ClickAwayListener';

export type Props = Omit<TooltipProps, 'onClose' | 'onOpen' | 'open' | 'children'> & {
  render?: (handleTooltipOpen: () => void, handleTooltipClose: () => void) => ReactElement;
  children?: ReactElement;
};

const defaultProps = {
  render: undefined,
  children: undefined,
};

function ForcedTooltip({ render, children, ...props }: Props) {
  const [open, setOpen] = useState(false);

  const handleTooltipClose = () => {
    setOpen(false);
  };

  const handleTooltipOpen = () => {
    setOpen(true);
  };

  return (
    <ClickAwayListener onClickAway={handleTooltipClose}>
      <div>
        <MuiTooltip
          PopperProps={{
            disablePortal: true,
          }}
          onClose={handleTooltipClose}
          onOpen={handleTooltipOpen}
          open={open}
          {...props}
        >
          {children || (render && render(handleTooltipOpen, handleTooltipClose)) || <div />}
        </MuiTooltip>
      </div>
    </ClickAwayListener>
  );
}

ForcedTooltip.defaultProps = defaultProps;

export default ForcedTooltip;

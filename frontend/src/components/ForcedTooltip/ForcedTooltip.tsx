import React, { ReactElement, useState } from 'react';
import MuiTooltip, { TooltipProps } from '@mui/material/Tooltip';
import ClickAwayListener from '@mui/material/ClickAwayListener';

export type Props = Omit<TooltipProps, 'onClose' | 'onOpen' | 'open' | 'children'> & {
  render?: (onTooltipOpen: () => void, onTooltipClose: () => void) => ReactElement;
  children?: ReactElement;
};

function ForcedTooltip({ render = undefined, children = undefined, ...props }: Props) {
  const [open, setOpen] = useState(false);

  const onTooltipClose = () => {
    setOpen(false);
  };

  const onTooltipOpen = () => {
    setOpen(true);
  };

  return (
    <ClickAwayListener onClickAway={onTooltipClose}>
      <div>
        <MuiTooltip
          PopperProps={{
            disablePortal: true,
          }}
          onClose={onTooltipClose}
          onOpen={onTooltipOpen}
          open={open}
          {...props}
        >
          {children || (render && render(onTooltipOpen, onTooltipClose)) || <div />}
        </MuiTooltip>
      </div>
    </ClickAwayListener>
  );
}

export default ForcedTooltip;

import React, { ReactNode, useState } from 'react';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Tooltip, { TooltipProps } from '@mui/material/Tooltip';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiHelpCircle } from '@mdi/js';
import ClickAwayListener from '@mui/material/ClickAwayListener';

interface Props {
  tooltipTitle: ReactNode;
  tooltipProps?: Omit<TooltipProps, 'onClose' | 'onOpen' | 'open' | 'title'>;
  iconButtonProps?: Omit<IconButtonProps, 'onClick'>;
}

const defaultProps = {
  tooltipProps: {},
  iconButtonProps: {},
};

function HelpTooltipButton({ tooltipTitle, tooltipProps, iconButtonProps }: Props) {
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
        <Tooltip
          PopperProps={{
            disablePortal: true,
          }}
          onClose={handleTooltipClose}
          onOpen={handleTooltipOpen}
          open={open}
          title={<>{tooltipTitle}</>}
          {...tooltipProps}
        >
          <span>
            <IconButton onClick={handleTooltipOpen} {...iconButtonProps}>
              <SvgIcon>
                <path d={mdiHelpCircle} />
              </SvgIcon>
            </IconButton>
          </span>
        </Tooltip>
      </div>
    </ClickAwayListener>
  );
}

HelpTooltipButton.defaultProps = defaultProps;

export default HelpTooltipButton;

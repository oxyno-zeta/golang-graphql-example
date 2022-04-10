import React, { ReactNode, useState } from 'react';
import IconButton from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import HelpIcon from '@mui/icons-material/Help';
import ClickAwayListener from '@mui/material/ClickAwayListener';

interface Props {
  tooltipTitle: ReactNode;
}

function HelpTooltipButton({ tooltipTitle }: Props) {
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
        >
          <IconButton onClick={handleTooltipOpen}>
            <HelpIcon />
          </IconButton>
        </Tooltip>
      </div>
    </ClickAwayListener>
  );
}

export default HelpTooltipButton;

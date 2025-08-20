import React, { type ReactNode } from 'react';
import IconButton, { type IconButtonProps } from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiHelpCircle } from '@mdi/js';
import ForcedTooltip, { type ForcedTooltipProps } from '../ForcedTooltip';

export interface Props {
  readonly tooltipTitle: ReactNode;
  readonly tooltipProps?: Omit<ForcedTooltipProps, 'title'>;
  readonly iconButtonProps?: Omit<IconButtonProps, 'onClick'>;
}

function HelpTooltipButton({ tooltipTitle, tooltipProps = {}, iconButtonProps = {} }: Props) {
  return (
    <ForcedTooltip
      title={<>{tooltipTitle}</>}
      {...tooltipProps}
      render={(onTooltipOpen) => (
        <span>
          <IconButton onClick={onTooltipOpen} {...iconButtonProps}>
            <SvgIcon>
              <path d={mdiHelpCircle} />
            </SvgIcon>
          </IconButton>
        </span>
      )}
    />
  );
}

export default HelpTooltipButton;

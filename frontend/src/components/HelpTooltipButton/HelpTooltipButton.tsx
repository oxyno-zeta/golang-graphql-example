import React, { ReactNode } from 'react';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiHelpCircle } from '@mdi/js';
import ForcedTooltip, { ForcedTooltipProps } from '../ForcedTooltip';

export interface Props {
  tooltipTitle: ReactNode;
  tooltipProps?: Omit<ForcedTooltipProps, 'title'>;
  iconButtonProps?: Omit<IconButtonProps, 'onClick'>;
}

const defaultProps = {
  tooltipProps: {},
  iconButtonProps: {},
};

function HelpTooltipButton({ tooltipTitle, tooltipProps, iconButtonProps }: Props) {
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

HelpTooltipButton.defaultProps = defaultProps;

export default HelpTooltipButton;

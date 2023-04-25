import React, { ReactNode } from 'react';
import SvgIcon, { SvgIconProps } from '@mui/material/SvgIcon';
import { mdiHelpCircle } from '@mdi/js';
import ForcedTooltip, { ForcedTooltipProps } from '../ForcedTooltip';

export interface Props {
  tooltipTitle: ReactNode;
  tooltipProps?: Omit<ForcedTooltipProps, 'title'>;
  svgIconProps?: Omit<SvgIconProps, 'onClick'>;
  svgIconContainerStyle?: Record<string, string>;
}

const defaultProps = {
  tooltipProps: {},
  svgIconProps: {},
  svgIconContainerStyle: { padding: '8px', verticalAlign: 'middle', textAlign: 'center', display: 'inline-flex' },
};

function HelpForcedTooltip({ tooltipTitle, tooltipProps, svgIconProps, svgIconContainerStyle }: Props) {
  return (
    <ForcedTooltip
      title={<>{tooltipTitle}</>}
      {...tooltipProps}
      render={(onTooltipOpen) => (
        <span style={svgIconContainerStyle}>
          <SvgIcon onClick={() => onTooltipOpen()} {...svgIconProps}>
            <path d={mdiHelpCircle} />
          </SvgIcon>
        </span>
      )}
    />
  );
}

HelpForcedTooltip.defaultProps = defaultProps;

export default HelpForcedTooltip;

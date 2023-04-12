import React, { ReactNode } from 'react';
import SvgIcon, { SvgIconProps } from '@mui/material/SvgIcon';
import { mdiInformation } from '@mdi/js';
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

function InfoForcedTooltip({ tooltipTitle, tooltipProps, svgIconProps, svgIconContainerStyle }: Props) {
  return (
    <ForcedTooltip
      title={<>{tooltipTitle}</>}
      {...tooltipProps}
      render={(handleTooltipOpen) => (
        <span style={svgIconContainerStyle}>
          <SvgIcon onClick={() => handleTooltipOpen()} {...svgIconProps}>
            <path d={mdiInformation} />
          </SvgIcon>
        </span>
      )}
    />
  );
}

InfoForcedTooltip.defaultProps = defaultProps;

export default InfoForcedTooltip;

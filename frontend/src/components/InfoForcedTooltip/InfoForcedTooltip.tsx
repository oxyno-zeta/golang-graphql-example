import React, { ReactNode } from 'react';
import SvgIcon, { SvgIconProps } from '@mui/material/SvgIcon';
import { mdiInformation } from '@mdi/js';
import ForcedTooltip, { ForcedTooltipProps } from '../ForcedTooltip';

export interface Props {
  readonly tooltipTitle: ReactNode;
  readonly tooltipProps?: Omit<ForcedTooltipProps, 'title'>;
  readonly svgIconProps?: Omit<SvgIconProps, 'onClick'>;
  readonly svgIconContainerStyle?: Record<string, string>;
}

function InfoForcedTooltip({
  tooltipTitle,
  tooltipProps = {},
  svgIconProps = {},
  svgIconContainerStyle = {
    padding: '8px',
    verticalAlign: 'middle',
    textAlign: 'center',
    display: 'inline-flex',
  },
}: Props) {
  return (
    <ForcedTooltip
      title={<>{tooltipTitle}</>}
      {...tooltipProps}
      render={(onTooltipOpen) => (
        <span style={svgIconContainerStyle}>
          <SvgIcon onClick={() => onTooltipOpen()} {...svgIconProps}>
            <path d={mdiInformation} />
          </SvgIcon>
        </span>
      )}
    />
  );
}

export default InfoForcedTooltip;

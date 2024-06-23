import React, { ReactNode } from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import type { BoxProps, TypographyProps } from '@mui/material';

export interface Props {
  labelText?: string;
  labelElement?: ReactNode;
  boxProps?: BoxProps;
  labelTypographyProps?: TypographyProps;
  valueElement?: ReactNode;
  valueText?: string;
  valueTypographyProps?: TypographyProps;
}

function OneLineLabelValueDisplay({
  boxProps = {},
  labelTypographyProps = {},
  labelElement = undefined,
  labelText = '',
  valueElement = undefined,
  valueText = '',
  valueTypographyProps = {},
}: Props) {
  return (
    <Box sx={{ display: { sm: 'flex' } }} {...boxProps}>
      {labelText && (
        <Typography sx={{ fontWeight: 'bold', marginRight: { sm: '5px' } }} {...labelTypographyProps}>
          {labelText}
        </Typography>
      )}
      {labelElement && <>{labelElement}</>}
      {valueText && <Typography {...valueTypographyProps}>{valueText}</Typography>}
      {valueElement && <>{valueElement}</>}
    </Box>
  );
}

export default OneLineLabelValueDisplay;

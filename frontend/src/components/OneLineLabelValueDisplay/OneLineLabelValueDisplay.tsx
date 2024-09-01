import React, { ReactNode } from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import type { BoxProps, TypographyProps } from '@mui/material';

export interface Props {
  readonly labelText?: string;
  readonly labelElement?: ReactNode;
  readonly boxProps?: BoxProps;
  readonly labelTypographyProps?: TypographyProps;
  readonly valueElement?: ReactNode;
  readonly valueText?: string;
  readonly valueTypographyProps?: TypographyProps;
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
      {labelText ? (
        <Typography sx={{ fontWeight: 'bold', marginRight: { sm: '5px' } }} {...labelTypographyProps}>
          {labelText}
        </Typography>
      ) : null}
      {labelElement ? <>{labelElement}</> : null}
      {valueText ? <Typography {...valueTypographyProps}>{valueText}</Typography> : null}
      {valueElement ? <>{valueElement}</> : null}
    </Box>
  );
}

export default OneLineLabelValueDisplay;

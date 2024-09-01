import React from 'react';
import Box from '@mui/material/Box';
import CircularProgress, { CircularProgressProps } from '@mui/material/CircularProgress';
import type { SxProps, TypographyProps } from '@mui/material';
import Typography from '@mui/material/Typography';

export interface Props {
  /**
   * Container Box sx.
   * Following this documentation: https://mui.com/system/getting-started/the-sx-prop/ .
   */
  readonly containerBoxSx?: SxProps;
  /**
   * Circular progress props.
   * Following this documentation: https://mui.com/material-ui/api/circular-progress/#props .
   */
  readonly circularProgressProps?: Partial<CircularProgressProps>;
  /**
   * subtitle.
   */
  readonly subtitle?: string;
  /**
   * Typography props.
   * Following this documentation: https://mui.com/material-ui/api/typography/#props .
   */
  readonly subtitleTypographyProps?: Partial<TypographyProps>;
}

function CenterLoading({
  containerBoxSx = {},
  circularProgressProps = {},
  subtitle = undefined,
  subtitleTypographyProps = {},
}: Props) {
  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        margin: '10px 0',
        ...containerBoxSx,
      }}
    >
      <CircularProgress {...circularProgressProps} />
      {subtitle ? (
        <Typography
          color="text.secondary"
          sx={{ fontSize: '12px', marginTop: '7px' }}
          variant="subtitle2"
          {...subtitleTypographyProps}
        >
          {subtitle}
        </Typography>
      ) : null}
    </Box>
  );
}

export default CenterLoading;

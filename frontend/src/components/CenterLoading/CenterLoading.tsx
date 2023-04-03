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
  containerBoxSx?: SxProps;
  /**
   * Circular progress props.
   * Following this documentation: https://mui.com/material-ui/api/circular-progress/#props .
   */
  circularProgressProps?: Partial<CircularProgressProps>;
  /**
   * subtitle.
   */
  subtitle?: string;
  /**
   * Typography props.
   * Following this documentation: https://mui.com/material-ui/api/typography/#props .
   */
  subtitleTypographyProps?: Partial<TypographyProps>;
}

const defaultProps = {
  containerBoxSx: {},
  circularProgressProps: {},
  subtitleTypographyProps: {},
  subtitle: undefined,
};

function CenterLoading({ containerBoxSx, circularProgressProps, subtitle, subtitleTypographyProps }: Props) {
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
      {subtitle && (
        <Typography
          color="text.secondary"
          variant="subtitle2"
          sx={{ fontSize: '12px', marginTop: '7px' }}
          {...subtitleTypographyProps}
        >
          {subtitle}
        </Typography>
      )}
    </Box>
  );
}

CenterLoading.defaultProps = defaultProps;

export default CenterLoading;

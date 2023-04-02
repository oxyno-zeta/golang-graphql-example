import React from 'react';
import Box from '@mui/material/Box';
import CircularProgress, { CircularProgressProps } from '@mui/material/CircularProgress';
import type { SxProps } from '@mui/material';

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
}

const defaultProps = {
  containerBoxSx: {},
  circularProgressProps: {},
};

function CenterLoading({ containerBoxSx, circularProgressProps }: Props) {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'center', margin: '10px 0', ...containerBoxSx }}>
      <CircularProgress {...circularProgressProps} />
    </Box>
  );
}

CenterLoading.defaultProps = defaultProps;

export default CenterLoading;

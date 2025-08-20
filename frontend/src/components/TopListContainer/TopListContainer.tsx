import React, { type ReactNode } from 'react';
import Box from '@mui/material/Box';
import type { SxProps } from '@mui/material';

export interface Props {
  readonly sx?: SxProps;
  readonly children: ReactNode;
}

function TopListContainer({ children, sx = {} }: Props) {
  return (
    <Box
      sx={{
        display: 'flex',
        flexFlow: 'row wrap',
        margin: '10px',
        gap: '5px',
        justifyContent: { xs: 'center', sm: 'flex-end' },
        ...sx,
      }}
    >
      {children}
    </Box>
  );
}

export default TopListContainer;

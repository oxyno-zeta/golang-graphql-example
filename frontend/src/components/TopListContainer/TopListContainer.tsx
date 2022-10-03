import React, { ReactNode } from 'react';
import Box from '@mui/material/Box';
import { SxProps } from '@mui/material';

interface Props {
  sx?: SxProps;
  children: ReactNode;
}

const defaultProps = {
  sx: {},
};

function TopListContainer({ children, sx }: Props) {
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

TopListContainer.defaultProps = defaultProps;

export default TopListContainer;

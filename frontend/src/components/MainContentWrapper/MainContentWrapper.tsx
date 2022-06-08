import React, { ReactNode } from 'react';
import Box from '@mui/material/Box';
import { TopBarSpacer } from '../TopBar';

interface Props {
  children: ReactNode;
}

function MainContentWrapper({ children }: Props) {
  return (
    <Box sx={{ margin: '0 20px 20px 20px' }}>
      <TopBarSpacer />
      {children}
    </Box>
  );
}

export default MainContentWrapper;

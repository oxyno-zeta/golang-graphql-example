import React, { ReactNode } from 'react';
import Box from '@mui/material/Box';
import type { SxProps } from '@mui/material';
import { TopBarSpacer } from '../TopBar';

interface Props {
  children: ReactNode;
  containerBoxSx?: SxProps;
}

const defaultProps = {
  containerBoxSx: {},
};

function MainContentWrapper({ children, containerBoxSx }: Props) {
  return (
    <Box sx={{ margin: '0 20px 20px 20px', ...containerBoxSx }}>
      <TopBarSpacer />
      {children}
    </Box>
  );
}

MainContentWrapper.defaultProps = defaultProps;

export default MainContentWrapper;

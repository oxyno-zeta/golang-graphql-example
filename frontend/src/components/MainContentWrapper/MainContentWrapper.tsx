import React, { type ReactNode } from 'react';
import Box from '@mui/material/Box';
import type { SxProps } from '@mui/material';
import { TopBarSpacer } from '../TopBar';

interface Props {
  readonly children: ReactNode;
  readonly containerBoxSx?: SxProps;
  readonly disableTopSpacer?: boolean;
}

function MainContentWrapper({ children, disableTopSpacer = false, containerBoxSx = {} }: Props) {
  return (
    <Box sx={{ margin: '0 20px 20px 20px', ...containerBoxSx }}>
      {!disableTopSpacer && <TopBarSpacer />}
      {children}
    </Box>
  );
}

export default MainContentWrapper;

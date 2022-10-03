import React, { ReactNode, useState } from 'react';
import Box from '@mui/material/Box';
import Drawer, { DrawerProps } from '@mui/material/Drawer';
import type { SxProps } from '@mui/material';
import MainContentWrapper from '../../MainContentWrapper';

interface Props {
  drawerWidth: number | string;
  drawerContentElement: ReactNode;
  renderContent: (handleDrawerToggle: () => void) => ReactNode;
  mobileDrawerProps?: Partial<Omit<DrawerProps, 'open' | 'onClose'>>;
  drawerProps?: Partial<Omit<DrawerProps, 'open'>>;
  drawerContainerBoxSx?: SxProps;
  mainContainerBoxSx?: SxProps;
}

const defaultProps = {
  mobileDrawerProps: {},
  drawerProps: {},
  drawerContainerBoxSx: {},
  mainContainerBoxSx: {},
};

function PageDrawer({
  drawerContentElement,
  drawerWidth,
  renderContent,
  mobileDrawerProps,
  drawerProps,
  drawerContainerBoxSx,
  mainContainerBoxSx,
}: Props) {
  // States
  const [isMobileOpen, setMobileOpen] = useState(false);

  const handleMobileDrawerToggle = () => {
    setMobileOpen((v) => !v);
  };

  return (
    <div style={{ display: 'flex' }}>
      <Box component="nav" sx={{ width: { lg: drawerWidth }, flexShrink: { lg: 0 }, ...drawerContainerBoxSx }}>
        {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
        <Drawer
          variant="temporary"
          open={isMobileOpen}
          onClose={handleMobileDrawerToggle}
          ModalProps={{
            keepMounted: true, // Better open performance on mobile.
          }}
          sx={{
            display: { xs: 'block', lg: 'none' },
            '& .MuiDrawer-paper': {
              boxSizing: 'border-box',
              width: drawerWidth,
            },
          }}
          {...mobileDrawerProps}
        >
          {drawerContentElement}
        </Drawer>
        <Drawer
          variant="persistent"
          sx={{
            display: { xs: 'none', lg: 'block' },
            '& .MuiDrawer-paper': {
              boxSizing: 'border-box',
              width: drawerWidth,
            },
          }}
          open
          {...drawerProps}
        >
          {drawerContentElement}
        </Drawer>
      </Box>

      <Box
        component="main"
        sx={{
          flexGrow: 1,
          width: { sm: `calc(100% - ${drawerWidth}px)` },
          overflowY: 'auto',
          ...mainContainerBoxSx,
        }}
      >
        <MainContentWrapper>
          <div style={{ height: '20px' }} />
          {renderContent(handleMobileDrawerToggle)}
        </MainContentWrapper>
      </Box>
    </div>
  );
}

PageDrawer.defaultProps = defaultProps;

export default PageDrawer;

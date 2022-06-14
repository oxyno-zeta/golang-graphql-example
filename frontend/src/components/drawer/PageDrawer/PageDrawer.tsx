import React, { ReactNode, useState } from 'react';
import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import MainContentWrapper from '../../MainContentWrapper';

interface Props {
  drawerWidth: number | string;
  drawerContentElement: ReactNode;
  renderContent: (handleDrawerToggle: () => void) => ReactNode;
}

function PageDrawer({ drawerContentElement, drawerWidth, renderContent }: Props) {
  // States
  const [isMobileOpen, setMobileOpen] = useState(false);

  const handleMobileDrawerToggle = () => {
    setMobileOpen((v) => !v);
  };

  return (
    <div style={{ display: 'flex' }}>
      <Box component="nav" sx={{ width: { lg: drawerWidth }, flexShrink: { lg: 0 } }}>
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
        >
          {drawerContentElement}
        </Drawer>
      </Box>

      <Box
        component="main"
        sx={{
          flexGrow: 1,
          padding: '0 20px 20px 20px',
          width: { sm: `calc(100% - ${drawerWidth}px)` },
          overflowY: 'auto',
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

export default PageDrawer;

import React, { ReactNode } from 'react';
import Box from '@mui/material/Box';
import Drawer, { DrawerProps } from '@mui/material/Drawer';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import type { SxProps } from '@mui/material';

interface Props {
  drawerWidth: number | string;
  drawerElement: ReactNode;
  handleMobileDrawerToggle: () => void;
  mobileDrawerProps?: Partial<Omit<DrawerProps, 'open' | 'onClose'>>;
  drawerProps?: Partial<Omit<DrawerProps, 'open'>>;
  drawerContainerBoxSx?: SxProps;
}

const defaultProps = {
  mobileDrawerProps: {},
  drawerProps: {},
  drawerContainerBoxSx: {},
};

function ContentDisplayDrawer({
  drawerElement,
  drawerWidth,
  handleMobileDrawerToggle,
  mobileDrawerProps,
  drawerProps,
  drawerContainerBoxSx,
}: Props) {
  const open = Boolean(drawerElement);
  const theme = useTheme();
  const sizeMatching = useMediaQuery(theme.breakpoints.up('lg'));

  return (
    <Box
      sx={{
        display: open ? 'block' : 'none',
        width: { lg: drawerWidth },
        flexShrink: { lg: 0 },
        ...drawerContainerBoxSx,
      }}
    >
      {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
      {open && !sizeMatching && (
        <Drawer
          variant="temporary"
          open={open}
          onClose={handleMobileDrawerToggle}
          anchor="right"
          sx={{
            display: 'block',
            '& .MuiDrawer-paper': {
              boxSizing: 'border-box',
              width: drawerWidth,
            },
          }}
          {...mobileDrawerProps}
        >
          {drawerElement}
        </Drawer>
      )}
      <Drawer
        variant="persistent"
        sx={{
          display: { xs: 'none', lg: 'block' },
          '& .MuiDrawer-paper': {
            boxSizing: 'border-box',
            width: drawerWidth,
          },
        }}
        anchor="right"
        open={open}
        {...drawerProps}
      >
        {drawerElement}
      </Drawer>
    </Box>
  );
}

ContentDisplayDrawer.defaultProps = defaultProps;

export default ContentDisplayDrawer;

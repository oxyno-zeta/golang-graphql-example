import React, { ReactNode, useCallback, useState } from 'react';
import Box from '@mui/material/Box';
import Drawer, { DrawerProps } from '@mui/material/Drawer';
import type { Theme, CSSObject } from '@mui/material';
import type { SystemStyleObject } from '@mui/system';
import Divider from '@mui/material/Divider';
import ListItem from '@mui/material/ListItem';
import List from '@mui/material/List';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import ListItemIcon from '@mui/material/ListItemIcon';
import { mdiChevronDoubleLeft, mdiChevronDoubleRight } from '@mdi/js';
import SvgIcon from '@mui/material/SvgIcon';
import { useTranslation } from 'react-i18next';
import { TopBarSpacer } from '~components/TopBar';
import MainContentWrapper from '../../MainContentWrapper';

interface DrawerContentProps {
  listItemButtonSx: SystemStyleObject<Theme>;
  listItemIconSx: SystemStyleObject<Theme>;
  listItemTextSx: SystemStyleObject<Theme>;
}

export interface Props {
  renderDrawerContent: (props: DrawerContentProps) => ReactNode;
  renderContent: (onDrawerToggle: () => void, isNormalCollapsed: boolean) => ReactNode;
  mobileDrawerProps?: Partial<Omit<DrawerProps, 'open' | 'onClose'>>;
  drawerProps?: Partial<Omit<DrawerProps, 'open'>>;
  drawerContainerBoxSx?: SystemStyleObject<Theme>;
  mainContainerBoxSx?: SystemStyleObject<Theme>;
  disableTopSpacer?: boolean;
  // Default drawer width for init.
  defaultDrawerWidth: number;
  minDrawerWidth?: number;
  maxDrawerWidth?: number;
  disableResize?: boolean;
  disableCollapse?: boolean;
}

const defaultMinDrawerWidth = 150;
const defaultMaxDrawerWidth = 400;

const defaultProps = {
  mobileDrawerProps: {},
  drawerProps: {},
  drawerContainerBoxSx: {},
  mainContainerBoxSx: {},
  disableTopSpacer: false,
  minDrawerWidth: defaultMinDrawerWidth,
  maxDrawerWidth: defaultMaxDrawerWidth,
  disableResize: false,
  disableCollapse: false,
};

function PageDrawer({
  defaultDrawerWidth,
  renderDrawerContent,
  renderContent,
  mobileDrawerProps,
  drawerProps,
  drawerContainerBoxSx,
  mainContainerBoxSx,
  disableTopSpacer,
  minDrawerWidth = defaultMinDrawerWidth,
  maxDrawerWidth = defaultMaxDrawerWidth,
  disableResize,
  disableCollapse,
}: Props) {
  // Setup translate
  const { t } = useTranslation();

  // States
  const [isMobileOpen, setMobileOpen] = useState(false);
  const [isNormalOpened, setNormalOpened] = useState(true);
  const [drawerWidth, setDrawerWidth] = useState(defaultDrawerWidth);

  const onMobileDrawerToggle = () => {
    setMobileOpen((v) => !v);
  };

  const onNormalDrawerToggle = () => {
    setNormalOpened((v) => !v);
  };

  const handleMouseMove = useCallback((e: MouseEvent) => {
    const newWidth = e.clientX - document.body.offsetLeft;

    if (newWidth > minDrawerWidth && newWidth < maxDrawerWidth) {
      setDrawerWidth(newWidth);
    }
  }, []);

  function handleMouseDown() {
    document.addEventListener('mouseup', handleMouseUp, true);
    document.addEventListener('mousemove', handleMouseMove, true);
  }

  function handleMouseUp() {
    document.removeEventListener('mouseup', handleMouseUp, true);
    document.removeEventListener('mousemove', handleMouseMove, true);
  }

  const openedMixin = (): CSSObject => ({
    width: `${drawerWidth}px`,
    overflowX: 'hidden',
  });

  const closedMixin = (theme: Theme): CSSObject => ({
    overflowX: 'hidden',
    width: `calc(${theme.spacing(7)} + 1px)`,
    [theme.breakpoints.up('sm')]: {
      width: `calc(${theme.spacing(8)} + 1px)`,
    },
  });

  return (
    <div style={{ display: 'flex' }}>
      <Box
        component="nav"
        sx={(theme: Theme) => ({
          flexShrink: { lg: 0 },
          width: { lg: isNormalOpened ? `${drawerWidth}px` : `calc(${theme.spacing(7)} + 1px)` },
          ...drawerContainerBoxSx,
        })}
      >
        {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
        <Drawer
          variant="temporary"
          open={isMobileOpen}
          onClose={onMobileDrawerToggle}
          ModalProps={{
            keepMounted: true, // Better open performance on mobile.
          }}
          sx={{
            display: { xs: 'block', lg: 'none' },
            '& .MuiDrawer-paper': {
              boxSizing: 'border-box',
            },
          }}
          {...mobileDrawerProps}
        >
          {!disableTopSpacer && <TopBarSpacer />}
          {renderDrawerContent({
            listItemButtonSx: {},
            listItemIconSx: {},
            listItemTextSx: {},
          })}
        </Drawer>
        <Drawer
          variant="persistent"
          open
          sx={(theme) => ({
            display: { xs: 'none', lg: 'block' },
            ...(isNormalOpened && {
              ...openedMixin(),
              '& .MuiDrawer-paper': { ...(disableResize ? {} : { borderRight: 'unset' }), ...openedMixin() },
            }),
            ...(!isNormalOpened && {
              ...closedMixin(theme),
              '& .MuiDrawer-paper': closedMixin(theme),
            }),
          })}
          {...drawerProps}
        >
          {isNormalOpened && !disableResize && (
            <Divider
              role="button"
              onMouseDown={() => handleMouseDown()}
              orientation="vertical"
              flexItem
              sx={(theme) => ({
                cursor: 'ew-resize',
                position: 'absolute',
                borderRightWidth: '1px',
                backgroundColor: 'unset',
                '&:hover': { borderWidth: '2px', backgroundColor: theme.palette.divider },
                width: 0,
                top: 0,
                right: 0,
                bottom: 0,
              })}
            />
          )}
          {!disableTopSpacer && <TopBarSpacer />}
          <div style={{ overflowY: 'auto', height: '100%', display: 'flex', flexDirection: 'column' }}>
            {renderDrawerContent({
              listItemButtonSx: isNormalOpened ? {} : { justifyContent: 'center' },
              listItemIconSx: isNormalOpened ? {} : { minWidth: 'unset' },
              listItemTextSx: isNormalOpened ? {} : { display: 'none' },
            })}
            {!disableCollapse && (
              <div style={{ marginTop: 'auto' }}>
                <List>
                  <ListItem disablePadding dense>
                    <ListItemButton
                      dense
                      onClick={onNormalDrawerToggle}
                      sx={isNormalOpened ? {} : { justifyContent: 'center' }}
                    >
                      <ListItemIcon sx={isNormalOpened ? {} : { minWidth: 'unset' }}>
                        <SvgIcon>
                          <path d={isNormalOpened ? mdiChevronDoubleLeft : mdiChevronDoubleRight} />
                        </SvgIcon>
                      </ListItemIcon>
                      <ListItemText sx={isNormalOpened ? {} : { display: 'none' }}>
                        {t('common.collapseSidebarAction')}
                      </ListItemText>
                    </ListItemButton>
                  </ListItem>
                </List>
              </div>
            )}
          </div>
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
        <MainContentWrapper disableTopSpacer={disableTopSpacer}>
          {!disableTopSpacer && <div style={{ height: '20px' }} />}
          {renderContent(onMobileDrawerToggle, isNormalOpened)}
        </MainContentWrapper>
      </Box>
    </div>
  );
}

PageDrawer.defaultProps = defaultProps;

export default PageDrawer;

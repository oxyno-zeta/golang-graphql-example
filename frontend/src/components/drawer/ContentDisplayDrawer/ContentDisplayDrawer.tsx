import React, { ReactNode, useCallback, useState } from 'react';
import Box from '@mui/material/Box';
import Drawer, { DrawerProps } from '@mui/material/Drawer';
import useMediaQuery from '@mui/material/useMediaQuery';
import Tooltip from '@mui/material/Tooltip';
import IconButton from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import { useTheme } from '@mui/material/styles';
import type { Theme } from '@mui/material';
import type { SystemStyleObject } from '@mui/system';
import { mdiClose } from '@mdi/js';
import { useTranslation } from 'react-i18next';
import { TopBarSpacer } from '~components/TopBar';

export interface Props {
  // Default drawer width for init.
  readonly defaultDrawerWidth: number;
  readonly drawerElement: ReactNode;
  readonly children: ReactNode;
  readonly onClose: () => void;
  readonly mobileDrawerProps?: Partial<Omit<DrawerProps, 'open' | 'onClose'>>;
  readonly drawerProps?: Partial<Omit<DrawerProps, 'open'>>;
  readonly mainContainerBoxSx?: Omit<SystemStyleObject<Theme>, 'width'>;
  readonly drawerContainerBoxSx?: Omit<SystemStyleObject<Theme>, 'width' | 'display' | 'flexShrink'>;
  readonly disableTopSpacer?: boolean;
  readonly titleElement?: ReactNode;
  readonly minDrawerWidth?: number;
  readonly maxDrawerWidth?: number;
  readonly disableResize?: boolean;
}

const defaultMinDrawerWidth = 150;
const defaultMaxDrawerWidth = 400;

function ContentDisplayDrawer({
  children,
  drawerElement,
  defaultDrawerWidth,
  onClose,
  mobileDrawerProps = {},
  drawerProps = {},
  drawerContainerBoxSx = {},
  mainContainerBoxSx = {},
  disableTopSpacer = false,
  titleElement = null,
  minDrawerWidth = defaultMinDrawerWidth,
  maxDrawerWidth = defaultMaxDrawerWidth,
  disableResize = false,
}: Props) {
  // Setup translate
  const { t } = useTranslation();
  // State
  const [drawerWidth, setDrawerWidth] = useState(defaultDrawerWidth);
  // Variables
  const open = Boolean(drawerElement);
  const theme = useTheme();
  const sizeMatching = useMediaQuery(theme.breakpoints.up('lg'));

  const handleMouseMove = useCallback(
    (e: MouseEvent) => {
      const newWidth = document.body.offsetWidth - (e.clientX - document.body.offsetLeft);

      if (newWidth > minDrawerWidth && newWidth < maxDrawerWidth) {
        setDrawerWidth(newWidth);
      }
    },
    [minDrawerWidth, maxDrawerWidth],
  );

  function handleMouseDown() {
    document.addEventListener('mouseup', handleMouseUp, true);
    document.addEventListener('mousemove', handleMouseMove, true);
  }

  function handleMouseUp() {
    document.removeEventListener('mouseup', handleMouseUp, true);
    document.removeEventListener('mousemove', handleMouseMove, true);
  }

  const content = (
    <>
      {!disableResize && (
        <Divider
          flexItem
          onMouseDown={() => handleMouseDown()}
          orientation="vertical"
          role="button"
          sx={(th) => ({
            cursor: 'ew-resize',
            position: 'absolute',
            borderRightWidth: '1px',
            backgroundColor: 'unset',
            '&:hover': { borderWidth: '2px', backgroundColor: th.palette.divider },
            width: 0,
            top: 0,
            left: 0,
            bottom: 0,
          })}
        />
      )}
      <div style={{ display: 'flex', alignItems: 'center', margin: '10px' }}>
        {titleElement || (
          <Typography sx={{ marginRight: 'auto' }} variant="h6">
            {t('common.details')}
          </Typography>
        )}
        <Tooltip title={<>{t('common.closeAction')}</>}>
          <span>
            <IconButton onClick={onClose} sx={{ marginLeft: 'auto' }}>
              <SvgIcon>
                <path d={mdiClose} />
              </SvgIcon>
            </IconButton>
          </span>
        </Tooltip>
      </div>
      {drawerElement}
    </>
  );

  return (
    <Box sx={{ ...(open ? { marginRight: `${drawerWidth}px` } : {}), ...mainContainerBoxSx }}>
      {children}
      <Box
        sx={{
          display: open ? 'block' : 'none',
          width: { lg: `${drawerWidth}px` },
          flexShrink: { lg: 0 },
          ...drawerContainerBoxSx,
        }}
      >
        {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
        {open && !sizeMatching ? (
          <Drawer
            anchor="right"
            onClose={onClose}
            open={open}
            sx={{
              display: 'block',
              '& .MuiDrawer-paper': {
                boxSizing: 'border-box',
                width: drawerWidth,
              },
            }}
            variant="temporary"
            {...mobileDrawerProps}
          >
            {content}
          </Drawer>
        ) : null}
        <Drawer
          anchor="right"
          open={open}
          sx={{
            display: { xs: 'none', lg: 'block' },
            '& .MuiDrawer-paper': {
              boxSizing: 'border-box',
              width: drawerWidth,
              ...(disableResize ? {} : { borderLeft: 'unset' }),
            },
          }}
          variant="persistent"
          {...drawerProps}
        >
          {!disableTopSpacer && <TopBarSpacer />}
          {content}
        </Drawer>
      </Box>
    </Box>
  );
}

export default ContentDisplayDrawer;

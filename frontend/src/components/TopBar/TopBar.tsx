import React from 'react';
import { Link } from 'react-router';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Avatar from '@mui/material/Avatar';
import { useTranslation } from 'react-i18next';
import TopBarUserMenu, { type TopBarUserMenuProps } from './TopBarUserMenu';

type TopBarProps = {
  readonly TopBarUserMenuComponent?: React.ElementType;
  readonly topBarUserMenuProps?: TopBarUserMenuProps;
  readonly disableUserMenu?: boolean;
};

function TopBar({
  TopBarUserMenuComponent = TopBarUserMenu,
  topBarUserMenuProps = {},
  disableUserMenu = false,
}: TopBarProps) {
  // Setup translate
  const { t } = useTranslation();

  return (
    <AppBar id="topbar" position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
      <Toolbar variant="dense">
        <Avatar component={Link} src="/logo.png" to="/" />
        <Typography
          component={Link}
          sx={{ marginLeft: '10px', textDecoration: 'none', color: 'inherit' }}
          to="/"
          variant="h6"
        >
          {t('common.mainTitle')}
        </Typography>
        <div style={{ flexGrow: 1 }} />
        {!disableUserMenu && <TopBarUserMenuComponent {...topBarUserMenuProps} />}
      </Toolbar>
    </AppBar>
  );
}

export default TopBar;

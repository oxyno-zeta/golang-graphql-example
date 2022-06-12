import React from 'react';
import { Link } from 'react-router-dom';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Avatar from '@mui/material/Avatar';
import { useTranslation } from 'react-i18next';
import IconToggleColorMode from '../theming/IconToggleColorMode';

function TopBar() {
  // Setup translate
  const { t } = useTranslation();

  return (
    <AppBar position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
      <Toolbar variant="dense">
        <Avatar src="/logo.png" component={Link} to="/" />
        <Typography
          variant="h6"
          component={Link}
          sx={{ marginLeft: '10px', textDecoration: 'none', color: 'inherit' }}
          to="/"
        >
          {t('common.mainTitle')}
        </Typography>
        <div style={{ flexGrow: 1 }} />
        <IconToggleColorMode />
      </Toolbar>
    </AppBar>
  );
}

export default TopBar;

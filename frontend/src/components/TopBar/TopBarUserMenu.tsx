import React, { useState, useContext } from 'react';
import { useTranslation } from 'react-i18next';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import Tooltip from '@mui/material/Tooltip';
import MenuItem from '@mui/material/MenuItem';
import Divider from '@mui/material/Divider';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiAccountCircle } from '@mdi/js';
import ConfigContext from '../../contexts/ConfigContext';
import ToggleColorModeMenuItem from '../theming/ToggleColorModeMenuItem';

//
// Build logout URL
//
// Logout URL must be build according to this example:
// ENCODED = http://{KEYCLOAK_URL}/auth/realms/{REALM_NAME}/protocol/openid-connect/logout?post_logout_redirect_uri={ENCODED_REDIRECT_URI}
// Documentation of Keycloak: https://www.keycloak.org/docs/latest/securing_apps/index.html#logout
// Final: /oauth2/sign_out?rd={ENCODED}
// According to: https://oauth2-proxy.github.io/oauth2-proxy/docs/features/endpoints/
function buildLogoutURL(signOutURLString: string, oidcClientID: string) {
  // Parse sign out url
  const signOutURL = new URL(signOutURLString);
  // Add redirect param
  // Blocked by Keycloak v19
  // Blocked by this PR: https://github.com/keycloak/keycloak/issues/12680
  // Workaround to have code ready
  if (signOutURLString !== '' && oidcClientID !== '') {
    // Encode current origin (aka http://DOMAIN_INCLUDING_PORTS)
    const currentEncodedURI = encodeURIComponent(window.location.origin);

    signOutURL.searchParams.set('post_logout_redirect_uri', currentEncodedURI);
    signOutURL.searchParams.set('client_id', currentEncodedURI);
  }
  // Encode it
  const targetEncodedURI = encodeURIComponent(signOutURL.toString());

  // Create final URL
  return `/oauth2/sign_out?rd=${targetEncodedURI}`;
}

function TopBarUserMenu() {
  // Setup translate
  const { t } = useTranslation();
  // States
  const [anchorElUser, setAnchorElUser] = useState<null | HTMLElement>(null);
  const config = useContext(ConfigContext);

  // Hooks
  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget.parentElement);
  };
  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  return (
    <>
      <Tooltip title={t('common.accountMenu')}>
        <span>
          <IconButton onClick={handleOpenUserMenu}>
            <SvgIcon>
              <path d={mdiAccountCircle} />
            </SvgIcon>
          </IconButton>
        </span>
      </Tooltip>
      <Menu
        sx={{ marginTop: '28px' }}
        anchorEl={anchorElUser}
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        keepMounted
        transformOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        open={Boolean(anchorElUser)}
        onClose={handleCloseUserMenu}
      >
        <div style={{ display: 'flex', alignItems: 'center', margin: '0 10px' }}>
          <div style={{ margin: '5px' }}>
            <div style={{ width: 'calc(100% - 20px)' }}>
              <Typography style={{ fontSize: 14 }} gutterBottom>
                Fake User
              </Typography>
              <Typography style={{ fontSize: 11 }} color="text.secondary" gutterBottom>
                fake@fake.com
              </Typography>
            </div>

            <div style={{ margin: '0 10px' }}>
              <Divider />
            </div>

            <div style={{ margin: '5px 0' }}>
              <ToggleColorModeMenuItem />
            </div>
          </div>
        </div>

        {config.oidcSignOutURL && (
          <div style={{ margin: '0 25px' }}>
            <Divider />
          </div>
        )}

        {config.oidcSignOutURL && (
          <MenuItem
            component="a"
            href={buildLogoutURL(config.oidcSignOutURL, config.oidcClientID || '')}
            rel="noopener noreferrer"
          >
            {t('common.signOutAction')}
          </MenuItem>
        )}
      </Menu>
    </>
  );
}

export default TopBarUserMenu;

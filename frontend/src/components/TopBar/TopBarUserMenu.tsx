import React, { useState, useContext } from 'react';
import { useTranslation } from 'react-i18next';
import IconButton from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import MenuItem from '@mui/material/MenuItem';
import Divider from '@mui/material/Divider';
import SvgIcon from '@mui/material/SvgIcon';
import MenuList from '@mui/material/MenuList';
import Popover from '@mui/material/Popover';
import { mdiAccountCircle } from '@mdi/js';
import ConfigContext from '../../contexts/ConfigContext';
import ToggleColorModeMenuItem from '../theming/ToggleColorModeMenuItem';
import TimezoneSelector from '../timezone/TimezoneSelector';
import UserInfo from './components/UserInfo';

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

  // Expand
  const { oidcSignOutURL, oidcClientID } = config;

  // Hooks
  const onOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget.parentElement);
  };
  const onCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  return (
    <>
      <Tooltip title={t('common.accountMenu')}>
        <span>
          <IconButton onClick={onOpenUserMenu}>
            <SvgIcon>
              <path d={mdiAccountCircle} />
            </SvgIcon>
          </IconButton>
        </span>
      </Tooltip>
      <Popover
        sx={{ marginTop: '28px' }}
        anchorEl={anchorElUser}
        anchorOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        transformOrigin={{
          vertical: 'top',
          horizontal: 'right',
        }}
        open={Boolean(anchorElUser)}
        onClose={onCloseUserMenu}
      >
        <div style={{ margin: '10px 16px 10px 16px' }}>
          <div style={{ width: 'calc(100% - 20px)' }}>
            <UserInfo />
          </div>

          <div style={{ margin: '0 9px' }}>
            <Divider />
          </div>

          <div style={{ margin: '5px 0' }}>
            <ToggleColorModeMenuItem />
          </div>
          <div style={{ margin: '15px 0 5px 0' }}>
            <TimezoneSelector />
          </div>
        </div>

        {oidcSignOutURL && (
          <div style={{ margin: '0 25px' }}>
            <Divider />
          </div>
        )}

        <MenuList>
          {oidcSignOutURL && (
            <MenuItem component="a" href={buildLogoutURL(oidcSignOutURL, oidcClientID || '')} rel="noopener noreferrer">
              {t('common.signOutAction')}
            </MenuItem>
          )}
        </MenuList>
      </Popover>
    </>
  );
}

export default TopBarUserMenu;

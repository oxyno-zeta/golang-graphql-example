import React, { type ReactNode, useContext, useMemo, useState } from 'react';
import Cookies from 'universal-cookie';
import PageDrawerSettingsContext from '~contexts/PageDrawerSettingsContext';
import ConfigContext from '~contexts/ConfigContext';

export interface Props {
  readonly children: ReactNode;
}

const cookieName = 'left-menu-collapsed';

function PageDrawerSettingsProvider({ children }: Props) {
  // Get cookies object
  const cookies = useMemo(() => new Cookies(), []);
  // Get stored collapsed menu value
  const storedCollapsedMenu = cookies.get(cookieName);

  // Compute initial value
  let initCollapsedVal = storedCollapsedMenu;
  if (!initCollapsedVal) {
    initCollapsedVal = false;
  } else {
    initCollapsedVal = initCollapsedVal === 'true';
  }

  // Get config from context
  const cfg = useContext(ConfigContext);

  // Expand
  const { configCookieDomain } = cfg;

  // States
  const [isCollapsed, setIsCollapsed] = useState<boolean>(initCollapsedVal);

  // Create color mode context
  const contextValue = useMemo(() => {
    // Set cookie
    const setCookie = (input: boolean) => {
      cookies.set(cookieName, input, {
        path: '/',
        maxAge: 31536000, // 1 year
        domain: configCookieDomain,
      });
    };

    return {
      isCollapsed: () => isCollapsed,
      toggleCollapsed: () => {
        setIsCollapsed((v) => {
          // Save in storage
          setCookie(!v);

          return !v;
        });
      },
    };
  }, [isCollapsed, configCookieDomain, cookies]);

  return <PageDrawerSettingsContext.Provider value={contextValue}>{children}</PageDrawerSettingsContext.Provider>;
}

export default PageDrawerSettingsProvider;

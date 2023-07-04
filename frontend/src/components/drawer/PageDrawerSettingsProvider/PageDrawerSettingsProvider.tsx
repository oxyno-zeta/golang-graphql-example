import React, { ReactNode, useContext, useMemo, useState } from 'react';
import PageDrawerSettingsContext from '~contexts/PageDrawerSettingsContext';
import Cookies from 'universal-cookie';
import ConfigContext from '~contexts/ConfigContext';

export interface Props {
  children: ReactNode;
}

const cookieName = 'left-menu-collapsed';

function PageDrawerSettingsProvider({ children }: Props) {
  // Get cookies object
  const cookies = new Cookies();
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
  const [isCollapsed, setCollapsed] = useState<boolean>(initCollapsedVal);

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
        setCollapsed((v) => {
          // Save in storage
          setCookie(!v);

          return !v;
        });
      },
    };
  }, [isCollapsed]);

  return <PageDrawerSettingsContext.Provider value={contextValue}>{children}</PageDrawerSettingsContext.Provider>;
}

export default PageDrawerSettingsProvider;

import React, { type ReactNode, useState, useMemo, useContext } from 'react';
import Cookies from 'universal-cookie';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import GridTableViewSwitcherContext from '~contexts/GridTableViewSwitcherContext';
import ConfigContext from '../../../contexts/ConfigContext';

export interface Props {
  readonly children: ReactNode;
}

const cookieName = 'grid-view';

function GridTableViewSwitcherProvider({ children }: Props) {
  // Get config from context
  const cfg = useContext(ConfigContext);
  // Get cookies object
  const cookies = useMemo(() => new Cookies(), []);
  // Get stored value
  const storedValue = cookies.get(cookieName);
  // Compute initial value
  let initVal = false;
  if (storedValue === null) {
    initVal = false;
  } else {
    initVal = storedValue === 'true';
  }
  // Theming
  const theme = useTheme();
  const sizeMatching = useMediaQuery(theme.breakpoints.down('lg'));
  // Check if screen size is matching
  if (sizeMatching) {
    initVal = true;
  }

  // State for grid view
  const [gridView, setGridView] = useState<boolean>(initVal);

  // Expand
  const { configCookieDomain } = cfg;

  // Create table grid view switcher context
  const tableGridViewSwitcherCtx = useMemo(() => {
    // Set cookie
    const setCookie = (input: boolean) => {
      cookies.set(cookieName, input, {
        path: '/',
        maxAge: 31536000, // 1 year
        domain: configCookieDomain,
      });
    };

    return {
      toggleGridTableView: () => {
        setGridView((oldValue) => {
          // Compute new value
          const newVal = !oldValue;
          // Save in storage
          setCookie(newVal);

          return newVal;
        });
      },
      isGridViewEnabled: () => gridView,
    };
  }, [gridView, configCookieDomain, cookies]);

  return (
    <GridTableViewSwitcherContext.Provider value={tableGridViewSwitcherCtx}>
      {children}
    </GridTableViewSwitcherContext.Provider>
  );
}

export default GridTableViewSwitcherProvider;

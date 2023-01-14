import React, { ReactNode, useState, useMemo, useContext } from 'react';
import useMediaQuery from '@mui/material/useMediaQuery';
import { ThemeProvider as MuiThemeProvider, createTheme, ThemeOptions } from '@mui/material/styles';
import { PaletteMode } from '@mui/material';
import Cookies from 'universal-cookie';
import ColorModeContext from '../../../contexts/ColorModeContext';
import ConfigContext from '../../../contexts/ConfigContext';

interface Props {
  children: ReactNode;
  themeOptions: ThemeOptions;
}

function ThemeProvider({ children, themeOptions }: Props) {
  // Get config from context
  const cfg = useContext(ConfigContext);
  // Get cookies object
  const cookies = new Cookies();
  // Check prefer color scheme from system
  const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');
  // Get stored theme mode
  const storedThemeMode = cookies.get('palette-mode');
  // Compute initial value
  let initVal = storedThemeMode;
  if (initVal === null || (initVal !== 'dark' && initVal !== 'light')) {
    initVal = prefersDarkMode ? 'dark' : 'light';
  }

  // State for mode
  const [mode, setMode] = useState<PaletteMode>(initVal as PaletteMode);

  // Expand
  const { configCookieDomain } = cfg;

  // Create color mode context
  const colorMode = useMemo(() => {
    // Set cookie
    const setCookie = (input: PaletteMode) => {
      cookies.set('palette-mode', input, {
        path: '/',
        maxAge: 31536000, // 1 year
        domain: configCookieDomain,
      });
    };

    return {
      toggleColorMode: () => {
        setMode((prevMode) => {
          // Compute new value
          const newVal = prevMode === 'light' ? 'dark' : 'light';
          // Save in storage
          setCookie(newVal);

          return newVal;
        });
      },
      setColorMode: (input: PaletteMode | null) => {
        setMode((prevMode: PaletteMode) => {
          // Compute new value
          const newVal = input || prevMode;
          // Save in storage
          setCookie(newVal);

          return newVal;
        });
      },
    };
  }, []);

  const theme = useMemo(() => {
    // Initialize working copy
    const work = { ...themeOptions };
    // Patch palette mode from options
    // Check if palette object is present
    // Otherwise set it directly
    if (work.palette) {
      work.palette.mode = mode;
    } else {
      work.palette = { mode };
    }

    return createTheme(work);
  }, [mode, themeOptions]);

  return (
    <ColorModeContext.Provider value={colorMode}>
      <MuiThemeProvider theme={theme}>{children}</MuiThemeProvider>
    </ColorModeContext.Provider>
  );
}

export default ThemeProvider;

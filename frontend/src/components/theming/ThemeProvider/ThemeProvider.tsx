import React, { ReactNode, useState, useMemo } from 'react';
import useMediaQuery from '@mui/material/useMediaQuery';
import { ThemeProvider as MuiThemeProvider, createTheme, ThemeOptions } from '@mui/material/styles';
import { PaletteMode } from '@mui/material';
import ColorModeContext from '../../../contexts/ColorModeContext';

interface Props {
  children: ReactNode;
  themeOptions: ThemeOptions;
}

function ThemeProvider({ children, themeOptions }: Props) {
  // Check prefer color scheme from system
  const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');
  // Get stored theme mode
  const storedThemeMode = localStorage.getItem('paletteMode');
  // Compute initial value
  let initVal = storedThemeMode;
  if (initVal === null || (initVal !== 'dark' && initVal !== 'light')) {
    initVal = prefersDarkMode ? 'dark' : 'light';
  }

  // State for mode
  const [mode, setMode] = useState<PaletteMode>(initVal as PaletteMode);

  // Create color mode context
  const colorMode = useMemo(
    () => ({
      toggleColorMode: () => {
        setMode((prevMode) => {
          // Compute new value
          const newVal = prevMode === 'light' ? 'dark' : 'light';
          // Save in storage
          localStorage.setItem('paletteMode', newVal);

          return newVal;
        });
      },
    }),
    [],
  );

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

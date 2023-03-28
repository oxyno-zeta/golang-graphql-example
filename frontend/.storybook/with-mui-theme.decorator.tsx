import React, { useMemo } from 'react';
import { CssBaseline, ThemeProvider, createTheme } from '@mui/material';

const themes = {
  dark: createTheme({
    palette: {
      mode: 'dark',
    },
  }),
  light: createTheme({
    palette: {
      mode: 'light',
    },
  }),
};

export const withMuiTheme = (Story, context) => {
  const { theme: themeKey } = context.globals;

  // only recompute the theme if the themeKey changes
  const theme = useMemo(() => themes[themeKey] || themes['light'], [themeKey]);

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <Story />
    </ThemeProvider>
  );
};

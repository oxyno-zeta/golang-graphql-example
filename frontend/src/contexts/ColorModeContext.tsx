import { createContext } from 'react';
import { PaletteMode } from '@mui/material';

export default createContext({
  toggleColorMode: () => {},
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  setColorMode: (mode: PaletteMode | null) => {}, // This is just an empty function, ThemeProvider will set the real one
});

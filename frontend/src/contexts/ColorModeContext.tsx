import { createContext } from 'react';
import { PaletteMode } from '@mui/material';

export interface ColorModeContextModel {
  toggleColorMode: () => void;
  setColorMode: (mode: PaletteMode | null) => void;
}

export default createContext<ColorModeContextModel>({
  toggleColorMode: () => {},
  setColorMode: () => {},
});

import { createContext } from 'react';

export interface PageDrawerSettingsContextModel {
  toggleCollapsed: () => void;
  getCollapsed: () => boolean;
}

export default createContext<PageDrawerSettingsContextModel>({
  toggleCollapsed: () => {},
  getCollapsed: () => false,
});

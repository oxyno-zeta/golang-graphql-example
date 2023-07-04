import { createContext } from 'react';

export interface PageDrawerSettingsContextModel {
  toggleCollapsed: () => void;
  isCollapsed: () => boolean;
}

export default createContext<PageDrawerSettingsContextModel>({
  toggleCollapsed: () => {},
  isCollapsed: () => false,
});

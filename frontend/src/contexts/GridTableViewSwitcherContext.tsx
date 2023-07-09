import { createContext } from 'react';

export interface GridTableViewSwitcherContextModel {
  toggleGridTableView: () => void;
  isGridViewEnabled: () => boolean;
}

export default createContext<GridTableViewSwitcherContextModel>({
  toggleGridTableView: () => {},
  isGridViewEnabled: () => false,
});

import { createContext } from 'react';

export interface PageDrawerContextModel {
  onDrawerToggle: () => void;
}

export default createContext<PageDrawerContextModel>({
  onDrawerToggle: () => {},
});

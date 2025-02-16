import { createContext } from 'react';
import type { BreadcrumbData } from '../types';

export interface AutoBreadcrumbContextModel {
  pushAutoBreadcrumb: (input: BreadcrumbData) => void;
  popAutoBreadcrumb: () => void;
  getBreadcrumbData: () => BreadcrumbData[];
}

export default createContext<AutoBreadcrumbContextModel>({
  pushAutoBreadcrumb: () => {},
  popAutoBreadcrumb: () => {},
  getBreadcrumbData: () => [],
});

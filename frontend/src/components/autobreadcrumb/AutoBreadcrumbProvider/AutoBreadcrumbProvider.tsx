import React, { type ReactNode, useMemo, useState } from 'react';
import AutoBreadcrumbContext, { type AutoBreadcrumbContextModel } from '../contexts/AutoBreadcrumbContext';
import type { BreadcrumbData } from '../types';

interface Props {
  readonly children: ReactNode;
}

function AutoBreadcrumbProvider({ children }: Props) {
  const [state, setState] = useState<BreadcrumbData[]>([]);

  const contextValue: AutoBreadcrumbContextModel = useMemo(() => {
    return {
      pushAutoBreadcrumb: (input: BreadcrumbData) => {
        state.push(input);
        setState(state);
      },
      popAutoBreadcrumb: () => {
        state.pop();
        setState(state);
      },
      getBreadcrumbData: () => state,
    };
  }, [state, setState]);

  return <AutoBreadcrumbContext.Provider value={contextValue}>{children}</AutoBreadcrumbContext.Provider>;
}

export default AutoBreadcrumbProvider;

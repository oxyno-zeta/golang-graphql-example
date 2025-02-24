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
        if (!state.some((v: BreadcrumbData) => v.id === input.id)) {
          setState((s) => {
            return [...s, input].sort((a, b) => a.depth - b.depth);
          });
        }
      },
      popAutoBreadcrumb: (input: BreadcrumbData) => {
        setState((s) => s.filter((v) => v.id !== input.id));
      },
      getBreadcrumbData: () => state,
    };
  }, [state, setState]);

  return <AutoBreadcrumbContext.Provider value={contextValue}>{children}</AutoBreadcrumbContext.Provider>;
}

export default AutoBreadcrumbProvider;

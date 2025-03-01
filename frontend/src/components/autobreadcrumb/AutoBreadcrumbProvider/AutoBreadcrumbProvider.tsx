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
            const sorted = [...s, input].sort((a, b) => a.depth - b.depth);

            // Check that we aren't in production
            if (process.env.NODE_ENV !== 'production') {
              // Map all depth
              const allDepths = sorted.map((it) => it.depth);
              // Map all ids
              const allIds = sorted.map((it) => it.id);
              // Unique depth
              const uniqueDepths = allDepths.filter((value, index, array) => array.indexOf(value) === index);
              // Unique id
              const uniqueIds = allIds.filter((value, index, array) => array.indexOf(value) === index);

              // Check id uniqueness
              if (uniqueIds.length !== allIds.length) {
                console.error('Same id is provided more than once. Fix to have one id on one path');
              }

              // Check depth uniqueness
              if (uniqueDepths.length !== allDepths.length) {
                console.error('Same depth is provided more than once. Fix to have one depth on one path');
              }

              // Check that no depth is missing
              if (sorted[sorted.length - 1].depth !== sorted.length - 1) {
                console.error(
                  "One or many breadcrumb data seems to miss or last depth isn't the right one. (This error can arrive when multiple depth are crossed at the same time)",
                );
              }
            }

            return sorted;
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

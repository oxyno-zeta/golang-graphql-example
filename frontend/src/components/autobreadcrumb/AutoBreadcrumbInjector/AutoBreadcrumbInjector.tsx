import { type ReactNode, useContext, useEffect } from 'react';
import AutoBreadcrumbContext from '../contexts/AutoBreadcrumbContext';
import type { BreadcrumbData } from '../types';

interface Props {
  readonly item: BreadcrumbData;
  readonly children: ReactNode;
}

function AutoBreadcrumbInjector({ item, children }: Props) {
  const { pushAutoBreadcrumb, popAutoBreadcrumb } = useContext(AutoBreadcrumbContext);

  useEffect(() => {
    pushAutoBreadcrumb(item);

    return () => {
      popAutoBreadcrumb(item);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [item]);

  return children;
}

export default AutoBreadcrumbInjector;

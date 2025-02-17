import { type ReactNode, useContext, useEffect } from 'react';
import AutoBreadcrumbContext from '../contexts/AutoBreadcrumbContext';
import { BreadcrumbData } from '../types';

interface Props {
  readonly item: BreadcrumbData;
  readonly children: ReactNode;
}

function AutoBreadcrumbInjector({ item, children }: Props) {
  const { pushAutoBreadcrumb, popAutoBreadcrumb } = useContext(AutoBreadcrumbContext);

  useEffect(() => {
    pushAutoBreadcrumb(item);

    return () => {
      console.log('pop called in injector');
      popAutoBreadcrumb();
    };
  }, [item, popAutoBreadcrumb, pushAutoBreadcrumb]);

  return children;
}

export default AutoBreadcrumbInjector;

import { type ReactNode, useContext, useEffect } from 'react';
import AutoBreadcrumbContext from '../contexts/AutoBreadcrumbContext';
import { BreadcrumbData } from '../types';
import { useLocation, useResolvedPath } from 'react-router';

interface Props {
  readonly item: BreadcrumbData;
  readonly children: ReactNode;
}

function AutoBreadcrumbInjector({ item, children }: Props) {
  const { pushAutoBreadcrumb, popAutoBreadcrumb } = useContext(AutoBreadcrumbContext);
  // Get location data
  const locationData = useLocation();
  console.log(locationData);
  console.log(useResolvedPath(`../${locationData.pathname}/`));

  console.log(item);
  useEffect(() => {
    pushAutoBreadcrumb(item);

    return () => {
      console.log('pop called in injector', item);
      popAutoBreadcrumb(item);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [item]);

  return children;
}

export default AutoBreadcrumbInjector;

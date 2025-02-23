import React, { useContext } from 'react';
import { useLocation, useParams, useResolvedPath, resolvePath } from 'react-router';
import Breadcrumbs, { BreadcrumbsProps } from '@mui/material/Breadcrumbs';
import FixedBreadcrumb from './FixedBreadcrumb';
import GraphQLBreadcrumb from './GraphQLBreadcrumb';
import AutoBreadcrumbContext from '../contexts/AutoBreadcrumbContext';

export type Props = BreadcrumbsProps;

function AutoBreadcrumb(props: Props) {
  // Get context
  const ctx = useContext(AutoBreadcrumbContext);
  // Get params
  const params = useParams();
  // Get location data
  const locationData = useLocation();
  // console.log(locationData);
  const mm = useResolvedPath(`${locationData.pathname}/../`);
  console.log(mm);

  // Get breadcrumb data
  const crumbs = ctx.getBreadcrumbData();
  console.log(crumbs);

  return (
    <Breadcrumbs {...props}>
      {crumbs.map((breadcrumbData, i) => {
        // Compute last boolean
        const last = i === crumbs.length - 1;
        console.log('==========');
        console.log(build(crumbs.length - i));
        console.log(resolvePath(build(crumbs.length - i - 1), locationData.pathname));
        console.log('==========');

        if (breadcrumbData.fixed) {
          return (
            <FixedBreadcrumb
              breadcrumbData={breadcrumbData.fixed}
              key={breadcrumbData.id}
              last={last}
              pathname={resolvePath(build(crumbs.length - i - 1), locationData.pathname).pathname}
            />
          );
        }

        return (
          <GraphQLBreadcrumb
            breadcrumbData={breadcrumbData.graphql}
            key={breadcrumbData.id}
            last={last}
            params={params}
            pathname={locationData.pathname}
          />
        );
      })}
    </Breadcrumbs>
  );
}

function build(n: number) {
  let res = '';

  for (let i = 0; i < n; i++) {
    res += '../';
  }

  return res;
}

export default AutoBreadcrumb;

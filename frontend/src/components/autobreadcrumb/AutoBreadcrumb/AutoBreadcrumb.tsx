import React, { useContext } from 'react';
import { useLocation, useParams, resolvePath } from 'react-router';
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

  // Get breadcrumb data
  const crumbs = ctx.getBreadcrumbData();

  return (
    <Breadcrumbs {...props}>
      {crumbs.map((breadcrumbData, i) => {
        // Compute last boolean
        const last = i === crumbs.length - 1;
        // Compute pathname
        const computedPathname = resolvePath(
          build(crumbs.length - breadcrumbData.depth - 1),
          locationData.pathname,
        ).pathname;

        if (breadcrumbData.fixed) {
          return (
            <FixedBreadcrumb
              breadcrumbData={breadcrumbData.fixed}
              key={breadcrumbData.id}
              last={last}
              pathname={computedPathname}
            />
          );
        }

        return (
          <GraphQLBreadcrumb
            breadcrumbData={breadcrumbData.graphql}
            key={breadcrumbData.id}
            last={last}
            params={params}
            pathname={computedPathname}
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

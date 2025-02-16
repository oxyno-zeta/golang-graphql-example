import React, { useContext } from 'react';
import { useParams } from 'react-router-dom';
import Breadcrumbs, { BreadcrumbsProps } from '@mui/material/Breadcrumbs';
import type { BreadcrumbData } from '../types';
import FixedBreadcrumb from './FixedBreadcrumb';
import GraphQLBreadcrumb from './GraphQLBreadcrumb';
import AutoBreadcrumbContext from '../contexts/AutoBreadcrumbContext';

export type Props = BreadcrumbsProps;

function AutoBreadcrumb(props: Props) {
  // Get context
  const ctx = useContext(AutoBreadcrumbContext);
  // Get params
  const params = useParams();

  // Filter breadcrumbs
  const crumbs = matches
    // first get rid of any matches that don't have handle and crumb
    .filter((match) => Boolean((match.handle as RouteHandle)?.breadcrumb));

  return (
    <Breadcrumbs {...props}>
      {crumbs.map((y, i) => {
        // Compute last boolean
        const last = i === crumbs.length - 1;

        // Get route handle
        const routeHandle: RouteHandle = y.handle as RouteHandle;

        // Get data
        const breadcrumbData: BreadcrumbData = routeHandle.breadcrumb as BreadcrumbData;

        if (breadcrumbData.fixed) {
          return (
            <FixedBreadcrumb
              breadcrumbData={breadcrumbData.fixed}
              key={breadcrumbData.id || y.id}
              last={last}
              pathname={y.pathname}
            />
          );
        }

        return (
          <GraphQLBreadcrumb
            breadcrumbData={breadcrumbData.graphql}
            key={breadcrumbData.id || y.id}
            last={last}
            params={params}
            pathname={y.pathname}
          />
        );
      })}
    </Breadcrumbs>
  );
}

export default AutoBreadcrumb;

import React, { useContext } from 'react';
import { useLocation, useParams, useResolvedPath } from 'react-router';
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
  console.log(locationData);
  console.log(useResolvedPath(locationData.pathname));

  // Get breadcrumb data
  const crumbs = ctx.getBreadcrumbData();
  console.log(crumbs);

  return (
    <Breadcrumbs {...props}>
      {crumbs.map((breadcrumbData, i) => {
        // Compute last boolean
        const last = i === crumbs.length - 1;

        if (breadcrumbData.fixed) {
          return (
            <FixedBreadcrumb
              breadcrumbData={breadcrumbData.fixed}
              key={breadcrumbData.id}
              last={last}
              pathname={locationData.pathname}
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

export default AutoBreadcrumb;

import React from 'react';
import type { Params } from 'react-router-dom';
import { useQuery } from '@apollo/client';
import Skeleton from '@mui/material/Skeleton';
import type { BreadcrumbGraphQLDataConfig } from './types';
import FixedBreadcrumb from './FixedBreadcrumb';

interface Props {
  readonly params: Params<string>;
  readonly breadcrumbData: BreadcrumbGraphQLDataConfig;
  readonly last: boolean;
  readonly pathname: string;
}

function GraphQLBreadcrumb({ params, breadcrumbData, last, pathname }: Props) {
  // Build variables
  const variables = breadcrumbData.buildVariables ? breadcrumbData.buildVariables(params) : {};

  // Query
  const { data, loading, error } = useQuery(breadcrumbData.query, {
    variables,
    fetchPolicy: 'cache-first',
    ...breadcrumbData.queryOptions,
  });

  // Check loading or error
  if (error || loading) {
    // Check error
    if (error) {
      // Log
      console.error(error);
    }

    return <Skeleton variant="text" width={40} {...(breadcrumbData.skeletonProps || {})} />;
  }

  // Get text
  const text = breadcrumbData.getTextContent(data);

  return (
    <FixedBreadcrumb
      breadcrumbData={{
        textContent: text,
        linkProps: breadcrumbData.linkProps,
        typographyProps: breadcrumbData.typographyProps,
      }}
      last={last}
      pathname={pathname}
    />
  );
}

export default GraphQLBreadcrumb;

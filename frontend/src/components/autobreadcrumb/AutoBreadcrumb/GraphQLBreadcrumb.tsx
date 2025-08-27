import React from 'react';
import type { Params } from 'react-router';
import { useQuery } from '@apollo/client/react';
import Skeleton from '@mui/material/Skeleton';
import { useTranslation } from 'react-i18next';
import type { BreadcrumbGraphQLDataConfig } from '../types';
import FixedBreadcrumb from './FixedBreadcrumb';

interface Props {
  readonly params: Params;
  readonly breadcrumbData: BreadcrumbGraphQLDataConfig;
  readonly last: boolean;
  readonly pathname: string;
  readonly disablePageTitle?: boolean;
}

function GraphQLBreadcrumb({ params, breadcrumbData, last, pathname, disablePageTitle = false }: Props) {
  // Translate
  const { t } = useTranslation();

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
  const text = breadcrumbData.getTextContent(data, t);

  // Create override computed path function if necessary
  let overrideComputedPath: ((p: string, params: Params) => string) | undefined;
  if (breadcrumbData.overrideComputedPath) {
    overrideComputedPath = (p: string, inputParams: Params) =>
      breadcrumbData.overrideComputedPath!(p, inputParams, data);
  }

  return (
    <FixedBreadcrumb
      breadcrumbData={{
        disableTranslate: true,
        textContent: text,
        linkProps: breadcrumbData.linkProps,
        typographyProps: breadcrumbData.typographyProps,
        overrideComputedPath,
      }}
      last={last}
      pathname={pathname}
      disablePageTitle={disablePageTitle}
      params={params}
    />
  );
}

export default GraphQLBreadcrumb;

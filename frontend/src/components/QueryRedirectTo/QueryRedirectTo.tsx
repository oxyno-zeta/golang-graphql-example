import React from 'react';
import { useTranslation } from 'react-i18next';
import { type DocumentNode, type OperationVariables, type MaybeMasked } from '@apollo/client';
import { useQuery } from '@apollo/client/react';
import { useParams, type Params, Navigate, type NavigateProps } from 'react-router';
import CenterLoading, { type CenterLoadingProps } from '~components/CenterLoading';
import ErrorsDisplay, { type ErrorsDisplayProps } from '~components/ErrorsDisplay';
import NoData, { type NoDataTypographyProps } from '~components/NoData';

export interface Props<T, P extends OperationVariables> {
  // Query document to execute.
  readonly query: DocumentNode;
  // Build navigate to function will return a path.
  // If path result is empty or undefined, then "no data" is displayed.
  // If you don't want any "no data" displayed, manage the redirect path
  readonly buildNavigateTo: (params: MaybeMasked<T> | undefined) => string | undefined | null;
  // Build query variables function will return a query variable object.
  readonly buildQueryVariables: (params: Params) => P;
  // Disable center loading subtitle.
  readonly disableCenterLoadingSubtitle?: boolean;
  // Query hook options.
  readonly queryHookOptions?: Omit<useQuery.Options<T, P>, 'variables'>;
  // No data Typography props.
  readonly noDataTypographyProps?: NoDataTypographyProps;
  // Center loading props.
  readonly centerLoadingProps?: CenterLoadingProps;
  // Graphql Errors props.
  readonly graphqlErrorsProps?: Omit<ErrorsDisplayProps, 'error|errors'>;
  // Navigate props.
  readonly navigateProps?: Omit<NavigateProps, 'to'>;
}

function QueryRedirectTo<T, P extends OperationVariables>({
  query,
  buildQueryVariables,
  buildNavigateTo,
  disableCenterLoadingSubtitle = false,
  queryHookOptions = undefined,
  noDataTypographyProps = {},
  centerLoadingProps = {},
  graphqlErrorsProps = {},
  navigateProps = {},
}: Props<T, P>) {
  // Setup translate
  const { t } = useTranslation();
  // Get params
  const params = useParams();

  // Build query variables
  const queryVariables: P = buildQueryVariables(params);

  // Query
  const { data, loading, error } = useQuery<T, P>(query, {
    variables: queryVariables,
    fetchPolicy: 'network-only',
    ...(queryHookOptions || {}),
  });

  // Check loading
  if (loading) {
    return (
      <CenterLoading
        containerBoxSx={{ margin: '15px 0' }}
        subtitle={!disableCenterLoadingSubtitle ? t('common.loadingText') : undefined}
        {...centerLoadingProps}
      />
    );
  }

  // Check error
  if (error) {
    return <ErrorsDisplay error={error} {...graphqlErrorsProps} />;
  }

  // Build navigate to
  const to = buildNavigateTo(data);

  // Check if to isn't present
  if (!to) {
    return <NoData {...noDataTypographyProps} />;
  }

  return <Navigate to={to} {...navigateProps} />;
}

export default QueryRedirectTo;

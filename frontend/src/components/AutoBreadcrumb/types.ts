import type { QueryHookOptions } from '@apollo/client';
import type { LinkProps, SkeletonProps, TypographyProps } from '@mui/material';
import type { DocumentNode } from 'graphql';
import type { LinkProps as RouterLinkProps, Params } from 'react-router-dom';

export type RouteHandle = {
  breadcrumb?: BreadcrumbData;
};

export type BreadcrumbData = BreadcrumbGraphQLData | BreadcrumbFixedData;

export type BreadcrumbFixedData = {
  // Id is needed for each element to ensure to be unique.
  id: string;
  fixed: BreadcrumbFixedDataConfig;
  graphql: undefined;
};

export type BreadcrumbGraphQLData = {
  // Id is needed for each element to ensure to be unique.
  id: string;
  fixed: undefined;
  graphql: BreadcrumbGraphQLDataConfig;
};

export type BreadcrumbFixedDataConfig = {
  textContent: string;
  linkProps?: Omit<LinkProps & RouterLinkProps, 'to'>;
  typographyProps?: TypographyProps;
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type BreadcrumbGraphQLDataConfig<V = Record<string, any>, D = any> = {
  query: DocumentNode;
  getTextContent: (data: D) => string;
  queryOptions?: Omit<QueryHookOptions<D>, 'variables'>;
  buildVariables?: (params: Params<string>) => V;
  skeletonProps?: SkeletonProps;
  linkProps?: Omit<LinkProps & RouterLinkProps, 'to'>;
  typographyProps?: TypographyProps;
};

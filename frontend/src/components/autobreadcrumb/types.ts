import type { QueryHookOptions } from '@apollo/client';
import type { LinkProps, SkeletonProps, TypographyProps } from '@mui/material';
import type { DocumentNode } from 'graphql';
import type { LinkProps as RouterLinkProps, Params } from 'react-router';

export type BreadcrumbData = BreadcrumbGraphQLData | BreadcrumbFixedData | BreadcrumbDataIgnoredRoute;

export interface BreadcrumbDataIgnoredRoute {
  // Id is needed for each element to ensure to be unique.
  id: string;
  // Depth to compute routes
  // Start at 0
  depth: number;
  ignored: boolean;
  fixed?: undefined;
  graphql?: undefined;
}

export interface BreadcrumbFixedData {
  // Id is needed for each element to ensure to be unique.
  id: string;
  // Depth to compute routes
  // Start at 0
  depth: number;
  fixed: BreadcrumbFixedDataConfig;
  graphql?: undefined;
  ignored?: undefined;
}

export interface BreadcrumbGraphQLData {
  // Id is needed for each element to ensure to be unique.
  id: string;
  // Depth to compute routes
  // Start at 0
  depth: number;
  fixed?: undefined;
  ignored?: undefined;
  graphql: BreadcrumbGraphQLDataConfig;
}

export interface BreadcrumbFixedDataConfig {
  textContent: string;
  linkProps?: Omit<LinkProps & RouterLinkProps, 'to'>;
  typographyProps?: TypographyProps;
  overrideComputedPath?: (computedPath: string, params: Params) => string;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export interface BreadcrumbGraphQLDataConfig<V = Record<string, any>, D = any> {
  query: DocumentNode;
  getTextContent: (data: D) => string;
  queryOptions?: Omit<QueryHookOptions<D>, 'variables'>;
  buildVariables?: (params: Params) => V;
  skeletonProps?: SkeletonProps;
  linkProps?: Omit<LinkProps & RouterLinkProps, 'to'>;
  typographyProps?: TypographyProps;
  overrideComputedPath?: (computedPath: string, params: Params, data: D) => string;
}

import React from 'react';
import type { Params, RouteObject } from 'react-router-dom';
import { gql } from '@apollo/client';
import NotFoundRoute from '~components/NotFoundRoute';
import MainContentWrapper from '~components/MainContentWrapper';
import Todos from '~routes/Todos';
import QueryRedirectTo from '~components/QueryRedirectTo';
import type { TodoModel } from '~models/todos';
import type { ConnectionModel } from '~models/general';

const QUERY = gql`
  query Q($name: String) {
    todos(filter: { text: { contains: $name } }) {
      edges {
        node {
          id
        }
      }
    }
  }
`;

interface QueryResult {
  todos: ConnectionModel<TodoModel>;
}

interface QueryVariables {
  name: string;
}

const buildNavigateTo = (params?: QueryResult) =>
  params?.todos?.edges && params?.todos?.edges[0] && `/fake/${params?.todos?.edges[0].node.id}`;

const buildQueryVariables = (params: Params<string>) => ({ name: params.name as string });

const router: RouteObject[] = [
  {
    path: '/',
    index: true,
    element: (
      <MainContentWrapper>
        <Todos />
      </MainContentWrapper>
    ),
  },
  {
    path: 'redirect-to/:name',
    element: (
      <MainContentWrapper>
        <QueryRedirectTo<QueryResult, QueryVariables>
          buildNavigateTo={buildNavigateTo}
          buildQueryVariables={buildQueryVariables}
          query={QUERY}
        />
      </MainContentWrapper>
    ),
  },
  {
    path: '*',
    element: (
      <MainContentWrapper>
        <NotFoundRoute />
      </MainContentWrapper>
    ),
  },
];

export default router;

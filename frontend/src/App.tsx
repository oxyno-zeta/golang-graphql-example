import React from 'react';
import { Routes, Route } from 'react-router-dom';
import { gql } from '@apollo/client';
import NotFoundRoute from '~components/NotFoundRoute';
import MainContentWrapper from '~components/MainContentWrapper';
import Todos from '~routes/Todos';
import QueryRedirectTo from '~components/QueryRedirectTo';
import { TodoModel } from '~models/todos';
import { ConnectionModel } from '~models/general';

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

function App() {
  return (
    <Routes>
      <Route path="/">
        <Route
          index
          element={
            <MainContentWrapper>
              <Todos />
            </MainContentWrapper>
          }
        />
        <Route
          path="/redirect-to/:name"
          element={
            <MainContentWrapper>
              <QueryRedirectTo<QueryResult, QueryVariables>
                query={QUERY}
                buildNavigateTo={(params) =>
                  params?.todos?.edges && params?.todos?.edges[0] && `/fake/${params?.todos?.edges[0].node.id}`
                }
                buildQueryVariables={(params) => ({ name: params.name as string })}
              />
            </MainContentWrapper>
          }
        />
        <Route
          path="*"
          element={
            <MainContentWrapper>
              <NotFoundRoute />
            </MainContentWrapper>
          }
        />
      </Route>
    </Routes>
  );
}

export default App;

import React, { type ReactNode } from 'react';
import { MemoryRouter, Route, Routes } from 'react-router';
import { render, waitFor } from '@testing-library/react';
import { ApolloError, gql } from '@apollo/client';
import { MockedProvider, type MockedResponse } from '@apollo/client/testing';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import { GraphQLError } from 'graphql';
import QueryRedirectTo from './QueryRedirectTo';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

const QUERY = gql`
  query Q($name: String) {
    fake(name: $name) {
      id
    }
  }
`;

interface Props {
  readonly children: ReactNode;
  readonly mockedResponse: MockedResponse;
}

function TestComponent({ children, mockedResponse }: Props) {
  return (
    <MockedProvider addTypename={false} mocks={[mockedResponse]}>
      <MemoryRouter initialEntries={['/route/fake-name']}>
        <Routes>
          <Route element={<div>Fake</div>} path="/fake" />
          <Route element={<>{children}</>} path="/route/:name" />
        </Routes>
      </MemoryRouter>
    </MockedProvider>
  );
}

describe('QueryRedirectTo', () => {
  it('should display a loading when query is pending', async () => {
    const { container, findByRole } = render(
      <TestComponent
        mockedResponse={{
          request: {
            query: QUERY,
            variables: { name: 'fake-name' },
          },
        }}
      >
        <QueryRedirectTo buildNavigateTo={() => '/fake'} buildQueryVariables={({ name }) => ({ name })} query={QUERY} />
      </TestComponent>,
    );

    expect(container).toMatchSnapshot();

    // Find progressbar
    expect(container).toHaveTextContent('common.loadingText');
    expect(await findByRole('progressbar')).not.toBeNull();
  });

  it('should display a loading without subtitle when query is pending', async () => {
    const { container, findByRole } = render(
      <TestComponent
        mockedResponse={{
          request: {
            query: QUERY,
            variables: { name: 'fake-name' },
          },
        }}
      >
        <QueryRedirectTo
          buildNavigateTo={() => '/fake'}
          buildQueryVariables={({ name }) => ({ name })}
          disableCenterLoadingSubtitle
          query={QUERY}
        />
      </TestComponent>,
    );

    expect(container).toMatchSnapshot();

    // Find progressbar
    expect(container).not.toHaveTextContent('common.loadingText');
    expect(await findByRole('progressbar')).not.toBeNull();
  });

  it('should display an error when an error is raised on query', async () => {
    const { container, findByRole } = render(
      <TestComponent
        mockedResponse={{
          request: {
            query: QUERY,
            variables: { name: 'fake-name' },
          },
          error: new ApolloError({
            graphQLErrors: [
              new GraphQLError('forbidden graphql error', {
                extensions: { code: 'FORBIDDEN' },
              }),
            ],
          }),
        }}
      >
        <QueryRedirectTo buildNavigateTo={() => '/fake'} buildQueryVariables={({ name }) => ({ name })} query={QUERY} />
      </TestComponent>,
    );

    expect(container).toMatchSnapshot();

    // Find progressbar
    expect(await findByRole('progressbar')).not.toBeNull();

    // Now find errors
    expect(container).toHaveTextContent('common.errors');
    expect(container).toMatchSnapshot();
  });

  it('should display no data when no data is returned from query', async () => {
    const { container, findByRole } = render(
      <TestComponent
        mockedResponse={{
          request: {
            query: QUERY,
            variables: { name: 'fake-name' },
          },
          result: {
            data: null,
          },
        }}
      >
        <QueryRedirectTo buildNavigateTo={() => null} buildQueryVariables={({ name }) => ({ name })} query={QUERY} />
      </TestComponent>,
    );

    expect(container).toMatchSnapshot();

    // Find progressbar
    expect(await findByRole('progressbar')).not.toBeNull();

    // Now find errors
    expect(container).toHaveTextContent('common.noData');
    expect(container).toMatchSnapshot();
  });

  it('should redirect to new path on success', async () => {
    const { container, findByRole } = render(
      <TestComponent
        mockedResponse={{
          request: {
            query: QUERY,
            variables: { name: 'fake-name' },
          },
          result: {
            data: { fake: { id: '1' } },
          },
        }}
      >
        <QueryRedirectTo buildNavigateTo={() => '/fake'} buildQueryVariables={({ name }) => ({ name })} query={QUERY} />
      </TestComponent>,
    );

    expect(container).toMatchSnapshot();

    // Find progressbar
    expect(await findByRole('progressbar')).not.toBeNull();

    // Workaround to avoid "react component change without any act called"...
    await waitFor(() => 0);

    // Now find errors
    expect(container).toHaveTextContent('Fake');
    expect(container).toMatchSnapshot();
  });
});

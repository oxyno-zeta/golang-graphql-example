import React, { type ReactNode } from 'react';
import { MemoryRouter, Route, Routes } from 'react-router';
import { render, waitFor } from '@testing-library/react';
import { CombinedGraphQLErrors, gql } from '@apollo/client';
import { type MockLink } from '@apollo/client/testing';
import { MockedProvider } from '@apollo/client/testing/react';
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
  readonly mockedResponse: MockLink.MockedResponse;
}

function TestComponent({ children, mockedResponse }: Props) {
  return (
    <MockedProvider mocks={[mockedResponse]}>
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
          delay: Infinity,
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
          delay: Infinity,
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
          error: new CombinedGraphQLErrors({
            errors: [
              new GraphQLError('forbidden graphql error', {
                extensions: { code: 'FORBIDDEN' },
              }),
            ],
          }),
          delay: 0,
        }}
      >
        <QueryRedirectTo buildNavigateTo={() => '/fake'} buildQueryVariables={({ name }) => ({ name })} query={QUERY} />
      </TestComponent>,
    );

    expect(container).toMatchSnapshot();

    await waitFor(async () => {
      expect(await findByRole('progressbar')).not.toBeInTheDocument();
    });

    // Now find errors
    expect(container).toMatchSnapshot();
    expect(container).toHaveTextContent('common.errors');
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
          delay: 0,
        }}
      >
        <QueryRedirectTo buildNavigateTo={() => null} buildQueryVariables={({ name }) => ({ name })} query={QUERY} />
      </TestComponent>,
    );

    expect(container).toMatchSnapshot();

    await waitFor(async () => {
      expect(await findByRole('progressbar')).not.toBeInTheDocument();
    });

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
            data: { fake: { id: '1', __typename: 'Fake' } },
          },
          delay: 0,
        }}
      >
        <QueryRedirectTo buildNavigateTo={() => '/fake'} buildQueryVariables={({ name }) => ({ name })} query={QUERY} />
      </TestComponent>,
    );

    expect(container).toMatchSnapshot();

    await waitFor(async () => {
      expect(await findByRole('progressbar')).not.toBeInTheDocument();
    });

    // Workaround to avoid "react component change without any act called"...
    await waitFor(() => 0);

    // Now find errors
    expect(container).toHaveTextContent('Fake');
    expect(container).toMatchSnapshot();
  });
});

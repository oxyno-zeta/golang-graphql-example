import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import { Link, Outlet, RouterProvider, createMemoryRouter } from 'react-router-dom';
import AutoBreadcrumb from './AutoBreadcrumb';
import {
  SimpleErrorQuery,
  SimpleQuery1,
  SimpleQuery2,
  SlowQuery,
  mockedResponses,
} from './AutoBreadcrumb.storage-test';
import type { RouteHandle } from './types';

export default {
  title: 'Components/AutoBreadcrumb',
  component: AutoBreadcrumb,
} as Meta<typeof AutoBreadcrumb>;

const allFixedRoutes = [
  {
    path: '/',
    element: (
      <>
        <div style={{ marginBottom: '30px' }}>
          Menu:
          <br />
          <Link to="/">/</Link>
          <br />
          <Link to="/level1">/level1</Link>
          <br />
          <Link to="/level1/level2">/level1/level2</Link>
        </div>
        Component:
        <AutoBreadcrumb />
        <Outlet />
      </>
    ),
    handle: {
      breadcrumb: { id: 'fake-id-1', fixed: { textContent: 'fake' } },
    } as RouteHandle,
    children: [
      {
        path: 'level1',
        element: (
          <>
            <div />
            <Outlet />
          </>
        ),
        handle: {
          breadcrumb: { id: 'fake-id-2', fixed: { textContent: 'fake' } },
        } as RouteHandle,
        children: [
          {
            path: 'level2',
            element: <div />,
            handle: {
              breadcrumb: { id: 'fake-id-3', fixed: { textContent: 'fake' } },
            } as RouteHandle,
          },
        ],
      },
    ],
  },
];

export const AllFixed: StoryFn<typeof AutoBreadcrumb> = function C() {
  const router = createMemoryRouter(allFixedRoutes, { initialEntries: ['/level1/level2'], initialIndex: 0 });

  return <RouterProvider router={router} />;
};

const allGraphqlRoutes = [
  {
    path: '/',
    element: (
      <>
        <div style={{ marginBottom: '30px' }}>
          Menu:
          <br />
          <Link to="/">/</Link>
          <br />
          <Link to="/level1">/level1</Link>
          <br />
          <Link to="/level1/error">/level1/error</Link>
          <br />
          <Link to="/level1/slow">/level1/slow</Link>
        </div>
        Component:
        <AutoBreadcrumb />
        <Outlet />
      </>
    ),
    handle: {
      breadcrumb: { id: 'fake-id-1', graphql: { query: SimpleQuery1, getTextContent: (data) => data.name } },
    } as RouteHandle,
    children: [
      {
        path: 'level1',
        element: (
          <>
            <div />
            <Outlet />
          </>
        ),
        handle: {
          breadcrumb: { id: 'fake-id-2', graphql: { query: SimpleQuery2, getTextContent: (data) => data.name2 } },
        } as RouteHandle,
        children: [
          {
            path: 'error',
            element: <div />,
            handle: {
              breadcrumb: { id: 'fake-id-3', graphql: { query: SimpleErrorQuery } },
            } as RouteHandle,
          },
          {
            path: 'slow',
            element: <div />,
            handle: {
              breadcrumb: { id: 'fake-id-4', graphql: { query: SlowQuery, getTextContent: (data) => data.slow } },
            } as RouteHandle,
          },
        ],
      },
    ],
  },
];

export const AllGraphql: StoryFn<typeof AutoBreadcrumb> = function C() {
  const router = createMemoryRouter(allGraphqlRoutes, { initialEntries: ['/level1'], initialIndex: 0 });

  return <RouterProvider router={router} />;
};
AllGraphql.parameters = {
  apolloClient: {
    // Example coming from https://storybook.js.org/addons/storybook-addon-apollo-client
    mocks: mockedResponses,
  },
};

const allMixedRoutes = [
  {
    path: '/',
    element: (
      <>
        <div style={{ marginBottom: '30px' }}>
          Menu:
          <br />
          <Link to="/">/</Link>
          <br />
          <Link to="/level1">/level1</Link>
          <br />
          <Link to="/level1/level2">/level1/level2</Link>
        </div>
        Component:
        <AutoBreadcrumb />
        <Outlet />
      </>
    ),
    handle: {
      breadcrumb: { id: 'fake-id-1', fixed: { textContent: 'fake' } },
    } as RouteHandle,
    children: [
      {
        path: 'level1',
        element: (
          <>
            <div />
            <Outlet />
          </>
        ),
        handle: {
          breadcrumb: { id: 'fake-id-2', graphql: { query: SlowQuery, getTextContent: (data) => data.slow } },
        } as RouteHandle,
        children: [
          {
            path: 'level2',
            element: <div />,
            handle: {
              breadcrumb: { id: 'fake-id-3', fixed: { textContent: 'fake' } },
            } as RouteHandle,
          },
        ],
      },
    ],
  },
];

export const AllMixed: StoryFn<typeof AutoBreadcrumb> = function C() {
  const router = createMemoryRouter(allMixedRoutes, { initialEntries: ['/level1/level2'], initialIndex: 0 });

  return <RouterProvider router={router} />;
};
AllMixed.parameters = {
  apolloClient: {
    // Example coming from https://storybook.js.org/addons/storybook-addon-apollo-client
    mocks: mockedResponses,
  },
};

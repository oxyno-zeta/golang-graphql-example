import React from 'react';
import { Outlet, createMemoryRouter, RouterProvider } from 'react-router-dom';
import { render, screen } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import { MockedProvider } from '@apollo/client/testing';
import type { RouteHandle } from './types';
import {
  SimpleErrorQuery,
  SimpleQuery1,
  SimpleQuery2,
  SlowQuery,
  mockedResponses,
} from './AutoBreadcrumb.storage-test';

import AutoBreadcrumb from './AutoBreadcrumb';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('AutoBreadcrumb', () => {
  describe('Fixed texts', () => {
    const allFixedRoutes = [
      {
        path: '/',
        element: (
          <>
            <AutoBreadcrumb />
            <Outlet />
          </>
        ),
        handle: {
          breadcrumb: { id: 'fake-id-1', fixed: { textContent: 'root' } },
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
              breadcrumb: { id: 'fake-id-2', fixed: { textContent: 'level1' } },
            } as RouteHandle,
            children: [
              {
                path: 'level2',
                element: <div />,
                handle: {
                  breadcrumb: { id: 'fake-id-3', fixed: { textContent: 'level2' } },
                } as RouteHandle,
              },
            ],
          },
        ],
      },
    ];

    it('should be ok to display on 1 level', () => {
      const router = createMemoryRouter(allFixedRoutes, { initialEntries: ['/'] });

      const { container } = render(<RouterProvider router={router} />);

      expect(container).toMatchSnapshot();

      expect(container).toHaveTextContent('root');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(1);

      expect(liElements[0].children[0].localName).toEqual('p');
      expect(liElements[0].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements[0].children[0]).toHaveTextContent('root');
    });

    it('should be ok to display on 2 levels', () => {
      const router = createMemoryRouter(allFixedRoutes, { initialEntries: ['/level1/'] });

      const { container } = render(<RouterProvider router={router} />);

      expect(container).toMatchSnapshot();

      expect(container).toHaveTextContent('root');
      expect(container).toHaveTextContent('/');
      expect(container).toHaveTextContent('level1');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(3);

      expect(liElements[0].children[0].localName).toEqual('a');
      expect(liElements[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements[0].children[0]).toHaveTextContent('root');
      expect(liElements[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements[1].localName).toEqual('li');
      expect(liElements[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[1]).toHaveTextContent('/');

      expect(liElements[2].children[0].localName).toEqual('p');
      expect(liElements[2].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements[2].children[0]).toHaveTextContent('level1');
    });

    it('should be ok to display on 3 levels', () => {
      const router = createMemoryRouter(allFixedRoutes, { initialEntries: ['/level1/level2/'] });

      const { container } = render(<RouterProvider router={router} />);

      expect(container).toMatchSnapshot();

      expect(container).toHaveTextContent('root');
      expect(container).toHaveTextContent('/');
      expect(container).toHaveTextContent('level1');
      expect(container).toHaveTextContent('level2');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(5);

      expect(liElements[0].children[0].localName).toEqual('a');
      expect(liElements[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements[0].children[0]).toHaveTextContent('root');
      expect(liElements[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements[1].localName).toEqual('li');
      expect(liElements[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[1]).toHaveTextContent('/');

      expect(liElements[2].children[0].localName).toEqual('a');
      expect(liElements[2].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements[2].children[0]).toHaveTextContent('level1');
      expect(liElements[2].children[0]).toHaveAttribute('href', '/level1');

      expect(liElements[3].localName).toEqual('li');
      expect(liElements[3]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[3]).toHaveTextContent('/');

      expect(liElements[4].children[0].localName).toEqual('p');
      expect(liElements[4].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements[4].children[0]).toHaveTextContent('level2');
    });
  });

  describe('GraphQL texts', () => {
    const allGraphqlRoutes = [
      {
        path: '/',
        element: (
          <>
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

    it('should be ok to display on 1 level', async () => {
      const router = createMemoryRouter(allGraphqlRoutes, { initialEntries: ['/'] });

      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <RouterProvider router={router} />
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      expect(container).not.toHaveTextContent('Query1');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(1);

      expect(liElements[0].children[0].localName).toEqual('span');
      expect(liElements[0].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');

      expect(await screen.findByText('Query1')).toBeInTheDocument();

      expect(container).toMatchSnapshot();

      const navElement2 = container.querySelector('nav');
      expect(navElement2).not.toBeNull();
      const olElement2 = container.querySelector('ol');
      expect(olElement2).not.toBeNull();
      const liElements2 = container.querySelectorAll('li');
      expect(liElements2).not.toBeNull();
      expect(liElements2).toHaveLength(1);

      expect(liElements2[0].children[0].localName).toEqual('p');
      expect(liElements2[0].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements2[0].children[0]).toHaveTextContent('Query1');
    });

    it('should be ok to display on 2 levels', async () => {
      const router = createMemoryRouter(allGraphqlRoutes, { initialEntries: ['/level1/'] });

      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <RouterProvider router={router} />
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      expect(container).not.toHaveTextContent('Query1');
      expect(container).toHaveTextContent('/');
      expect(container).not.toHaveTextContent('Query2');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(3);

      expect(liElements[0].children[0].localName).toEqual('span');
      expect(liElements[0].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');

      expect(liElements[1].localName).toEqual('li');
      expect(liElements[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[1]).toHaveTextContent('/');

      expect(liElements[2].children[0].localName).toEqual('span');
      expect(liElements[2].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');

      expect(await screen.findByText('Query1')).toBeInTheDocument();
      expect(await screen.findByText('Query2')).toBeInTheDocument();

      expect(container).toMatchSnapshot();

      const navElement2 = container.querySelector('nav');
      expect(navElement2).not.toBeNull();
      const olElement2 = container.querySelector('ol');
      expect(olElement2).not.toBeNull();
      const liElements2 = container.querySelectorAll('li');
      expect(liElements2).not.toBeNull();
      expect(liElements2).toHaveLength(3);

      expect(liElements2[0].children[0].localName).toEqual('a');
      expect(liElements2[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements2[0].children[0]).toHaveTextContent('Query1');
      expect(liElements2[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements2[1].localName).toEqual('li');
      expect(liElements2[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements2[1]).toHaveTextContent('/');

      expect(liElements2[2].children[0].localName).toEqual('p');
      expect(liElements2[2].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements2[2].children[0]).toHaveTextContent('Query2');
    });

    it('should be ok to display on 2 levels and third with an error', async () => {
      const router = createMemoryRouter(allGraphqlRoutes, { initialEntries: ['/level1/error'] });

      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <RouterProvider router={router} />
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      expect(container).not.toHaveTextContent('Query1');
      expect(container).toHaveTextContent('/');
      expect(container).not.toHaveTextContent('Query2');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(5);

      expect(liElements[0].children[0].localName).toEqual('span');
      expect(liElements[0].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');

      expect(liElements[1].localName).toEqual('li');
      expect(liElements[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[1]).toHaveTextContent('/');

      expect(liElements[2].children[0].localName).toEqual('span');
      expect(liElements[2].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');

      expect(liElements[3].localName).toEqual('li');
      expect(liElements[3]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[3]).toHaveTextContent('/');

      expect(liElements[4].children[0].localName).toEqual('span');
      expect(liElements[4].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');

      expect(await screen.findByText('Query1')).toBeInTheDocument();
      expect(await screen.findByText('Query2')).toBeInTheDocument();

      expect(container).toMatchSnapshot();

      const navElement2 = container.querySelector('nav');
      expect(navElement2).not.toBeNull();
      const olElement2 = container.querySelector('ol');
      expect(olElement2).not.toBeNull();
      const liElements2 = container.querySelectorAll('li');
      expect(liElements2).not.toBeNull();
      expect(liElements2).toHaveLength(5);

      expect(liElements2[0].children[0].localName).toEqual('a');
      expect(liElements2[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements2[0].children[0]).toHaveTextContent('Query1');
      expect(liElements2[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements2[1].localName).toEqual('li');
      expect(liElements2[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements2[1]).toHaveTextContent('/');

      expect(liElements2[2].children[0].localName).toEqual('a');
      expect(liElements2[2].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements2[2].children[0]).toHaveTextContent('Query2');
      expect(liElements2[2].children[0]).toHaveAttribute('href', '/level1');

      expect(liElements[3].localName).toEqual('li');
      expect(liElements[3]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[3]).toHaveTextContent('/');

      expect(liElements[4].children[0].localName).toEqual('span');
      expect(liElements[4].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');
    });
  });

  describe('Mixed texts', () => {
    const allMixedRoutes = [
      {
        path: '/',
        element: (
          <>
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
              breadcrumb: { id: 'fake-id-2', graphql: { query: SimpleQuery2, getTextContent: (data) => data.name2 } },
            } as RouteHandle,
          },
        ],
      },
    ];

    it('should be ok to display on 1 level', async () => {
      const router = createMemoryRouter(allMixedRoutes, { initialEntries: ['/'] });

      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <RouterProvider router={router} />
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      const navElement2 = container.querySelector('nav');
      expect(navElement2).not.toBeNull();
      const olElement2 = container.querySelector('ol');
      expect(olElement2).not.toBeNull();
      const liElements2 = container.querySelectorAll('li');
      expect(liElements2).not.toBeNull();
      expect(liElements2).toHaveLength(1);

      expect(liElements2[0].children[0].localName).toEqual('p');
      expect(liElements2[0].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements2[0].children[0]).toHaveTextContent('fake');
    });

    it('should be ok to display on 2 levels', async () => {
      const router = createMemoryRouter(allMixedRoutes, { initialEntries: ['/level1/'] });

      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <RouterProvider router={router} />
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      expect(container).toHaveTextContent('fake');
      expect(container).toHaveTextContent('/');
      expect(container).not.toHaveTextContent('Query2');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(3);

      expect(liElements[0].children[0].localName).toEqual('a');
      expect(liElements[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements[0].children[0]).toHaveTextContent('fake');
      expect(liElements[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements[1].localName).toEqual('li');
      expect(liElements[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[1]).toHaveTextContent('/');

      expect(liElements[2].children[0].localName).toEqual('span');
      expect(liElements[2].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');

      expect(await screen.findByText('Query2')).toBeInTheDocument();

      expect(container).toMatchSnapshot();

      const navElement2 = container.querySelector('nav');
      expect(navElement2).not.toBeNull();
      const olElement2 = container.querySelector('ol');
      expect(olElement2).not.toBeNull();
      const liElements2 = container.querySelectorAll('li');
      expect(liElements2).not.toBeNull();
      expect(liElements2).toHaveLength(3);

      expect(liElements2[0].children[0].localName).toEqual('a');
      expect(liElements2[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements2[0].children[0]).toHaveTextContent('fake');
      expect(liElements2[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements2[1].localName).toEqual('li');
      expect(liElements2[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements2[1]).toHaveTextContent('/');

      expect(liElements2[2].children[0].localName).toEqual('p');
      expect(liElements2[2].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements2[2].children[0]).toHaveTextContent('Query2');
    });
  });

  describe('Mixed texts and ignored routes', () => {
    const routes = [
      {
        path: '/',
        element: (
          <>
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
            children: [
              {
                path: 'level2',
                element: (
                  <>
                    <div />
                    <Outlet />
                  </>
                ),
                handle: {
                  breadcrumb: {
                    id: 'fake-id-2',
                    graphql: { query: SimpleQuery1, getTextContent: (data) => data.name },
                  },
                } as RouteHandle,
              },
            ],
          },
        ],
      },
    ];

    it('should be ok to display on 1 level', async () => {
      const router = createMemoryRouter(routes, { initialEntries: ['/'] });

      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <RouterProvider router={router} />
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      const navElement2 = container.querySelector('nav');
      expect(navElement2).not.toBeNull();
      const olElement2 = container.querySelector('ol');
      expect(olElement2).not.toBeNull();
      const liElements2 = container.querySelectorAll('li');
      expect(liElements2).not.toBeNull();
      expect(liElements2).toHaveLength(1);

      expect(liElements2[0].children[0].localName).toEqual('p');
      expect(liElements2[0].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements2[0].children[0]).toHaveTextContent('fake');
    });

    it('should be ok to display on 1 level with 1 ignored route', async () => {
      const router = createMemoryRouter(routes, { initialEntries: ['/level1/'] });

      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <RouterProvider router={router} />
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      const navElement2 = container.querySelector('nav');
      expect(navElement2).not.toBeNull();
      const olElement2 = container.querySelector('ol');
      expect(olElement2).not.toBeNull();
      const liElements2 = container.querySelectorAll('li');
      expect(liElements2).not.toBeNull();
      expect(liElements2).toHaveLength(1);

      expect(liElements2[0].children[0].localName).toEqual('p');
      expect(liElements2[0].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements2[0].children[0]).toHaveTextContent('fake');
    });

    it('should be ok to display on 2 levels with 1 ignored route', async () => {
      const router = createMemoryRouter(routes, { initialEntries: ['/level1/level2/'] });

      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <RouterProvider router={router} />
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      expect(container).toHaveTextContent('fake');
      expect(container).toHaveTextContent('/');
      expect(container).not.toHaveTextContent('Query1');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(3);

      expect(liElements[0].children[0].localName).toEqual('a');
      expect(liElements[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements[0].children[0]).toHaveTextContent('fake');
      expect(liElements[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements[1].localName).toEqual('li');
      expect(liElements[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[1]).toHaveTextContent('/');

      expect(liElements[2].children[0].localName).toEqual('span');
      expect(liElements[2].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');

      expect(await screen.findByText('Query1')).toBeInTheDocument();

      expect(container).toMatchSnapshot();

      const navElement2 = container.querySelector('nav');
      expect(navElement2).not.toBeNull();
      const olElement2 = container.querySelector('ol');
      expect(olElement2).not.toBeNull();
      const liElements2 = container.querySelectorAll('li');
      expect(liElements2).not.toBeNull();
      expect(liElements2).toHaveLength(3);

      expect(liElements2[0].children[0].localName).toEqual('a');
      expect(liElements2[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements2[0].children[0]).toHaveTextContent('fake');
      expect(liElements2[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements2[1].localName).toEqual('li');
      expect(liElements2[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements2[1]).toHaveTextContent('/');

      expect(liElements2[2].children[0].localName).toEqual('p');
      expect(liElements2[2].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements2[2].children[0]).toHaveTextContent('Query1');
    });
  });
});

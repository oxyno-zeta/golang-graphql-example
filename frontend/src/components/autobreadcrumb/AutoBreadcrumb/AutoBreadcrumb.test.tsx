import React from 'react';
import { Outlet, Routes, Route, MemoryRouter } from 'react-router';
import { render, screen } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import { MockedProvider } from '@apollo/client/testing';
import {
  SimpleErrorQuery,
  SimpleQuery1,
  SimpleQuery2,
  SlowQuery,
  mockedResponses,
} from './AutoBreadcrumb.storage-test';

import AutoBreadcrumb from './AutoBreadcrumb';
import AutoBreadcrumbInjector from '../AutoBreadcrumbInjector';
import AutoBreadcrumbProvider from '../AutoBreadcrumbProvider';

jest.mock('react-i18next', () => ({ useTranslation: () => ({ t: (key: string) => key }) }));

describe('autobreadcrumb/AutoBreadcrumb', () => {
  describe('Fixed texts', () => {
    const allFixedRoutes = (
      <Routes>
        <Route
          path="/"
          element={
            <AutoBreadcrumbInjector item={{ depth: 0, id: 'fake-id-1', fixed: { textContent: 'root' } }}>
              <AutoBreadcrumb />
              <Outlet />
            </AutoBreadcrumbInjector>
          }
        >
          <Route
            path="level1"
            element={
              <AutoBreadcrumbInjector item={{ depth: 1, id: 'fake-id-2', fixed: { textContent: 'level1' } }}>
                <Outlet />
              </AutoBreadcrumbInjector>
            }
          >
            <Route
              path="level2"
              element={
                <AutoBreadcrumbInjector item={{ depth: 2, id: 'fake-id-3', fixed: { textContent: 'level2' } }}>
                  <div />
                </AutoBreadcrumbInjector>
              }
            />
          </Route>
        </Route>
      </Routes>
    );

    it('should be ok to display on 1 level', () => {
      const { container } = render(
        <AutoBreadcrumbProvider>
          <MemoryRouter initialIndex={0} initialEntries={['/']}>
            {allFixedRoutes}
          </MemoryRouter>
        </AutoBreadcrumbProvider>,
      );

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
      const { container } = render(
        <AutoBreadcrumbProvider>
          <MemoryRouter initialIndex={0} initialEntries={['/level1/']}>
            {allFixedRoutes}
          </MemoryRouter>
        </AutoBreadcrumbProvider>,
      );

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
      const { container } = render(
        <AutoBreadcrumbProvider>
          <MemoryRouter initialIndex={0} initialEntries={['/level1/level2/']}>
            {allFixedRoutes}
          </MemoryRouter>
        </AutoBreadcrumbProvider>,
      );

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
      expect(liElements[2].children[0]).toHaveAttribute('href', '/level1/');

      expect(liElements[3].localName).toEqual('li');
      expect(liElements[3]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[3]).toHaveTextContent('/');

      expect(liElements[4].children[0].localName).toEqual('p');
      expect(liElements[4].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements[4].children[0]).toHaveTextContent('level2');
    });
  });

  describe('GraphQL texts', () => {
    const allGraphqlRoutes = (
      <Routes>
        <Route
          path="/"
          element={
            <AutoBreadcrumbInjector
              item={{
                depth: 0,
                id: 'fake-id-1',
                graphql: { query: SimpleQuery1, getTextContent: (data) => data.name },
              }}
            >
              <AutoBreadcrumb />
              <Outlet />
            </AutoBreadcrumbInjector>
          }
        >
          <Route
            path="level1"
            element={
              <AutoBreadcrumbInjector
                item={{
                  depth: 1,
                  id: 'fake-id-2',
                  graphql: { query: SimpleQuery2, getTextContent: (data) => data.name2 },
                }}
              >
                <Outlet />
              </AutoBreadcrumbInjector>
            }
          >
            <Route
              path="error"
              element={
                <AutoBreadcrumbInjector
                  item={{
                    depth: 2,
                    id: 'fake-id-3',
                    graphql: { query: SimpleErrorQuery, getTextContent: (data) => data.name2 },
                  }}
                >
                  <div />
                </AutoBreadcrumbInjector>
              }
            />
            <Route
              path="slow"
              element={
                <AutoBreadcrumbInjector
                  item={{
                    depth: 2,
                    id: 'fake-id-4',
                    graphql: { query: SlowQuery, getTextContent: (data) => data.slow },
                  }}
                >
                  <div />
                </AutoBreadcrumbInjector>
              }
            />
          </Route>
        </Route>
      </Routes>
    );

    it('should be ok to display on 1 level', async () => {
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/']}>
              {allGraphqlRoutes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
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
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/level1']}>
              {allGraphqlRoutes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
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
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/level1/error']}>
              {allGraphqlRoutes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
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
      expect(liElements2[2].children[0]).toHaveAttribute('href', '/level1/');

      expect(liElements[3].localName).toEqual('li');
      expect(liElements[3]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements[3]).toHaveTextContent('/');

      expect(liElements[4].children[0].localName).toEqual('span');
      expect(liElements[4].children[0]).toHaveClass('MuiSkeleton-root MuiSkeleton-text');
    });
  });

  describe('Mixed texts', () => {
    const allMixedRoutes = (
      <Routes>
        <Route
          path="/"
          element={
            <AutoBreadcrumbInjector
              item={{
                depth: 0,
                id: 'fake-id-1',
                fixed: { textContent: 'fake' },
              }}
            >
              <AutoBreadcrumb />
              <Outlet />
            </AutoBreadcrumbInjector>
          }
        >
          <Route
            path="level1"
            element={
              <AutoBreadcrumbInjector
                item={{
                  depth: 1,
                  id: 'fake-id-2',
                  graphql: { query: SimpleQuery2, getTextContent: (data) => data.name2 },
                }}
              >
                <Outlet />
              </AutoBreadcrumbInjector>
            }
          />
        </Route>
      </Routes>
    );

    it('should be ok to display on 1 level', async () => {
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/']}>
              {allMixedRoutes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
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
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/level1/']}>
              {allMixedRoutes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
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
    const routes = (
      <Routes>
        <Route
          path="/"
          element={
            <AutoBreadcrumbInjector
              item={{
                depth: 0,
                id: 'fake-id-0',
                fixed: { textContent: 'fake' },
              }}
            >
              <AutoBreadcrumb />
              <Outlet />
            </AutoBreadcrumbInjector>
          }
        >
          <Route
            path="level1"
            element={
              <AutoBreadcrumbInjector
                item={{
                  depth: 1,
                  id: 'fake-id-1',
                  ignored: true,
                }}
              >
                <Outlet />
              </AutoBreadcrumbInjector>
            }
          >
            <Route
              path="level2"
              element={
                <AutoBreadcrumbInjector
                  item={{
                    depth: 2,
                    id: 'fake-id-2',
                    graphql: { query: SimpleQuery1, getTextContent: (data) => data.name },
                  }}
                >
                  <Outlet />
                </AutoBreadcrumbInjector>
              }
            >
              <Route
                path="level3"
                element={
                  <AutoBreadcrumbInjector
                    item={{
                      depth: 3,
                      id: 'fake-id-3',
                      ignored: true,
                    }}
                  >
                    <Outlet />
                  </AutoBreadcrumbInjector>
                }
              >
                <Route
                  path="level4"
                  element={
                    <AutoBreadcrumbInjector
                      item={{
                        depth: 4,
                        id: 'fake-id-4',
                        graphql: { query: SimpleQuery2, getTextContent: (data) => data.name2 },
                      }}
                    >
                      <Outlet />
                    </AutoBreadcrumbInjector>
                  }
                />
              </Route>
            </Route>
          </Route>
        </Route>
      </Routes>
    );

    it('should be ok to display on 1 level', async () => {
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/']}>
              {routes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
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
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/level1/']}>
              {routes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
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
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/level1/level2/']}>
              {routes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
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

      expect(container).toMatchSnapshot();

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

    it('should be ok to display on 4 levels with 2 ignored routes', async () => {
      const { container } = render(
        <MockedProvider mocks={mockedResponses}>
          <AutoBreadcrumbProvider>
            <MemoryRouter initialIndex={0} initialEntries={['/level1/level2/level3/level4']}>
              {routes}
            </MemoryRouter>
          </AutoBreadcrumbProvider>
        </MockedProvider>,
      );

      expect(container).toMatchSnapshot();

      expect(container).toHaveTextContent('fake');
      expect(container).toHaveTextContent('/');
      expect(container).not.toHaveTextContent('Query1');
      expect(container).not.toHaveTextContent('Query2');

      const navElement = container.querySelector('nav');
      expect(navElement).not.toBeNull();
      const olElement = container.querySelector('ol');
      expect(olElement).not.toBeNull();
      const liElements = container.querySelectorAll('li');
      expect(liElements).not.toBeNull();
      expect(liElements).toHaveLength(5);

      expect(container).toMatchSnapshot();

      expect(liElements[0].children[0].localName).toEqual('a');
      expect(liElements[0].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements[0].children[0]).toHaveTextContent('fake');
      expect(liElements[0].children[0]).toHaveAttribute('href', '/');

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
      expect(liElements2[0].children[0]).toHaveTextContent('fake');
      expect(liElements2[0].children[0]).toHaveAttribute('href', '/');

      expect(liElements2[1].localName).toEqual('li');
      expect(liElements2[1]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements2[1]).toHaveTextContent('/');

      expect(liElements2[2].children[0].localName).toEqual('a');
      expect(liElements2[2].children[0]).toHaveClass('MuiTypography-root MuiLink-root');
      expect(liElements2[2].children[0]).toHaveTextContent('Query1');
      expect(liElements2[2].children[0]).toHaveAttribute('href', '/level1/level2/');

      expect(liElements2[3].localName).toEqual('li');
      expect(liElements2[3]).toHaveClass('MuiBreadcrumbs-separator');
      expect(liElements2[3]).toHaveTextContent('/');

      expect(liElements2[4].children[0].localName).toEqual('p');
      expect(liElements2[4].children[0]).toHaveClass('MuiTypography-root MuiTypography-body1');
      expect(liElements2[4].children[0]).toHaveTextContent('Query2');
    });
  });
});

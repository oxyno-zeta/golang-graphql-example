import React from 'react';
import { StoryFn, Meta } from '@storybook/react-vite';
import { Link, Outlet, MemoryRouter, Routes, Route } from 'react-router';
import AutoBreadcrumb from './AutoBreadcrumb';
import {
  SimpleErrorQuery,
  SimpleQuery1,
  SimpleQuery2,
  SlowQuery,
  mockedResponses,
} from './AutoBreadcrumb.storage-test';
import AutoBreadcrumbInjector from '../AutoBreadcrumbInjector';
import AutoBreadcrumbProvider from '../AutoBreadcrumbProvider';

export default {
  title: 'Components/autobreadcrumb/AutoBreadcrumb',
  component: AutoBreadcrumb,
} as Meta<typeof AutoBreadcrumb>;

const allFixedRoutes = (
  <Routes>
    <Route
      path="/"
      element={
        <AutoBreadcrumbInjector item={{ depth: 0, id: 'fake-id-0', fixed: { textContent: 'level0' } }}>
          <div style={{ marginBottom: '30px' }}>
            Menu:
            <br />
            <Link to="/">/</Link>
            <br />
            <Link to="/level1">/level1</Link>
            <br />
            <Link to="/level1/level2">/level1/level2</Link>
            <br />
            <Link to="/level1/level2/level3">/level1/level2/level3</Link>
          </div>
          <div>level0 content</div>
          <Outlet />
          <div style={{ backgroundColor: 'lightblue' }}>
            Component:
            <AutoBreadcrumb />
          </div>
        </AutoBreadcrumbInjector>
      }
    >
      <Route
        path="level1"
        element={
          <AutoBreadcrumbInjector item={{ depth: 1, id: 'fake-id-1', fixed: { textContent: 'level1' } }}>
            <div>level1 content</div>
            <Outlet />
          </AutoBreadcrumbInjector>
        }
      >
        <Route
          path=":level2"
          element={
            <AutoBreadcrumbInjector item={{ depth: 2, id: 'fake-id-2', fixed: { textContent: 'level2' } }}>
              <div>level2 content</div>
              <Outlet />
            </AutoBreadcrumbInjector>
          }
        >
          <Route
            path=":level3"
            element={
              <AutoBreadcrumbInjector item={{ depth: 3, id: 'fake-id-3', fixed: { textContent: 'level3' } }}>
                <div>level3 content</div>
              </AutoBreadcrumbInjector>
            }
          />
        </Route>
      </Route>
    </Route>
  </Routes>
);

export const AllFixed: StoryFn<typeof AutoBreadcrumb> = function C() {
  return (
    <AutoBreadcrumbProvider>
      <MemoryRouter initialIndex={0} initialEntries={['/level1/level2']}>
        {allFixedRoutes}
      </MemoryRouter>
    </AutoBreadcrumbProvider>
  );
};

const allGraphqlRoutes = (
  <Routes>
    <Route
      path="/"
      element={
        <AutoBreadcrumbInjector
          item={{ depth: 0, id: 'fake-id-0', graphql: { query: SimpleQuery1, getTextContent: (data) => data.name } }}
        >
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
          <div>level0 content</div>
          <Outlet />
          <div style={{ backgroundColor: 'lightblue' }}>
            Component:
            <AutoBreadcrumb />
          </div>
        </AutoBreadcrumbInjector>
      }
    >
      <Route
        path="level1"
        element={
          <AutoBreadcrumbInjector
            item={{ depth: 1, id: 'fake-id-1', graphql: { query: SimpleQuery2, getTextContent: (data) => data.name2 } }}
          >
            <div>level1 content</div>
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
                id: 'fake-id-2',
                graphql: { query: SimpleErrorQuery, getTextContent: (data) => data.name2 },
              }}
            >
              <div>error content</div>
              <Outlet />
            </AutoBreadcrumbInjector>
          }
        />
        <Route
          path="slow"
          element={
            <AutoBreadcrumbInjector
              item={{
                depth: 2,
                id: 'fake-id-3',
                graphql: { query: SlowQuery, getTextContent: (data) => data.slow },
              }}
            >
              <div>slow content</div>
              <Outlet />
            </AutoBreadcrumbInjector>
          }
        />
      </Route>
    </Route>
  </Routes>
);

export const AllGraphql: StoryFn<typeof AutoBreadcrumb> = function C() {
  return (
    <AutoBreadcrumbProvider>
      <MemoryRouter initialIndex={0} initialEntries={['/level1']}>
        {allGraphqlRoutes}
      </MemoryRouter>
    </AutoBreadcrumbProvider>
  );
};
AllGraphql.parameters = {
  apolloClient: {
    // Example coming from https://storybook.js.org/addons/storybook-addon-apollo-client
    mocks: mockedResponses,
  },
};

const allMixedRoutes = (
  <Routes>
    <Route
      path="/"
      element={
        <AutoBreadcrumbInjector
          item={{
            depth: 0,
            id: 'fake-id-0',
            fixed: { textContent: 'level0' },
          }}
        >
          <div style={{ marginBottom: '30px' }}>
            Menu:
            <br />
            <Link to="/">/</Link>
            <br />
            <Link to="/level1">/level1</Link>
            <br />
            <Link to="/level1/level2">/level1/level2</Link>
          </div>
          <div>level0 content</div>
          <Outlet />
          <div style={{ backgroundColor: 'lightblue' }}>
            Component:
            <AutoBreadcrumb />
          </div>
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
              graphql: { query: SlowQuery, getTextContent: (data) => data.slow },
            }}
          >
            <div>level1 content</div>
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
                fixed: { textContent: 'level2' },
              }}
            >
              <div>level2 content</div>
              <Outlet />
            </AutoBreadcrumbInjector>
          }
        />
      </Route>
    </Route>
  </Routes>
);

export const AllMixed: StoryFn<typeof AutoBreadcrumb> = function C() {
  return (
    <AutoBreadcrumbProvider>
      <MemoryRouter initialIndex={0} initialEntries={['/level1/level2']}>
        {allMixedRoutes}
      </MemoryRouter>
    </AutoBreadcrumbProvider>
  );
};
AllMixed.parameters = {
  apolloClient: {
    // Example coming from https://storybook.js.org/addons/storybook-addon-apollo-client
    mocks: mockedResponses,
  },
};

const ignoredRoutes = (
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
          <div style={{ marginBottom: '30px' }}>
            Menu:
            <br />
            <Link to="/">/</Link>
            <br />
            <Link to="/level1">/level1</Link>
            <br />
            <Link to="/level1/level2">/level1/level2</Link>
            <br />
            <Link to="/level1/level2/level3">/level1/level2/level3</Link>
            <br />
            <Link to="/level1/level2/level3/level4">/level1/level2/level3/level4</Link>
          </div>
          <div>level0 content</div>
          <Outlet />
          <div style={{ backgroundColor: 'lightblue' }}>
            Component:
            <AutoBreadcrumb />
          </div>
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
            <div>level1 content</div>
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
              <div>level2 content</div>
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
                <div>level3 content</div>
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
                  <div>level4 content</div>
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

export const IgnoredRoutes: StoryFn<typeof AutoBreadcrumb> = function C() {
  return (
    <AutoBreadcrumbProvider>
      <MemoryRouter initialIndex={0} initialEntries={['/level1/level2']}>
        {ignoredRoutes}
      </MemoryRouter>
    </AutoBreadcrumbProvider>
  );
};
IgnoredRoutes.parameters = {
  apolloClient: {
    // Example coming from https://storybook.js.org/addons/storybook-addon-apollo-client
    mocks: mockedResponses,
  },
};

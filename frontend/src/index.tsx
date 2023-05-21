import React from 'react';
import { createRoot } from 'react-dom/client';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
// import i18n
import './i18n';
import './yup-i18n';

import subRoutes from './routes/router-routes';
import App from './App';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: subRoutes,
  },
]);

const container = document.getElementById('root');
const root = createRoot(container!); // eslint-disable-line @typescript-eslint/no-non-null-assertion
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);

import React, { Suspense } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import CssBaseline from '@mui/material/CssBaseline';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import { ErrorBoundary } from 'react-error-boundary';
// import i18n
import './i18n';
import './yup-i18n';

import FallbackErrorBoundary from '~components/FallbackErrorBoundary';
import App from './App';
import MainPageCenterLoading from './components/MainPageCenterLoading';
import ConfigProvider from './components/ConfigProvider';
import TopBar from './components/TopBar';
import Footer from './components/Footer';
import ClientProvider from './components/ClientProvider';
import ThemeProvider from './components/theming/ThemeProvider';
import TimezoneProvider from './components/timezone/TimezoneProvider';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

const container = document.getElementById('root');
const root = createRoot(container!); // eslint-disable-line @typescript-eslint/no-non-null-assertion
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Suspense fallback={<MainPageCenterLoading />}>
        <ConfigProvider loadingComponent={<MainPageCenterLoading />}>
          <ErrorBoundary FallbackComponent={FallbackErrorBoundary}>
            <ClientProvider>
              <ThemeProvider themeOptions={{}}>
                <TimezoneProvider>
                  <CssBaseline />
                  <TopBar />
                  <App />
                  <Footer />
                </TimezoneProvider>
              </ThemeProvider>
            </ClientProvider>
          </ErrorBoundary>
        </ConfigProvider>
      </Suspense>
    </BrowserRouter>
  </React.StrictMode>,
);

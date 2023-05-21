import React, { Suspense } from 'react';
import { Outlet } from 'react-router-dom';
import CssBaseline from '@mui/material/CssBaseline';
import { ErrorBoundary } from 'react-error-boundary';
import FallbackErrorBoundary from '~components/FallbackErrorBoundary';
import MainPageCenterLoading from '~components/MainPageCenterLoading';
import ConfigProvider from '~components/ConfigProvider';
import TopBar from '~components/TopBar';
import Footer from '~components/Footer';
import ClientProvider from '~components/ClientProvider';
import ThemeProvider from '~components/theming/ThemeProvider';
import TimezoneProvider from '~components/timezone/TimezoneProvider';

function App() {
  return (
    <Suspense fallback={<MainPageCenterLoading />}>
      <ConfigProvider loadingComponent={<MainPageCenterLoading />}>
        <ErrorBoundary FallbackComponent={FallbackErrorBoundary}>
          <ClientProvider>
            <ThemeProvider themeOptions={{}}>
              <TimezoneProvider>
                <CssBaseline />
                <TopBar />
                <Outlet />
                <Footer />
              </TimezoneProvider>
            </ThemeProvider>
          </ClientProvider>
        </ErrorBoundary>
      </ConfigProvider>
    </Suspense>
  );
}

export default App;

import React, { Suspense } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import CssBaseline from '@mui/material/CssBaseline';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
// import i18n
import './i18n';
import './yup-i18n';

import App from './App';
import reportWebVitals from './reportWebVitals';
import MainPageCenterLoading from './components/MainPageCenterLoading';
import ConfigProvider from './components/ConfigProvider';
import TopBar from './components/TopBar';
import Footer from './components/Footer';
import ClientProvider from './components/ClientProvider';
import ThemeProvider from './components/theming/ThemeProvider';

// Extend dayjs
dayjs.extend(localizedFormat);

const container = document.getElementById('root');
const root = createRoot(container!); // eslint-disable-line @typescript-eslint/no-non-null-assertion
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <Suspense fallback={<MainPageCenterLoading />}>
        <ConfigProvider loadingComponent={<MainPageCenterLoading />}>
          <ClientProvider>
            <ThemeProvider themeOptions={{}}>
              <CssBaseline />
              <TopBar />
              <App />
              <Footer />
            </ThemeProvider>
          </ClientProvider>
        </ConfigProvider>
      </Suspense>
    </BrowserRouter>
  </React.StrictMode>,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

import * as jest from 'jest-mock';
import { MockedProvider } from '@apollo/client/testing';
import { withMuiTheme } from './with-mui-theme.decorator';
import i18n from './i18next.js';

// Inject jest correctly
window.jest = jest;

// Load Roboto fonts
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import '@fontsource/material-icons';

export const parameters = {
  i18n,
  locale: 'en',
  locales: {
    en: { title: 'English', right: 'US' },
  },
  apolloClient: {
    MockedProvider,
    // any props you want to pass to MockedProvider on every story
  },
  actions: { argTypesRegex: '^on[A-Z].*' },
  controls: {
    expanded: true,
    hideNoControlsWarning: true,
    matchers: {
      color: /(background|color)$/i,
      date: /Date$/,
    },
  },
};

export const globalTypes = {
  theme: {
    name: 'Theme',
    title: 'Theme',
    description: 'Theme for your components',
    defaultValue: 'light',
    toolbar: {
      icon: 'paintbrush',
      dynamicTitle: true,
      items: [
        { value: 'light', left: '☀️', title: 'Light mode' },
        { value: 'dark', left: '🌙', title: 'Dark mode' },
      ],
    },
  },
};

export const decorators = [withMuiTheme];

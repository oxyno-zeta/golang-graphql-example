import type { Preview } from '@storybook/react';
import * as jest from 'jest-mock';
import { MockedProvider } from '@apollo/client/testing';
import { withMuiTheme } from './with-mui-theme.decorator';
import i18n from './i18next.js';

// Inject jest correctly
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
window.jest = jest;

// Load Roboto fonts
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import '@fontsource/material-icons';

const preview: Preview = {
  globals: {
    locale: 'en',
    locales: {
      en: 'English',
    },
  },
  parameters: {
    i18n,
    apolloClient: {
      MockedProvider,
      // any props you want to pass to MockedProvider on every story
    },
    actions: { argTypesRegex: '^(on|handle)[A-Z].*' },
    controls: {
      expanded: true,
      hideNoControlsWarning: true,
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
  globalTypes: {
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
  },
  decorators: [withMuiTheme],
};

export default preview;

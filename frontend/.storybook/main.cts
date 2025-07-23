import type { StorybookConfig } from '@storybook/react-vite';
import { mergeConfig } from 'vite';

const config: StorybookConfig = {
  stories: ['../src/**/*.stories.@(js|jsx|ts|tsx)'],
  addons: [
    'storybook-addon-apollo-client',
    'storybook-addon-remix-react-router',
    'storybook-react-i18next',
    '@vueless/storybook-dark-mode',
  ],
  framework: {
    name: '@storybook/react-vite',
    options: {},
  },
  async viteFinal(config) {
    // Merge custom configuration into the default config
    return mergeConfig(config, {
      // Add storybook-specific dependencies to pre-optimization
      optimizeDeps: {
        include: ['@storybook/addon-designs'],
      },
    });
  },
  core: {
    disableTelemetry: true,
    builder: '@storybook/builder-vite',
  },
  typescript: {
    check: false,
    reactDocgen: 'react-docgen',
  },
  docs: {
    autodocs: true,
  },
};

export default config;

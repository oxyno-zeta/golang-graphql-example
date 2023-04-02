module.exports = {
  stories: ['../src/**/*.stories.mdx', '../src/**/*.stories.@(js|jsx|ts|tsx)'],
  addons: [
    '@storybook/addon-links',
    '@storybook/addon-essentials',
    '@storybook/addon-interactions',
    'storybook-addon-apollo-client',
    '@storybook/addon-queryparams',
    'storybook-react-context',
    'storybook-react-i18next',
  ],
  framework: '@storybook/react',
  core: {
    builder: '@storybook/builder-vite',
    disableTelemetry: true,
    noOpen: true,
  },
  async viteFinal(config) {
    return {
      ...config,
      rollupOptions: {
        ...config.rollupOptions,
        // Externalize deps that shouldn't be bundled
        external: ['react', 'react-dom'],
        output: {
          // Global vars to use in UMD build for externalized deps
          globals: {
            react: 'React',
            'react-dom': 'ReactDOM',
          },
        },
      },
    };
  },
  typescript: {
    check: false,
    checkOptions: {},
    reactDocgen: 'react-docgen-typescript',
    reactDocgenTypescriptOptions: {
      skipChildrenPropWithoutDoc: false,
      allowSyntheticDefaultImports: false, // speeds up storybook build time
      esModuleInterop: false, // speeds up storybook build time
      shouldExtractLiteralValuesFromEnum: true, // makes union prop types like variant and size appear as select controls
      shouldRemoveUndefinedFromOptional: true, // makes string and boolean types that can be undefined appear as inputs and switches
    },
  },
};

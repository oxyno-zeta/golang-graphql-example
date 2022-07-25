export type ConfigModel = {
  // Graphql Endpoint
  graphqlEndpoint: string;
  // Configuration cookie domain
  configCookieDomain: string;
};

export const defaultConfig: ConfigModel = {
  graphqlEndpoint: '/api/graphql',
  configCookieDomain: location.hostname, // eslint-disable-line no-restricted-globals
};

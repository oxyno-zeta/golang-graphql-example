export type ConfigModel = {
  // Graphql Endpoint
  graphqlEndpoint: string;
};

export const defaultConfig: ConfigModel = {
  graphqlEndpoint: '/api/graphql',
};

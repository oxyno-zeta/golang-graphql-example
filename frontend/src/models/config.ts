export type ConfigModel = {
  // Graphql Endpoint
  graphqlEndpoint: string;
  // Configuration cookie domain
  configCookieDomain: string;
  // Sign Out URL
  // Can be found by looking at: https://IDP.easymile.com/auth/realms/REALM/.well-known/openid-configuration
  // Example of Wellknown: https://idp-dev.easymile.com/auth/realms/em-only/.well-known/openid-configuration
  // Partial value example: https://idp-dev.easymile.com/auth/realms/em-only/protocol/openid-connect/logout
  // On this add the current URL as encoded parameter for redirect
  // Documentation: https://www.keycloak.org/docs/latest/securing_apps/index.html#logout
  // Blocked by this PR: https://github.com/keycloak/keycloak/issues/12680
  // TODO Waiting for Keycloak v19 but code is ready
  oidcSignOutURL?: string;
  oidcClientID?: string;
};

export const defaultConfig: ConfigModel = {
  graphqlEndpoint: '/api/graphql',
  configCookieDomain: location.hostname,
};

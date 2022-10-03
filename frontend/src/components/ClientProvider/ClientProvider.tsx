import React, { useContext, ReactNode } from 'react';
import { ApolloClient, InMemoryCache, createHttpLink, ServerError, ApolloProvider, ApolloLink } from '@apollo/client';
import { onError } from '@apollo/client/link/error';
import { ConfigModel } from '../../models/config';
import ConfigContext from '../../contexts/ConfigContext';

interface Props {
  children: ReactNode;
}

function generateClient(cfg: ConfigModel) {
  // Create apollo link to force header injection
  // That force is done to ensure a 401 Unauthorized error when backend is protected by Oauth2-Proxy
  // Documentation: https://oauth2-proxy.github.io/oauth2-proxy/docs/behaviour/
  const forceHeaders = new ApolloLink((operation, forward) => {
    operation.setContext(({ headers = {} }) => ({
      headers: {
        ...headers,
        Accept: 'application/json',
      },
    }));

    return forward(operation);
  });

  // Create apollo link
  // This is needed to create a valid forwarder of cookies (for authentication)
  // Documentation: https://www.apollographql.com/docs/react/networking/authentication/#cookie
  const apolloLink = createHttpLink({
    uri: cfg.graphqlEndpoint,
    // Do not touch this parameter
    credentials: 'include',
  });
  // Create an error link that will allow to force reload the page in case of a unauthorized error arrive
  // This can happened when a user stay on the page without changing or reloading it and
  // the cookie isn't valid anymore, in this case, the next request (done by Apollo Client) won't pass
  // and people will get an error. To avoid that, when an unauthorized error arrive, a page reload
  // will be triggered.
  const errorLink = onError(({ networkError }) => {
    // Check if error that is coming from server and it is an unauthorized status code.
    if (networkError && (networkError as ServerError).statusCode === 401) {
      window.location.reload();
    }
  });
  // Create apollo client
  const client = new ApolloClient({
    link: errorLink.concat(forceHeaders).concat(apolloLink),
    // Add memory cache
    cache: new InMemoryCache(),
    // Connect to dev tools only on dev
    connectToDevTools: process.env.NODE_ENV !== 'production',
  });

  return client;
}

function ClientProvider({ children }: Props) {
  const cfg = useContext(ConfigContext);

  return <ApolloProvider client={generateClient(cfg)}>{children}</ApolloProvider>;
}

export default ClientProvider;

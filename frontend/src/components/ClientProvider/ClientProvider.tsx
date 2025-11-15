import React, { useContext, type ReactNode } from 'react';
import { ApolloClient, InMemoryCache, ServerError, ApolloLink, HttpLink } from '@apollo/client';
import { ApolloProvider } from '@apollo/client/react';
import { ErrorLink } from '@apollo/client/link/error';
import { type ConfigModel } from '~models/config';
import { Observable } from 'rxjs';
import ConfigContext from '../../contexts/ConfigContext';
import { fromApolloContextErrorToWithTraceError } from '../../utils/WithTraceError';

interface Props {
  readonly children: ReactNode;
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
  const apolloLink = new HttpLink({
    uri: cfg.graphqlEndpoint,
    // Do not touch this parameter
    credentials: 'include',
  });
  // Create an error link that will allow to force reload the page in case of a unauthorized error arrive
  // This can happened when a user stay on the page without changing or reloading it and
  // the cookie isn't valid anymore, in this case, the next request (done by Apollo Client) won't pass
  // and people will get an error. To avoid that, when an unauthorized error arrive, a page reload
  // will be triggered.
  const errorLink = new ErrorLink(({ error }) => {
    // Check if error that is coming from server and it is an unauthorized status code.
    if (ServerError.is(error) && error.statusCode === 401) {
      window.location.reload();
    }
  });

  // Create link to manage trace id and request id in errors
  const withTraceErrorLink = new ErrorLink(
    ({ error, operation }) =>
      new Observable((observer) => {
        // Get context
        const context = operation.getContext();

        // Default
        observer.error(fromApolloContextErrorToWithTraceError(error, context));
      }),
  );

  // Create apollo client
  const client = new ApolloClient({
    link: ApolloLink.from([errorLink, forceHeaders, withTraceErrorLink, apolloLink]),
    // Add memory cache
    cache: new InMemoryCache(),
    devtools: {
      // Connect to dev tools only on dev
      enabled: process.env.NODE_ENV !== 'production',
    },
  });

  return client;
}

function ClientProvider({ children }: Props) {
  const cfg = useContext(ConfigContext);

  return <ApolloProvider client={generateClient(cfg)}>{children}</ApolloProvider>;
}

export default ClientProvider;

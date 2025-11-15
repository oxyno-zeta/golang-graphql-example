import { CombinedGraphQLErrors, type ServerError } from '@apollo/client';
import { AxiosError, type InternalAxiosRequestConfig } from 'axios';
import { GraphQLError } from 'graphql';
import WithTraceError, { fromAxiosErrorToWithTraceError } from '~utils/WithTraceError';

// Build data for tests
export const forbiddenNetworkError: ServerError = {
  name: 'Forbidden',
  message: 'Forbidden',
  response: new Response(),
  statusCode: 403,
  bodyText: '',
};
export const forbiddenAxiosError: AxiosError = new AxiosError('fake', 'ERR_BAD_REQUEST', undefined, null, {
  data: { error: 'forbidden', extensions: { code: 'FORBIDDEN' } },
  status: 403,
  statusText: 'Forbidden',
  headers: { 'X-Request-ID': 'request-id', 'X-Trace-ID': 'trace-id' },
  config: { url: 'http://fake.com', method: 'get' } as InternalAxiosRequestConfig,
});
export const forbiddenWithTraceAxiosError = fromAxiosErrorToWithTraceError(forbiddenAxiosError);
export const internalServerErrorAxiosError: AxiosError = new AxiosError('fake', 'ERR_BAD_REQUEST', undefined, null, {
  data: 'Internal server error',
  status: 500,
  statusText: 'Internal server error',
  headers: {},
  config: { url: 'http://fake.com', method: 'get' } as InternalAxiosRequestConfig,
});
export const internalServerErrorWithTraceDataWithTraceAxiosError =
  fromAxiosErrorToWithTraceError(internalServerErrorAxiosError);
export const simpleGraphqlErrorWithoutExtension: GraphQLError = new GraphQLError('simple graphql error');
export const simpleCombinedGraphQLErrorWithoutExtension = new CombinedGraphQLErrors({
  errors: [simpleGraphqlErrorWithoutExtension],
});
export const simpleForbiddenGraphqlError: GraphQLError = new GraphQLError('forbidden graphql error', {
  extensions: { code: 'FORBIDDEN' },
});
export const simpleForbiddenCombinedGraphQLError = new CombinedGraphQLErrors({ errors: [simpleForbiddenGraphqlError] });
export const simpleInternalServerErrorGraphqlError: GraphQLError = new GraphQLError(
  'internal server error graphql error',
  {
    extensions: { code: 'INTERNAL_SERVER_ERROR' },
  },
);
export const simpleInternalServerErrorCombinedGraphQLError = new CombinedGraphQLErrors({
  errors: [simpleInternalServerErrorGraphqlError],
});
export const simpleInternalServerWithTraceErrorCombinedGraphQLError = new WithTraceError(
  simpleInternalServerErrorCombinedGraphQLError,
  'request-id',
  'trace-id',
);

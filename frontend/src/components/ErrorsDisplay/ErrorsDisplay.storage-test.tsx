import { CombinedGraphQLErrors, type ServerError } from '@apollo/client';
import { AxiosError } from 'axios';
import { GraphQLError } from 'graphql';
import { WithTraceError } from '~components/ClientProvider';

// Build data for tests
export const forbiddenNetworkError: ServerError = {
  name: 'Forbidden',
  message: 'Forbidden',
  response: new Response(),
  statusCode: 403,
  bodyText: '',
};
export const forbiddenAxiosError: AxiosError = new AxiosError(
  'fake',
  '403',
  undefined,
  null,
  new Response({
    error: 'forbidden',
    extensions: { code: 'FORBIDDEN' },
  }),
);
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

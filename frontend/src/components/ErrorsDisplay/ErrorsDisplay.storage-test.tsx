import { CombinedGraphQLErrors, type ServerError } from '@apollo/client';
import { GraphQLError } from 'graphql';

// Build data for tests
export const forbiddenNetworkError: ServerError = {
  name: 'Forbidden',
  message: 'Forbidden',
  response: new Response(),
  statusCode: 403,
  bodyText: '',
};
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

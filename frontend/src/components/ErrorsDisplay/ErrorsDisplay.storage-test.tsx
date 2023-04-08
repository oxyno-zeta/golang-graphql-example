import { ServerError } from '@apollo/client';
import { GraphQLError } from 'graphql';

// Build data for tests
export const forbiddenNetworkError: ServerError = {
  name: 'Forbidden',
  message: 'Forbidden',
  response: new Response(),
  statusCode: 403,
  result: {},
};
export const forbiddenNetworkErrorWithErrors: ServerError = {
  name: 'Forbidden',
  message: 'Forbidden',
  response: new Response(),
  statusCode: 403,
  result: { errors: [{ message: 'fake message', path: ['fake', 'path'] }] },
};
export const forbiddenNetworkErrorWithMultipleErrors: ServerError = {
  name: 'Forbidden',
  message: 'Forbidden',
  response: new Response(),
  statusCode: 403,
  result: {
    errors: [
      { message: 'fake message', path: ['fake', 'path'] },
      { message: 'fake message 2', path: ['fake', 'path2'] },
    ],
  },
};
export const simpleGraphqlErrorWithoutExtension: GraphQLError = new GraphQLError('simple graphql error');
export const simpleForbiddenGraphqlError: GraphQLError = new GraphQLError('forbidden graphql error', {
  extensions: { code: 'FORBIDDEN' },
});
export const simpleInternalServerErrorGraphqlError: GraphQLError = new GraphQLError(
  'internal server error graphql error',
  {
    extensions: { code: 'INTERNAL_SERVER_ERROR' },
  },
);

import React from 'react';
import renderer from 'react-test-renderer';
import { ApolloError, ServerError } from '@apollo/client';
import { GraphQLError } from 'graphql';

import GraphqlErrors from './GraphqlErrors';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

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

describe('GraphqlErrors', () => {
  it('should return null when no error or errors are present', () => {
    const component = renderer.create(<GraphqlErrors />);
    const tree = component.toJSON();

    expect(tree).toBeNull();
  });

  it('should display a network error when error is present', () => {
    const component = renderer.create(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkError,
          })
        }
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(2);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('Forbidden');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display a network error with 1 error when error is present', () => {
    const component = renderer.create(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkErrorWithErrors,
          })
        }
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(2);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('fake.path fake message');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display a network error with multiple errors when error is present', () => {
    const component = renderer.create(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkErrorWithMultipleErrors,
          })
        }
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(3);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('fake.path fake message');
    expect(allP[2].props.children).toEqual('fake.path2 fake message 2');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display a network error when errors are present', () => {
    const component = renderer.create(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkError,
          }),
        ]}
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(2);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('Forbidden');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display a graphql error without extension when error is present', () => {
    const component = renderer.create(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'one graphql apollo error',
            graphQLErrors: [simpleGraphqlErrorWithoutExtension],
          })
        }
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(2);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('simple graphql error');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display a graphql error without extension when errors are present', () => {
    const component = renderer.create(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'one graphql apollo error',
            graphQLErrors: [simpleGraphqlErrorWithoutExtension],
          }),
        ]}
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(2);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('simple graphql error');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display a graphql error with extension when error is present', () => {
    const component = renderer.create(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'one graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError],
          })
        }
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(2);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('common.errorCode.FORBIDDEN');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display a graphql error with extension when errors are present', () => {
    const component = renderer.create(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'one graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError],
          }),
        ]}
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(2);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('common.errorCode.FORBIDDEN');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display two graphql error with extension when error is present', () => {
    const component = renderer.create(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'two graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError, simpleInternalServerErrorGraphqlError],
          })
        }
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(3);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('common.errorCode.FORBIDDEN');
    expect(allP[2].props.children).toEqual('common.errorCode.INTERNAL_SERVER_ERROR');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display two graphql error with extension when errors are present (1 item with 2)', () => {
    const component = renderer.create(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'two graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError, simpleInternalServerErrorGraphqlError],
          }),
        ]}
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(3);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('common.errorCode.FORBIDDEN');
    expect(allP[2].props.children).toEqual('common.errorCode.INTERNAL_SERVER_ERROR');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });

  it('should display two graphql error with extension when errors are present (2 items)', () => {
    const component = renderer.create(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'two graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError],
          }),
          new ApolloError({
            errorMessage: 'two graphql apollo error',
            graphQLErrors: [simpleInternalServerErrorGraphqlError],
          }),
        ]}
      />,
    );
    const tree = component.toJSON();

    const allP = component.root.findAllByType('p');
    expect(allP).toHaveLength(3);
    expect(allP[0].props.children).toEqual(['common.errors', ':']);
    expect(allP[1].props.children).toEqual('common.errorCode.FORBIDDEN');
    expect(allP[2].props.children).toEqual('common.errorCode.INTERNAL_SERVER_ERROR');

    expect(tree).not.toBeNull();
    expect(tree).toMatchSnapshot();
  });
});

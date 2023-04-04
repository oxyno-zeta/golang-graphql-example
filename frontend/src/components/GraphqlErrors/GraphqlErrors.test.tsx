import React from 'react';
import { render } from '@testing-library/react';
import { ApolloError } from '@apollo/client';
import {
  forbiddenNetworkError,
  forbiddenNetworkErrorWithErrors,
  forbiddenNetworkErrorWithMultipleErrors,
  simpleForbiddenGraphqlError,
  simpleGraphqlErrorWithoutExtension,
  simpleInternalServerErrorGraphqlError,
} from './GraphqlErrors.storage-test';

import GraphqlErrors from './GraphqlErrors';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('GraphqlErrors', () => {
  it('should return null when no error or errors are present', () => {
    const { container } = render(<GraphqlErrors />);

    expect(container).toMatchSnapshot();
  });

  it('should display a network error when error is present', () => {
    const { container } = render(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkError,
          })
        }
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'Forbidden'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a network error with 1 error when error is present', () => {
    const { container } = render(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkErrorWithErrors,
          })
        }
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'fake.path fake message'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a network error with multiple errors when error is present', () => {
    const { container } = render(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkErrorWithMultipleErrors,
          })
        }
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(3);

    // Prepare values
    const values = ['common.errors:', 'fake.path fake message', 'fake.path2 fake message 2'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a network error when errors are present', () => {
    const { container } = render(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkError,
          }),
        ]}
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'Forbidden'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a graphql error without extension when error is present', () => {
    const { container } = render(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'one graphql apollo error',
            graphQLErrors: [simpleGraphqlErrorWithoutExtension],
          })
        }
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'simple graphql error'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a graphql error without extension when errors are present', () => {
    const { container } = render(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'one graphql apollo error',
            graphQLErrors: [simpleGraphqlErrorWithoutExtension],
          }),
        ]}
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'simple graphql error'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a graphql error with extension when error is present', () => {
    const { container } = render(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'one graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError],
          })
        }
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'common.errorCode.FORBIDDEN'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a graphql error with extension when errors are present', () => {
    const { container } = render(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'one graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError],
          }),
        ]}
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'common.errorCode.FORBIDDEN'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display two graphql error with extension when error is present', () => {
    const { container } = render(
      <GraphqlErrors
        error={
          new ApolloError({
            errorMessage: 'two graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError, simpleInternalServerErrorGraphqlError],
          })
        }
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(3);

    // Prepare values
    const values = ['common.errors:', 'common.errorCode.FORBIDDEN', 'common.errorCode.INTERNAL_SERVER_ERROR'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display two graphql error with extension when errors are present (1 item with 2)', () => {
    const { container } = render(
      <GraphqlErrors
        errors={[
          new ApolloError({
            errorMessage: 'two graphql apollo error',
            graphQLErrors: [simpleForbiddenGraphqlError, simpleInternalServerErrorGraphqlError],
          }),
        ]}
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(3);

    // Prepare values
    const values = ['common.errors:', 'common.errorCode.FORBIDDEN', 'common.errorCode.INTERNAL_SERVER_ERROR'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display two graphql error with extension when errors are present (2 items)', () => {
    const { container } = render(
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

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(3);

    // Prepare values
    const values = ['common.errors:', 'common.errorCode.FORBIDDEN', 'common.errorCode.INTERNAL_SERVER_ERROR'];
    allP.forEach((item, index) => {
      expect(item.innerHTML).toEqual(values[index]);
    });

    expect(container).toMatchSnapshot();
  });
});

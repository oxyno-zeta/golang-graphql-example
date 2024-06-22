import React from 'react';
import { render } from '@testing-library/react';
import { ApolloError } from '@apollo/client';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import {
  forbiddenNetworkError,
  forbiddenNetworkErrorWithErrors,
  forbiddenNetworkErrorWithMultipleErrors,
  simpleForbiddenGraphqlError,
  simpleGraphqlErrorWithoutExtension,
  simpleInternalServerErrorGraphqlError,
} from './ErrorsDisplay.storage-test';

import ErrorsDisplay from './ErrorsDisplay';
import {
  GraphqlErrorsExtensionsCodeForbiddenCustomComponentMapKey,
  NetworkErrorCustomComponentMapKey,
} from './constants';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('ErrorsDisplay', () => {
  it('should return null when no error or errors are present', () => {
    const { container } = render(<ErrorsDisplay />);

    expect(container).toMatchSnapshot();
  });

  it('should display a network error when error is present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a network error when error is present with a custom component', () => {
    function Fake({ input }: { input: string }) {
      return <p>{input}</p>;
    }

    const { container } = render(
      <ErrorsDisplay
        error={
          new ApolloError({
            errorMessage: 'network apollo error',
            networkError: forbiddenNetworkError,
          })
        }
        customErrorComponents={{
          [NetworkErrorCustomComponentMapKey]: Fake,
        }}
        customErrorComponentProps={{
          [NetworkErrorCustomComponentMapKey]: { input: 'fake' },
        }}
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'Forbidden'];
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a network error with 1 error when error is present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a network error with multiple errors when error is present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a network error when errors are present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a graphql error without extension when error is present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a graphql error without extension when errors are present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a graphql error with extension when error is present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display a graphql error with extension when errors are present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display two graphql error with extension when error is present', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display two graphql error with extension when errors are present (1 item with 2)', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display two graphql error with extension when errors are present (2 items)', () => {
    const { container } = render(
      <ErrorsDisplay
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
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display two graphql error with extension when errors are present (2 items with 1 custom component)', () => {
    function Fake({ input }: { input: string }) {
      return <p>{input}</p>;
    }

    const { container } = render(
      <ErrorsDisplay
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
        customErrorComponents={{
          [GraphqlErrorsExtensionsCodeForbiddenCustomComponentMapKey]: Fake,
        }}
        customErrorComponentProps={{
          [GraphqlErrorsExtensionsCodeForbiddenCustomComponentMapKey]: { input: 'fake' },
        }}
      />,
    );

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(3);

    // Prepare values
    const values = ['common.errors:', 'fake', 'common.errorCode.INTERNAL_SERVER_ERROR'];
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });
    expect(container).not.toHaveTextContent('common.errorCode.FORBIDDEN');

    expect(container).toMatchSnapshot();
  });

  it('should display classic error when error is present', () => {
    const { container } = render(<ErrorsDisplay error={new Error('error1')} />);

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'error1'];
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display classic error when errors are present (1 item)', () => {
    const { container } = render(<ErrorsDisplay errors={[new Error('error1')]} />);

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(2);

    // Prepare values
    const values = ['common.errors:', 'error1'];
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });

  it('should display classic error when errors are present (2 items)', () => {
    const { container } = render(<ErrorsDisplay errors={[new Error('error1'), new Error('error2')]} />);

    const allP = container.querySelectorAll('p');
    expect(allP).toHaveLength(3);

    // Prepare values
    const values = ['common.errors:', 'error1', 'error2'];
    values.forEach((item) => {
      expect(container).toHaveTextContent(item);
    });

    expect(container).toMatchSnapshot();
  });
});

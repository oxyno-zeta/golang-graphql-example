import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import { ApolloError } from '@apollo/client';
import ErrorsDisplay, { type Props } from './ErrorsDisplay';
import {
  forbiddenNetworkError,
  simpleForbiddenGraphqlError,
  simpleGraphqlErrorWithoutExtension,
  simpleInternalServerErrorGraphqlError,
} from './ErrorsDisplay.storage-test';

export default {
  title: 'Components/ErrorsDisplay',
  component: ErrorsDisplay,
  args: {
    error: new ApolloError({
      errorMessage: 'network apollo error',
      networkError: forbiddenNetworkError,
    }),
  },
} as Meta<typeof ErrorsDisplay>;

export const Playground: StoryFn<typeof ErrorsDisplay> = function C(args: Props) {
  return <ErrorsDisplay {...args} />;
};

export const ClassicError: StoryFn<typeof ErrorsDisplay> = function C() {
  return <ErrorsDisplay error={new Error('fake error !')} />;
};

export const GraphQLNetworkError: StoryFn<typeof ErrorsDisplay> = function C() {
  return (
    <ErrorsDisplay
      error={
        new ApolloError({
          errorMessage: 'network apollo error',
          networkError: forbiddenNetworkError,
        })
      }
    />
  );
};

export const OneGraphQLErrorWithoutExtension: StoryFn<typeof ErrorsDisplay> = function C() {
  return (
    <ErrorsDisplay
      error={
        new ApolloError({
          errorMessage: 'one graphql apollo error',
          graphQLErrors: [simpleGraphqlErrorWithoutExtension],
        })
      }
    />
  );
};

export const OneGraphQLErrorWithExtension: StoryFn<typeof ErrorsDisplay> = function C() {
  return (
    <ErrorsDisplay
      error={
        new ApolloError({
          errorMessage: 'one graphql apollo error',
          graphQLErrors: [simpleForbiddenGraphqlError],
        })
      }
    />
  );
};

export const TwoGraphQLErrorWithExtension: StoryFn<typeof ErrorsDisplay> = function C() {
  return (
    <ErrorsDisplay
      error={
        new ApolloError({
          errorMessage: 'two graphql apollo error',
          graphQLErrors: [simpleForbiddenGraphqlError, simpleInternalServerErrorGraphqlError],
        })
      }
    />
  );
};

export const OneGraphQLErrorWithoutMargin: StoryFn<typeof ErrorsDisplay> = function C() {
  return (
    <ErrorsDisplay
      error={
        new ApolloError({
          errorMessage: 'one graphql apollo error',
          graphQLErrors: [simpleForbiddenGraphqlError],
        })
      }
      noMargin
    />
  );
};

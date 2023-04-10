import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import { ApolloError } from '@apollo/client';
import ErrorsDisplay, { Props } from './ErrorsDisplay';
import {
  forbiddenNetworkError,
  simpleForbiddenGraphqlError,
  simpleGraphqlErrorWithoutExtension,
  simpleInternalServerErrorGraphqlError,
} from './ErrorsDisplay.storage-test';

export default {
  title: 'Components/ErrorsDisplay',
  component: ErrorsDisplay,
} as Meta<typeof ErrorsDisplay>;

const Template: StoryFn<typeof ErrorsDisplay> = function C(args: Props) {
  return <ErrorsDisplay {...args} />;
};

export const Playground = {
  render: Template,
  args: {
    error: new ApolloError({
      errorMessage: 'network apollo error',
      networkError: forbiddenNetworkError,
    }),
  },
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
      noMargin
      error={
        new ApolloError({
          errorMessage: 'one graphql apollo error',
          graphQLErrors: [simpleForbiddenGraphqlError],
        })
      }
    />
  );
};

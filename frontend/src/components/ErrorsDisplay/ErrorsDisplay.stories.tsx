import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
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
} as ComponentMeta<typeof ErrorsDisplay>;

const Template: ComponentStory<typeof ErrorsDisplay> = function C(args: Props) {
  return <ErrorsDisplay {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {
  error: new ApolloError({
    errorMessage: 'network apollo error',
    networkError: forbiddenNetworkError,
  }),
};

export const ClassicError: ComponentStory<typeof ErrorsDisplay> = function C() {
  return <ErrorsDisplay error={new Error('fake error !')} />;
};

export const GraphQLNetworkError: ComponentStory<typeof ErrorsDisplay> = function C() {
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

export const OneGraphQLErrorWithoutExtension: ComponentStory<typeof ErrorsDisplay> = function C() {
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

export const OneGraphQLErrorWithExtension: ComponentStory<typeof ErrorsDisplay> = function C() {
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

export const TwoGraphQLErrorWithExtension: ComponentStory<typeof ErrorsDisplay> = function C() {
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

export const OneGraphQLErrorWithoutMargin: ComponentStory<typeof ErrorsDisplay> = function C() {
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

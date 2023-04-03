import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { ApolloError } from '@apollo/client';
import GraphqlErrors, { Props } from './GraphqlErrors';
import {
  forbiddenNetworkError,
  simpleForbiddenGraphqlError,
  simpleGraphqlErrorWithoutExtension,
  simpleInternalServerErrorGraphqlError,
} from './GraphqlErrors.storage-test';

export default {
  title: 'Components/GraphqlErrors',
  component: GraphqlErrors,
} as ComponentMeta<typeof GraphqlErrors>;

const Template: ComponentStory<typeof GraphqlErrors> = function C(args: Props) {
  return <GraphqlErrors {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {
  error: new ApolloError({
    errorMessage: 'network apollo error',
    networkError: forbiddenNetworkError,
  }),
};

export const NetworkError: ComponentStory<typeof GraphqlErrors> = function C() {
  return (
    <GraphqlErrors
      error={
        new ApolloError({
          errorMessage: 'network apollo error',
          networkError: forbiddenNetworkError,
        })
      }
    />
  );
};

export const OneGraphQLErrorWithoutExtension: ComponentStory<typeof GraphqlErrors> = function C() {
  return (
    <GraphqlErrors
      error={
        new ApolloError({
          errorMessage: 'one graphql apollo error',
          graphQLErrors: [simpleGraphqlErrorWithoutExtension],
        })
      }
    />
  );
};

export const OneGraphQLErrorWithExtension: ComponentStory<typeof GraphqlErrors> = function C() {
  return (
    <GraphqlErrors
      error={
        new ApolloError({
          errorMessage: 'one graphql apollo error',
          graphQLErrors: [simpleForbiddenGraphqlError],
        })
      }
    />
  );
};

export const TwoGraphQLErrorWithExtension: ComponentStory<typeof GraphqlErrors> = function C() {
  return (
    <GraphqlErrors
      error={
        new ApolloError({
          errorMessage: 'two graphql apollo error',
          graphQLErrors: [simpleForbiddenGraphqlError, simpleInternalServerErrorGraphqlError],
        })
      }
    />
  );
};

export const OneGraphQLErrorWithoutMargin: ComponentStory<typeof GraphqlErrors> = function C() {
  return (
    <GraphqlErrors
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

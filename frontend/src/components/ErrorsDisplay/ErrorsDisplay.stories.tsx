import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import { CombinedGraphQLErrors } from '@apollo/client';
import ErrorsDisplay, { type Props } from './ErrorsDisplay';
import {
  forbiddenNetworkError,
  simpleCombinedGraphQLErrorWithoutExtension,
  simpleForbiddenCombinedGraphQLError,
  simpleForbiddenGraphqlError,
  simpleInternalServerErrorGraphqlError,
  simpleInternalServerWithTraceErrorCombinedGraphQLError,
} from './ErrorsDisplay.storage-test';

export default {
  title: 'Components/ErrorsDisplay',
  component: ErrorsDisplay,
  args: {
    error: forbiddenNetworkError,
  },
} as Meta<typeof ErrorsDisplay>;

export const Playground: StoryFn<typeof ErrorsDisplay> = function C(args: Props) {
  return <ErrorsDisplay {...args} />;
};

export const ClassicError: StoryFn<typeof ErrorsDisplay> = function C() {
  return <ErrorsDisplay error={new Error('fake error !')} />;
};

export const GraphQLNetworkError: StoryFn<typeof ErrorsDisplay> = function C() {
  return <ErrorsDisplay error={forbiddenNetworkError} />;
};

export const OneGraphQLErrorWithoutExtension: StoryFn<typeof ErrorsDisplay> = function C() {
  return <ErrorsDisplay error={simpleCombinedGraphQLErrorWithoutExtension} />;
};

export const OneGraphQLErrorWithExtension: StoryFn<typeof ErrorsDisplay> = function C() {
  return <ErrorsDisplay error={simpleForbiddenCombinedGraphQLError} />;
};

export const TwoGraphQLErrorWithExtension: StoryFn<typeof ErrorsDisplay> = function C() {
  return (
    <ErrorsDisplay
      error={
        new CombinedGraphQLErrors({
          errors: [simpleForbiddenGraphqlError, simpleInternalServerErrorGraphqlError],
        })
      }
    />
  );
};

export const OneGraphQLErrorWithoutMargin: StoryFn<typeof ErrorsDisplay> = function C() {
  return <ErrorsDisplay error={simpleForbiddenCombinedGraphQLError} noMargin />;
};

export const OneWithTraceGraphQLError: StoryFn<typeof ErrorsDisplay> = function C() {
  return <ErrorsDisplay error={simpleInternalServerWithTraceErrorCombinedGraphQLError} />;
};

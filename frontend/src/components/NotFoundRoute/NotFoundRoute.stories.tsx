import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import NotFoundRoute, { Props } from './NotFoundRoute';

export default {
  title: 'Components/NotFoundRoute',
  component: NotFoundRoute,
} as Meta<typeof NotFoundRoute>;

export const Playground: StoryFn<typeof NotFoundRoute> = function C(args: Props) {
  return <NotFoundRoute {...args} />;
};

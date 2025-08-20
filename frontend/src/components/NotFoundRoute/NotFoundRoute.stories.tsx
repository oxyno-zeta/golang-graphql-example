import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import NotFoundRoute, { type Props } from './NotFoundRoute';

export default {
  title: 'Components/NotFoundRoute',
  component: NotFoundRoute,
} as Meta<typeof NotFoundRoute>;

export const Playground: StoryFn<typeof NotFoundRoute> = function C(args: Props) {
  return <NotFoundRoute {...args} />;
};

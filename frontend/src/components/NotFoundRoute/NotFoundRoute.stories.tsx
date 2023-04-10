import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import NotFoundRoute, { Props } from './NotFoundRoute';

export default {
  title: 'Components/NotFoundRoute',
  component: NotFoundRoute,
} as Meta<typeof NotFoundRoute>;

const Template: StoryFn<typeof NotFoundRoute> = function C(args: Props) {
  return <NotFoundRoute {...args} />;
};

export const Playground = {
  render: Template,
  args: {},
};

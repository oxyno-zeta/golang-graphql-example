import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import NotFoundRoute, { Props } from './NotFoundRoute';

export default {
  title: 'Components/NotFoundRoute',
  component: NotFoundRoute,
} as ComponentMeta<typeof NotFoundRoute>;

const Template: ComponentStory<typeof NotFoundRoute> = function C(args: Props) {
  return <NotFoundRoute {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {};

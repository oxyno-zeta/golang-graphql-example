import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import NoData, { Props } from './NoData';

export default {
  title: 'Components/NoData',
  component: NoData,
} as ComponentMeta<typeof NoData>;

const Template: ComponentStory<typeof NoData> = function C(args: Props) {
  return <NoData {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {};

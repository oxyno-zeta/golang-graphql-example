import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import NoData, { Props } from './NoData';

export default {
  title: 'Components/NoData',
  component: NoData,
} as Meta<typeof NoData>;

const Template: StoryFn<typeof NoData> = function C(args: Props) {
  return <NoData {...args} />;
};

export const Playground = {
  render: Template,
  args: {},
};

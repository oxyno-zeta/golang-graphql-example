import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import NoData, { Props } from './NoData';

export default {
  title: 'Components/NoData',
  component: NoData,
} as Meta<typeof NoData>;

export const Playground: StoryFn<typeof NoData> = function C(args: Props) {
  return <NoData {...args} />;
};

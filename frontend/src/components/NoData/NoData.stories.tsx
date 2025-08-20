import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import NoData, { type Props } from './NoData';

export default {
  title: 'Components/NoData',
  component: NoData,
} as Meta<typeof NoData>;

export const Playground: StoryFn<typeof NoData> = function C(args: Props) {
  return <NoData {...args} />;
};

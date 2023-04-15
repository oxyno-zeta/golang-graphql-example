import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import GridTableViewSwitcher, { Props } from './GridTableViewSwitcher';

export default {
  title: 'Components/GridTableViewSwitcher',
  component: GridTableViewSwitcher,
  argTypes: {
    setGridView: { name: 'setGridView' },
  },
  args: {
    gridView: true,
  },
} as Meta<typeof GridTableViewSwitcher>;

export const Playground: StoryFn<typeof GridTableViewSwitcher> = function C(args: Props) {
  return <GridTableViewSwitcher {...args} />;
};

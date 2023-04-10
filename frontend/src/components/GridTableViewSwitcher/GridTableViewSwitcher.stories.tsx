import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import GridTableViewSwitcher, { Props } from './GridTableViewSwitcher';

export default {
  title: 'Components/GridTableViewSwitcher',
  component: GridTableViewSwitcher,
  argTypes: {
    setGridView: { name: 'setGridView' },
  },
} as Meta<typeof GridTableViewSwitcher>;

const Template: StoryFn<typeof GridTableViewSwitcher> = function C(args: Props) {
  return <GridTableViewSwitcher {...args} />;
};

export const Playground = {
  render: Template,
  args: {
    gridView: true,
  },
};

import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import GridTableViewSwitcher, { Props } from './GridTableViewSwitcher';

export default {
  title: 'Components/GridTableViewSwitcher',
  component: GridTableViewSwitcher,
} as ComponentMeta<typeof GridTableViewSwitcher>;

const Template: ComponentStory<typeof GridTableViewSwitcher> = function C(args: Props) {
  return <GridTableViewSwitcher {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {
  setGridView: () => {},
  gridView: true,
};

import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import GridTableViewSwitcher, { type Props } from './GridTableViewSwitcher';
import GridTableViewSwitcherProvider from '../GridTableViewSwitcherProvider';

export default {
  title: 'Components/gridTableViewSwitch/GridTableViewSwitcher',
  component: GridTableViewSwitcher,
} as Meta<typeof GridTableViewSwitcher>;

export const Playground: StoryFn<typeof GridTableViewSwitcher> = function C(args: Props) {
  return (
    <GridTableViewSwitcherProvider>
      <GridTableViewSwitcher {...args} />;
    </GridTableViewSwitcherProvider>
  );
};

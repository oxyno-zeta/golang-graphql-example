import React from 'react';
import { StoryFn, Meta } from '@storybook/react-vite';
import MainPageCenterLoading, { Props } from './MainPageCenterLoading';

export default {
  title: 'Components/MainPageCenterLoading',
  component: MainPageCenterLoading,
} as Meta<typeof MainPageCenterLoading>;

export const Playground: StoryFn<typeof MainPageCenterLoading> = function C(args: Props) {
  return <MainPageCenterLoading {...args} />;
};

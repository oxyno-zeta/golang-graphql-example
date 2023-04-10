import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import MainPageCenterLoading, { Props } from './MainPageCenterLoading';

export default {
  title: 'Components/MainPageCenterLoading',
  component: MainPageCenterLoading,
} as Meta<typeof MainPageCenterLoading>;

const Template: StoryFn<typeof MainPageCenterLoading> = function C(args: Props) {
  return <MainPageCenterLoading {...args} />;
};

export const Playground = {
  render: Template,
  args: {},
};

import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import MainPageCenterLoading, { Props } from './MainPageCenterLoading';

export default {
  title: 'Components/MainPageCenterLoading',
  component: MainPageCenterLoading,
} as ComponentMeta<typeof MainPageCenterLoading>;

const Template: ComponentStory<typeof MainPageCenterLoading> = function C(args: Props) {
  return <MainPageCenterLoading {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {};

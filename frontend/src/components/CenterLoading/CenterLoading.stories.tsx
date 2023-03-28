import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import CenterLoading, { CenterLoadingProps } from './CenterLoading';

export default {
  title: 'Components/CenterLoading',
  component: CenterLoading,
} as ComponentMeta<typeof CenterLoading>;

const Template: ComponentStory<typeof CenterLoading> = function C(args: CenterLoadingProps) {
  return <CenterLoading {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {};

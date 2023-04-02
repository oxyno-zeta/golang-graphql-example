import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import Stack from '@mui/material/Stack';
import CenterLoading, { Props } from './CenterLoading';

export default {
  title: 'Components/CenterLoading',
  component: CenterLoading,
} as ComponentMeta<typeof CenterLoading>;

const Template: ComponentStory<typeof CenterLoading> = function C(args: Props) {
  return <CenterLoading {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {};

export const Colors: ComponentStory<typeof CenterLoading> = function C() {
  return (
    <Stack spacing={2} maxWidth={300}>
      <CenterLoading />
      <CenterLoading circularProgressProps={{ color: 'secondary' }} />
      <CenterLoading circularProgressProps={{ color: 'success' }} />
      <CenterLoading circularProgressProps={{ color: 'error' }} />
    </Stack>
  );
};

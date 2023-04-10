import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import Footer, { Props } from './Footer';

export default {
  title: 'Components/Footer',
  component: Footer,
} as Meta<typeof Footer>;

const Template: StoryFn<typeof Footer> = function C(args: Props) {
  return <Footer {...args} />;
};

export const Playground = {
  render: Template,
  args: {},
};

import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import Footer, { Props } from './Footer';

export default {
  title: 'Components/Footer',
  component: Footer,
} as ComponentMeta<typeof Footer>;

const Template: ComponentStory<typeof Footer> = function C(args: Props) {
  return <Footer {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {};

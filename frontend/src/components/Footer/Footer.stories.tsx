import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import Footer, { type Props } from './Footer';

export default {
  title: 'Components/Footer',
  component: Footer,
} as Meta<typeof Footer>;

export const Playground: StoryFn<typeof Footer> = function C(args: Props) {
  return <Footer {...args} />;
};

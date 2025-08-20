import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import OneLineLabelValueDisplay, { type Props } from './OneLineLabelValueDisplay';

export default {
  title: 'Components/OneLineLabelValueDisplay',
  component: OneLineLabelValueDisplay,
  args: { labelText: 'Fake tooltip !' },
} as Meta<typeof OneLineLabelValueDisplay>;

export const Playground: StoryFn<typeof OneLineLabelValueDisplay> = function C(args: Props) {
  return <OneLineLabelValueDisplay {...args} />;
};

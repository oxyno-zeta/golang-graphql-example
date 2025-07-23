import React from 'react';
import { StoryFn, Meta } from '@storybook/react-vite';
import OneLineLabelValueDisplay, { Props } from './OneLineLabelValueDisplay';

export default {
  title: 'Components/OneLineLabelValueDisplay',
  component: OneLineLabelValueDisplay,
  args: { labelText: 'Fake tooltip !' },
} as Meta<typeof OneLineLabelValueDisplay>;

export const Playground: StoryFn<typeof OneLineLabelValueDisplay> = function C(args: Props) {
  return <OneLineLabelValueDisplay {...args} />;
};

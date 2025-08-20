import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import HelpTooltipButton, { type Props } from './HelpTooltipButton';

export default {
  title: 'Components/HelpTooltipButton',
  component: HelpTooltipButton,
  args: { tooltipTitle: 'Fake tooltip !' },
} as Meta<typeof HelpTooltipButton>;

export const Playground: StoryFn<typeof HelpTooltipButton> = function C(args: Props) {
  return <HelpTooltipButton {...args} />;
};

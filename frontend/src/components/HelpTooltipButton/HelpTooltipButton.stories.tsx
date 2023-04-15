import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import HelpTooltipButton, { Props } from './HelpTooltipButton';

export default {
  title: 'Components/HelpTooltipButton',
  component: HelpTooltipButton,
  args: { tooltipTitle: 'Fake tooltip !' },
} as Meta<typeof HelpTooltipButton>;

export const Playground: StoryFn<typeof HelpTooltipButton> = function C(args: Props) {
  return <HelpTooltipButton {...args} />;
};

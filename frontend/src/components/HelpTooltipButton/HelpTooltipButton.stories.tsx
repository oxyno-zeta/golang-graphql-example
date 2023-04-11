import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import HelpTooltipButton, { Props } from './HelpTooltipButton';

export default {
  title: 'Components/HelpTooltipButton',
  component: HelpTooltipButton,
} as Meta<typeof HelpTooltipButton>;

const Template: StoryFn<typeof HelpTooltipButton> = function C(args: Props) {
  return <HelpTooltipButton {...args} />;
};

export const Playground = {
  render: Template,
  args: { tooltipTitle: 'Fake tooltip !' },
};

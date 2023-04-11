import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import HelpForcedTooltip, { Props } from './HelpForcedTooltip';

export default {
  title: 'Components/HelpForcedTooltip',
  component: HelpForcedTooltip,
} as Meta<typeof HelpForcedTooltip>;

const Template: StoryFn<typeof HelpForcedTooltip> = function C(args: Props) {
  return <HelpForcedTooltip {...args} />;
};

export const Playground = {
  render: Template,
  args: { tooltipTitle: 'Fake tooltip !' },
};

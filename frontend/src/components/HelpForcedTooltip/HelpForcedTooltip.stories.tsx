import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import HelpForcedTooltip, { Props } from './HelpForcedTooltip';

export default {
  title: 'Components/HelpForcedTooltip',
  component: HelpForcedTooltip,
  args: { tooltipTitle: 'Fake tooltip !' },
} as Meta<typeof HelpForcedTooltip>;

export const Playground: StoryFn<typeof HelpForcedTooltip> = function C(args: Props) {
  return <HelpForcedTooltip {...args} />;
};

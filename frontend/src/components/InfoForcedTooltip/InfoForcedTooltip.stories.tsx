import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import InfoForcedTooltip, { Props } from './InfoForcedTooltip';

export default {
  title: 'Components/InfoForcedTooltip',
  component: InfoForcedTooltip,
  args: { tooltipTitle: 'Fake tooltip !' },
} as Meta<typeof InfoForcedTooltip>;

export const Playground: StoryFn<typeof InfoForcedTooltip> = function C(args: Props) {
  return <InfoForcedTooltip {...args} />;
};

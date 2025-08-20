import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import InfoForcedTooltip, { type Props } from './InfoForcedTooltip';

export default {
  title: 'Components/InfoForcedTooltip',
  component: InfoForcedTooltip,
  args: { tooltipTitle: 'Fake tooltip !' },
} as Meta<typeof InfoForcedTooltip>;

export const Playground: StoryFn<typeof InfoForcedTooltip> = function C(args: Props) {
  return <InfoForcedTooltip {...args} />;
};

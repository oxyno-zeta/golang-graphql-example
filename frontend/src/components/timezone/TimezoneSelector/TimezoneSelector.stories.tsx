import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import TimezoneProvider from '~components/timezone/TimezoneProvider';
import TimezoneSelector from './TimezoneSelector';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

export default {
  title: 'Components/timezone/TimezoneSelector',
  component: TimezoneSelector,
} as Meta<typeof TimezoneSelector>;

export const Playbook: StoryFn<typeof TimezoneSelector> = function C() {
  return (
    <TimezoneProvider>
      <TimezoneSelector />
    </TimezoneProvider>
  );
};

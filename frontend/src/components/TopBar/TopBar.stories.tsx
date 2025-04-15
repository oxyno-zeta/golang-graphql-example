import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import { withRouter } from 'storybook-addon-remix-react-router';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import TimezoneProvider from '~components/timezone/TimezoneProvider';
import ConfigContext from '../../contexts/ConfigContext';
import { defaultConfig } from '../../models/config';
import TopBar from './TopBar';
import FakeUserInfo from './components/FakeUserInfo';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

export default {
  title: 'Components/TopBar',
  component: TopBar,
  decorators: [withRouter],
} as Meta<typeof TopBar>;

export const Playbook: StoryFn<typeof TopBar> = function C() {
  return (
    <TimezoneProvider>
      <ConfigContext.Provider value={defaultConfig}>
        <TopBar />
      </ConfigContext.Provider>
    </TimezoneProvider>
  );
};
Playbook.parameters = {
  reactRouter: {
    routePath: '/route',
  },
  apolloClient: {
    // Example coming from https://storybook.js.org/addons/storybook-addon-apollo-client
    // mocks: [
    //   {
    //     request: {
    //       query: DashboardPageQuery,
    //     },
    //     result: {
    //       data: {
    //         viewer: null,
    //       },
    //     },
    //   },
    // ],
  },
};

export const WithFakeUserInfo: StoryFn<typeof TopBar> = function C() {
  return (
    <TimezoneProvider>
      <ConfigContext.Provider value={defaultConfig}>
        <TopBar topBarUserMenuProps={{ UserInfoComponent: FakeUserInfo }} />
      </ConfigContext.Provider>
    </TimezoneProvider>
  );
};

export const DisableUserMenu: StoryFn<typeof TopBar> = function C() {
  return (
    <TimezoneProvider>
      <ConfigContext.Provider value={defaultConfig}>
        <TopBar disableUserMenu />
      </ConfigContext.Provider>
    </TimezoneProvider>
  );
};

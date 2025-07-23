import React, { ReactNode } from 'react';
import { StoryFn, Meta } from '@storybook/react-vite';
import { withRouter } from 'storybook-addon-remix-react-router';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import SvgIcon from '@mui/material/SvgIcon';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import { mdiAccessPoint } from '@mdi/js';
import TopBar, { TopBarSpacer } from '~components/TopBar';
import TimezoneProvider from '~components/timezone/TimezoneProvider';
import ConfigContext from '../../../contexts/ConfigContext';
import { defaultConfig } from '../../../models/config';
import ContentDisplayDrawer, { Props } from './ContentDisplayDrawer';
import ListNavItemButton from '../ListNavItemButton';
import PageDrawer from '../PageDrawer';
import PageDrawerSettingsProvider from '../PageDrawerSettingsProvider';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

export default {
  title: 'Components/drawer/ContentDisplayDrawer',
  component: ContentDisplayDrawer,
  args: {
    defaultDrawerWidth: 200,
    drawerElement: <div>Content !</div>,
  },
  decorators: [withRouter],
} as Meta<typeof ContentDisplayDrawer>;

function RemoveStorybookPadding({ children }: { readonly children: ReactNode }) {
  return <div style={{ margin: '-1rem' }}>{children}</div>;
}

function Content() {
  return (
    <div style={{ margin: '10px' }}>
      <Grid container spacing={2}>
        <Grid size={8}>
          <Paper>xs=8</Paper>
        </Grid>
        <Grid size={4}>
          <Paper>xs=4</Paper>
        </Grid>
        <Grid size={4}>
          <Paper>xs=4</Paper>
        </Grid>
        <Grid size={8}>
          <Paper>xs=8</Paper>
        </Grid>
      </Grid>
    </div>
  );
}

export const DisableTopSpace: StoryFn<typeof ContentDisplayDrawer> = function C({ disableTopSpacer, ...args }: Props) {
  return (
    <RemoveStorybookPadding>
      <ContentDisplayDrawer disableTopSpacer {...args}>
        <Content />
      </ContentDisplayDrawer>
    </RemoveStorybookPadding>
  );
};

export const DisableResize: StoryFn<typeof ContentDisplayDrawer> = function C({ disableResize, ...args }: Props) {
  return (
    <RemoveStorybookPadding>
      <ContentDisplayDrawer disableResize {...args}>
        <Content />
      </ContentDisplayDrawer>
    </RemoveStorybookPadding>
  );
};

export const WithTopBarOnly: StoryFn<typeof ContentDisplayDrawer> = function C({ ...args }: Props) {
  return (
    <RemoveStorybookPadding>
      <TimezoneProvider>
        <ConfigContext.Provider value={defaultConfig}>
          <TopBar />
          <ContentDisplayDrawer {...args}>
            <TopBarSpacer />
            <Content />
          </ContentDisplayDrawer>
        </ConfigContext.Provider>
      </TimezoneProvider>
    </RemoveStorybookPadding>
  );
};

export const WithTopBarAndRightMenu: StoryFn<typeof ContentDisplayDrawer> = function C({ ...args }: Props) {
  return (
    <RemoveStorybookPadding>
      <TimezoneProvider>
        <PageDrawerSettingsProvider>
          <ConfigContext.Provider value={defaultConfig}>
            <TopBar />
            <PageDrawer
              defaultDrawerWidth={200}
              renderDrawerContent={({ listItemButtonSx, listItemIconSx, listItemTextSx }) => (
                <>
                  <ListItem disablePadding>
                    <ListItemButton dense sx={listItemButtonSx}>
                      <ListItemIcon sx={listItemIconSx}>
                        <SvgIcon>
                          <path d={mdiAccessPoint} />
                        </SvgIcon>
                      </ListItemIcon>
                      <ListItemText
                        primary="Menu"
                        secondary="Sub menu"
                        secondaryTypographyProps={{ noWrap: true, fontSize: 12, lineHeight: '16px' }}
                        sx={listItemTextSx}
                      />
                    </ListItemButton>
                  </ListItem>
                  <ListItem dense disablePadding>
                    <ListItemButton dense sx={listItemButtonSx}>
                      <ListItemIcon sx={listItemIconSx}>
                        <SvgIcon>
                          <path d={mdiAccessPoint} />
                        </SvgIcon>
                      </ListItemIcon>
                      <ListItemText sx={listItemTextSx}>Menu</ListItemText>
                    </ListItemButton>
                  </ListItem>
                  <ListItem dense disablePadding>
                    <ListItemButton dense sx={listItemButtonSx}>
                      <ListItemIcon sx={listItemIconSx}>
                        <SvgIcon>
                          <path d={mdiAccessPoint} />
                        </SvgIcon>
                      </ListItemIcon>
                      <ListItemText sx={listItemTextSx}>Menu</ListItemText>
                    </ListItemButton>
                  </ListItem>
                  <ListItem dense disablePadding>
                    <ListItemButton dense sx={listItemButtonSx}>
                      <ListItemIcon sx={listItemIconSx}>
                        <SvgIcon>
                          <path d={mdiAccessPoint} />
                        </SvgIcon>
                      </ListItemIcon>
                      <ListItemText sx={listItemTextSx}>Menu</ListItemText>
                    </ListItemButton>
                  </ListItem>
                  <ListItem disablePadding>
                    <ListNavItemButton listItemButtonProps={{ sx: listItemButtonSx, dense: true }} to="/fail">
                      <ListItemIcon sx={listItemIconSx}>
                        <SvgIcon>
                          <path d={mdiAccessPoint} />
                        </SvgIcon>
                      </ListItemIcon>
                      <ListItemText sx={listItemTextSx}>Link</ListItemText>
                    </ListNavItemButton>
                  </ListItem>
                </>
              )}
            >
              <ContentDisplayDrawer {...args}>
                <Content />
              </ContentDisplayDrawer>
            </PageDrawer>
          </ConfigContext.Provider>
        </PageDrawerSettingsProvider>
      </TimezoneProvider>
    </RemoveStorybookPadding>
  );
};

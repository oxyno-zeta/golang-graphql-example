import React, { ReactNode, useContext } from 'react';
import { StoryFn, Meta } from '@storybook/react';
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
import TopBar from '~components/TopBar';
import TimezoneProvider from '~components/timezone/TimezoneProvider';
import PageDrawerContext from '~contexts/PageDrawerContext';
import ConfigContext from '../../../contexts/ConfigContext';
import { defaultConfig } from '../../../models/config';
import PageDrawer, { Props } from './PageDrawer';
import ListNavItemButton from '../ListNavItemButton';
import OpenDrawerButton from '../OpenDrawerButton';
import PageDrawerSettingsProvider from '../PageDrawerSettingsProvider';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

function RenderContent() {
  const { onDrawerToggle } = useContext(PageDrawerContext);
  return (
    <div>
      <div>
        <OpenDrawerButton onDrawerToggle={onDrawerToggle} />
      </div>
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

export default {
  title: 'Components/drawer/PageDrawer',
  component: PageDrawer,
  args: {
    defaultDrawerWidth: 200,
    renderDrawerContent: ({ listItemButtonSx, listItemIconSx, listItemTextSx }) => (
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
    ),
    children: <RenderContent />,
  },
  decorators: [withRouter],
} as Meta<typeof PageDrawer>;

function RemoveStorybookPadding({ children }: { readonly children: ReactNode }) {
  return <div style={{ margin: '-1rem' }}>{children}</div>;
}

export const DisableTopSpace: StoryFn<typeof PageDrawer> = function C({ disableTopSpacer, ...args }: Props) {
  return (
    <RemoveStorybookPadding>
      <PageDrawerSettingsProvider>
        <PageDrawer {...args} disableTopSpacer />
      </PageDrawerSettingsProvider>
    </RemoveStorybookPadding>
  );
};

export const DisableResize: StoryFn<typeof PageDrawer> = function C({ disableResize, ...args }: Props) {
  return (
    <RemoveStorybookPadding>
      <PageDrawerSettingsProvider>
        <PageDrawer {...args} disableResize />
      </PageDrawerSettingsProvider>
    </RemoveStorybookPadding>
  );
};

export const DisableCollapse: StoryFn<typeof PageDrawer> = function C({ disableCollapse, ...args }: Props) {
  return (
    <RemoveStorybookPadding>
      <PageDrawerSettingsProvider>
        <PageDrawer {...args} disableCollapse />
      </PageDrawerSettingsProvider>
    </RemoveStorybookPadding>
  );
};

export const WithTopBar: StoryFn<typeof PageDrawer> = function C({ ...args }: Props) {
  return (
    <RemoveStorybookPadding>
      <TimezoneProvider>
        <PageDrawerSettingsProvider>
          <ConfigContext.Provider value={defaultConfig}>
            <TopBar />
            <PageDrawer {...args} />
          </ConfigContext.Provider>
        </PageDrawerSettingsProvider>
      </TimezoneProvider>
    </RemoveStorybookPadding>
  );
};

import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import CssBaseline from '@mui/material/CssBaseline';
import Typography from '@mui/material/Typography';
import { useTheme } from '@mui/material/styles';
import ThemeProvider from '~components/theming/ThemeProvider';
import ToggleColorModeMenuItem, { Props } from './ToggleColorModeMenuItem';

export default {
  title: 'Components/theming/ToggleColorModeMenuItem',
  component: ToggleColorModeMenuItem,
} as Meta<typeof ToggleColorModeMenuItem>;

function Info() {
  const theme = useTheme();
  return <Typography>Theme mode:{theme.palette.mode}</Typography>;
}

export const Playbook: StoryFn<typeof ToggleColorModeMenuItem> = function C(args: Props) {
  return (
    <ThemeProvider themeOptions={{}}>
      <CssBaseline />
      <ToggleColorModeMenuItem {...args} />
      <Info />
    </ThemeProvider>
  );
};

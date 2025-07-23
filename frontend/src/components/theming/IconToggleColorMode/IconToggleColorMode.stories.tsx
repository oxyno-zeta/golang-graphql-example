import React from 'react';
import { StoryFn, Meta } from '@storybook/react-vite';
import CssBaseline from '@mui/material/CssBaseline';
import Typography from '@mui/material/Typography';
import { useTheme } from '@mui/material/styles';
import ThemeProvider from '~components/theming/ThemeProvider';
import IconToggleColorMode, { Props } from './IconToggleColorMode';

export default {
  title: 'Components/theming/IconToggleColorMode',
  component: IconToggleColorMode,
} as Meta<typeof IconToggleColorMode>;

function Info() {
  const theme = useTheme();
  return <Typography>Theme mode:{theme.palette.mode}</Typography>;
}

export const Playbook: StoryFn<typeof IconToggleColorMode> = function C(args: Props) {
  return (
    <ThemeProvider themeOptions={{}}>
      <CssBaseline />
      <IconToggleColorMode {...args} />
      <Info />
    </ThemeProvider>
  );
};

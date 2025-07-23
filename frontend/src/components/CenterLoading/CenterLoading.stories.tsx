import React from 'react';
import { StoryFn, Meta } from '@storybook/react-vite';
import Stack from '@mui/material/Stack';
import { useTranslation } from 'react-i18next';
import CenterLoading, { Props } from './CenterLoading';

export default {
  title: 'Components/CenterLoading',
  component: CenterLoading,
} as Meta<typeof CenterLoading>;

export const Playground: StoryFn<typeof CenterLoading> = function C(args: Props) {
  return <CenterLoading {...args} />;
};

export const Colors: StoryFn<typeof CenterLoading> = function C() {
  return (
    <Stack spacing={2}>
      <CenterLoading />
      <CenterLoading circularProgressProps={{ color: 'secondary' }} />
      <CenterLoading circularProgressProps={{ color: 'success' }} />
      <CenterLoading circularProgressProps={{ color: 'error' }} />
    </Stack>
  );
};

export const Subtitles: StoryFn<typeof CenterLoading> = function C() {
  const { t } = useTranslation();
  return (
    <Stack spacing={2}>
      <CenterLoading subtitle={t('common.loadingText')} />
      <CenterLoading circularProgressProps={{ color: 'error' }} subtitle={t('common.errors')} />
    </Stack>
  );
};

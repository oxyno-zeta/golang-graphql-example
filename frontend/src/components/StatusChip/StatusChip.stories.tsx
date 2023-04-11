import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import Stack from '@mui/material/Stack';
import StatusChip, { Props } from './StatusChip';

export default {
  title: 'Components/StatusChip',
  component: StatusChip,
  args: {
    label: 'Fake',
    // Cleaned because not planned to be used like this
    onDelete: undefined,
  },
} as Meta<typeof StatusChip>;

export const Playground: StoryFn<typeof StatusChip> = function C(args: Props) {
  return <StatusChip {...args} />;
};

export const Colors: StoryFn<typeof StatusChip> = function C(args: Props) {
  return (
    <Stack spacing={2}>
      <div>
        <StatusChip {...args} label="None" />
      </div>
      <div>
        <StatusChip {...args} color="default" label="default" />
      </div>
      <div>
        <StatusChip {...args} color="error" label="error" />
      </div>
      <div>
        <StatusChip {...args} color="info" label="info" />
      </div>
      <div>
        <StatusChip {...args} color="primary" label="primary" />
      </div>
      <div>
        <StatusChip {...args} color="secondary" label="secondary" />
      </div>
      <div>
        <StatusChip {...args} color="success" label="success" />
      </div>
      <div>
        <StatusChip {...args} color="warning" label="warning" />
      </div>
    </Stack>
  );
};

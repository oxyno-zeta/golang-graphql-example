import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import AcceptDialog, { Props } from './AcceptDialog';

const storybookProps = {
  open: true,
  title: 'Title',
  content: 'Main content !',
  dialogProps: { disablePortal: true },
};

export default {
  title: 'Components/AcceptDialog',
  component: AcceptDialog,
  args: storybookProps,
} as Meta<typeof AcceptDialog>;

export const Playground: StoryFn<typeof AcceptDialog> = function C({ handleOk, ...args }: Props) {
  return (
    <AcceptDialog
      {...args}
      handleOk={() => {
        handleOk();
        return Promise.resolve();
      }}
    />
  );
};

export const OkDisabled: StoryFn<typeof AcceptDialog> = function C(args: Props) {
  return <AcceptDialog okDisabled {...args} />;
};

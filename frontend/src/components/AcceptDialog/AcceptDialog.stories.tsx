import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import AcceptDialog, { Props } from './AcceptDialog';

const storybookProps = {
  open: true,
  title: 'Title',
  content: 'Main content !',
  handleClose: () => {
    // eslint-disable-next-line no-alert
    alert('closed !');
  },
  handleOk: () => {
    // eslint-disable-next-line no-alert
    alert('ok !');
    return Promise.resolve();
  },
  dialogProps: { disablePortal: true },
};

export default {
  title: 'Components/AcceptDialog',
  component: AcceptDialog,
} as ComponentMeta<typeof AcceptDialog>;

const Template: ComponentStory<typeof AcceptDialog> = function C(args: Props) {
  return <AcceptDialog {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {
  ...storybookProps,
};

export const OkDisabled: ComponentStory<typeof AcceptDialog> = function C() {
  return <AcceptDialog okDisabled {...storybookProps} />;
};

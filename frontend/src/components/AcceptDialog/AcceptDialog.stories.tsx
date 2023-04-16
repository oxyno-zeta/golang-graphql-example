import React, { useState } from 'react';
import { StoryFn, Meta } from '@storybook/react';
import Button from '@mui/material/Button';
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

export const Playground: StoryFn<typeof AcceptDialog> = function C({ handleOk, handleClose, ...args }: Props) {
  const [open, setOpen] = useState(false);
  return (
    <>
      <Button
        onClick={() => {
          setOpen((v) => !v);
        }}
      >
        Click
      </Button>
      <AcceptDialog
        {...args}
        open={open}
        handleClose={() => {
          handleClose();
          setOpen(false);
        }}
        handleOk={() => {
          handleOk();
          setOpen(false);
          return Promise.resolve();
        }}
      />
    </>
  );
};

export const OkDisabled: StoryFn<typeof AcceptDialog> = function C({ handleOk, handleClose, ...args }: Props) {
  const [open, setOpen] = useState(false);
  return (
    <>
      <Button
        onClick={() => {
          setOpen((v) => !v);
        }}
      >
        Click
      </Button>
      <AcceptDialog
        {...args}
        okDisabled
        open={open}
        handleClose={() => {
          handleClose();
          setOpen(false);
        }}
        handleOk={() => {
          handleOk();
          setOpen(false);
          return Promise.resolve();
        }}
      />
    </>
  );
};

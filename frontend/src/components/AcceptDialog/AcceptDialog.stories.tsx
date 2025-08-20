import React, { useState } from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import Button from '@mui/material/Button';
import AcceptDialog, { type Props } from './AcceptDialog';

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

export const Playground: StoryFn<typeof AcceptDialog> = function C({ onSubmit, onClose, ...args }: Props) {
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
        onClose={() => {
          onClose();
          setOpen(false);
        }}
        onSubmit={() => {
          onSubmit();
          setOpen(false);
          return Promise.resolve();
        }}
        open={open}
      />
    </>
  );
};

export const OkDisabled: StoryFn<typeof AcceptDialog> = function C({ onSubmit, onClose, ...args }: Props) {
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
        onClose={() => {
          onClose();
          setOpen(false);
        }}
        onSubmit={() => {
          onSubmit();
          setOpen(false);
          return Promise.resolve();
        }}
        open={open}
      />
    </>
  );
};

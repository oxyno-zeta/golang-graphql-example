import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import { mdiMenu } from '@mdi/js';
import IconButton from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import Title, { Props } from './Title';

export default {
  title: 'Components/Title',
  component: Title,
  args: {
    title: 'Fake',
  },
} as Meta<typeof Title>;

export const Playground: StoryFn<typeof Title> = function C(args: Props) {
  return <Title {...args} />;
};

export const LeftElement: StoryFn<typeof Title> = function C() {
  return (
    <Title
      leftElement={
        <IconButton>
          <SvgIcon>
            <path d={mdiMenu} />
          </SvgIcon>
        </IconButton>
      }
      title="Fake"
    />
  );
};

export const RightElement: StoryFn<typeof Title> = function C() {
  return (
    <Title
      rightElement={
        <IconButton>
          <SvgIcon>
            <path d={mdiMenu} />
          </SvgIcon>
        </IconButton>
      }
      title="Fake"
    />
  );
};

export const RightAndLeftElement: StoryFn<typeof Title> = function C() {
  return (
    <Title
      leftElement={
        <IconButton>
          <SvgIcon>
            <path d={mdiMenu} />
          </SvgIcon>
        </IconButton>
      }
      rightElement={
        <IconButton>
          <SvgIcon>
            <path d={mdiMenu} />
          </SvgIcon>
        </IconButton>
      }
      title="Fake"
    />
  );
};

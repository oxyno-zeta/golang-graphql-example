import React from 'react';
import { ComponentStory, ComponentMeta } from '@storybook/react';
import { mdiMenu } from '@mdi/js';
import IconButton from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import Title, { Props } from './Title';

export default {
  title: 'Components/Title',
  component: Title,
} as ComponentMeta<typeof Title>;

const Template: ComponentStory<typeof Title> = function C(args: Props) {
  return <Title {...args} />;
};

export const Playground = Template.bind({});
Playground.args = {
  title: 'Fake',
};

export const LeftElement: ComponentStory<typeof Title> = function C() {
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

export const RightElement: ComponentStory<typeof Title> = function C() {
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

export const RightAndLeftElement: ComponentStory<typeof Title> = function C() {
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

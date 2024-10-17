import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import { withRouter } from 'storybook-addon-remix-react-router';
import Pagination, { Props } from './Pagination';

export default {
  title: 'Components/Pagination',
  component: Pagination,
  decorators: [withRouter],
  args: {
    pageInfo: {
      hasNextPage: true,
      hasPreviousPage: false,
      endCursor: 'fake-end',
    },
    maxPaginationSize: 10,
  },
} as Meta<typeof Pagination>;

export const Playground: StoryFn<typeof Pagination> = function C(args: Props) {
  return <Pagination {...args} onFirstPage={undefined} onNextPage={undefined} onPreviousPage={undefined} />;
};
Playground.parameters = {
  reactRouter: {
    routePath: '/route',
  },
};

export const OnlyNextPage: StoryFn<typeof Pagination> = function C(args: Props) {
  return (
    <Pagination
      {...args}
      onFirstPage={undefined}
      onNextPage={undefined}
      onPreviousPage={undefined}
      pageInfo={{ hasNextPage: true, hasPreviousPage: false, endCursor: 'fake-end' }}
    />
  );
};
OnlyNextPage.parameters = {
  reactRouter: {
    routePath: '/route',
  },
};

export const OnlyPreviousPage: StoryFn<typeof Pagination> = function C(args: Props) {
  return (
    <Pagination
      {...args}
      onFirstPage={undefined}
      onNextPage={undefined}
      onPreviousPage={undefined}
      pageInfo={{ hasNextPage: false, hasPreviousPage: true, startCursor: 'fake-start' }}
    />
  );
};
OnlyPreviousPage.parameters = {
  reactRouter: {
    routePath: '/route',
  },
};

export const AllEnabled: StoryFn<typeof Pagination> = function C(args: Props) {
  return (
    <Pagination
      {...args}
      onFirstPage={undefined}
      onNextPage={undefined}
      onPreviousPage={undefined}
      pageInfo={{
        hasNextPage: true,
        hasPreviousPage: true,
        startCursor: 'fake-start',
        endCursor: 'fake-end',
      }}
    />
  );
};
AllEnabled.parameters = {
  reactRouter: {
    routePath: '/route',
  },
};

export const AllEnabledWithFunctions: StoryFn<typeof Pagination> = function C(args: Props) {
  return (
    <Pagination
      {...args}
      pageInfo={{
        hasNextPage: true,
        hasPreviousPage: true,
        startCursor: 'fake-start',
        endCursor: 'fake-end',
      }}
    />
  );
};
AllEnabledWithFunctions.parameters = {
  reactRouter: {
    routePath: '/route',
  },
};

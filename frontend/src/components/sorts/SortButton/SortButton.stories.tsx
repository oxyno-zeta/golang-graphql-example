import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import { type SortOrderFieldModel, type SortOrderModel } from '~models/general';
import SortButton, { type Props } from './SortButton';

interface TestSortOrderModel {
  createdAt?: SortOrderModel;
  updatedAt?: SortOrderModel;
  text?: SortOrderModel;
}

const testSortFields: SortOrderFieldModel[] = [
  { field: 'createdAt', display: 'common.fields.createdAt' },
  { field: 'updatedAt', display: 'common.fields.updatedAt' },
  { field: 'text', display: 'test.fields.text' },
];

export default {
  title: 'Components/sorts/SortButton',
  component: SortButton,
  args: { sortFields: testSortFields, sorts: [{ text: 'ASC' }] },
} as Meta<typeof SortButton>;

export const Playground: StoryFn<typeof SortButton> = function C(args: Props<TestSortOrderModel>) {
  return <SortButton {...args} />;
};

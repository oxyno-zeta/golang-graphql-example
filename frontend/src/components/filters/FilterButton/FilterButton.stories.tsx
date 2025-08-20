import React from 'react';
import { type StoryFn, type Meta } from '@storybook/react-vite';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import {
  type BooleanFilterModel,
  type DateFilterModel,
  type FilterDefinitionFieldsModel,
  type StringFilterModel,
} from '~models/general';
import { booleanOperations, dateOperations, stringOperations } from '~models/general-operations';
import FilterButton, { type Props } from './FilterButton';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

interface TestFilterModel {
  AND?: TestFilterModel[];
  OR?: TestFilterModel[];
  createdAt?: DateFilterModel;
  updatedAt?: DateFilterModel;
  text?: StringFilterModel;
  done?: BooleanFilterModel;
}

const testFilterDefinitionObject: FilterDefinitionFieldsModel = {
  createdAt: {
    display: 'common.fields.createdAt',
    description: 'longgggggggggggggggggggg description',
    operations: dateOperations,
  },
  text: {
    display: 'todos.fields.text',
    operations: stringOperations,
  },
  done: {
    display: 'todos.fields.done',
    operations: booleanOperations,
  },
};

export default {
  title: 'Components/filters/FilterButton',
  component: FilterButton,
  args: { filter: {}, filterDefinitionModel: testFilterDefinitionObject },
} as Meta<typeof FilterButton>;

export const Playground: StoryFn<typeof FilterButton> = function C(args: Props<TestFilterModel>) {
  return <FilterButton {...args} />;
};

export const WithPredefinedFilters: StoryFn<typeof FilterButton> = function C(args: Props<TestFilterModel>) {
  return <FilterButton {...args} predefinedFilterObjects={[{ display: 'fake', filter: { done: { eq: true } } }]} />;
};

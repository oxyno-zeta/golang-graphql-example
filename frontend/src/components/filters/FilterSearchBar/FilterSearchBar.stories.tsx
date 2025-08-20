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
import FilterSearchBar, { type Props } from './FilterSearchBar';

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
  title: 'Components/filters/FilterSearchBar',
  component: FilterSearchBar,
  args: { filter: {}, filterDefinitionModel: testFilterDefinitionObject },
} as Meta<typeof FilterSearchBar>;

export const Playground: StoryFn<typeof FilterSearchBar> = function C(args: Props<TestFilterModel>) {
  return <FilterSearchBar {...args} />;
};

export const WithPredefinedFilters: StoryFn<typeof FilterSearchBar> = function C(args: Props<TestFilterModel>) {
  return <FilterSearchBar {...args} predefinedFilterObjects={[{ display: 'fake', filter: { done: { eq: true } } }]} />;
};

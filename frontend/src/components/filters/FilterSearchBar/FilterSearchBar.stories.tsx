import React from 'react';
import { StoryFn, Meta } from '@storybook/react';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import { BooleanFilterModel, DateFilterModel, FilterDefinitionFieldsModel, StringFilterModel } from '~models/general';
import { booleanOperations, dateOperations, stringOperations } from '~models/general-operations';
import FilterSearchBar, { Props } from './FilterSearchBar';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

type TestFilterModel = {
  AND?: TestFilterModel[];
  OR?: TestFilterModel[];
  createdAt?: DateFilterModel;
  updatedAt?: DateFilterModel;
  text?: StringFilterModel;
  done?: BooleanFilterModel;
};

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

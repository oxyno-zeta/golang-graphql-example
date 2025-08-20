import {
  type SortOrderModel,
  type DateFilterModel,
  type StringFilterModel,
  type BooleanFilterModel,
  type FilterDefinitionFieldsModel,
  type SortOrderFieldModel,
} from './general';
import { stringOperations, booleanOperations, dateOperations } from './general-operations';

export interface TodoModel {
  id: string;
  createdAt: string;
  updatedAt: string;
  text: string;
  done: boolean;
}

export interface TodoSortOrderModel {
  createdAt?: SortOrderModel;
  updatedAt?: SortOrderModel;
  text?: SortOrderModel;
}

export interface TodoFilterModel {
  AND?: TodoFilterModel[];
  OR?: TodoFilterModel[];
  createdAt?: DateFilterModel;
  updatedAt?: DateFilterModel;
  text?: StringFilterModel;
  done?: BooleanFilterModel;
}

export const todoSortFields: SortOrderFieldModel[] = [
  { field: 'createdAt', display: 'common.fields.createdAt' },
  { field: 'updatedAt', display: 'common.fields.updatedAt' },
  { field: 'text', display: 'todos.fields.text' },
];

export const todoFilterDefinitionObject: FilterDefinitionFieldsModel = {
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

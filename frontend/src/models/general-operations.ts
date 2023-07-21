import { FilterDefinitionOperationsModel } from './general';

export const booleanOperations: FilterDefinitionOperationsModel<boolean> = {
  eq: {
    display: 'common.operations.eq',
    initialValue: true,
    input: false,
    enumValues: [
      { display: 'common.boolean.true', value: true },
      { display: 'common.boolean.false', value: false },
    ],
  },
  notEq: {
    display: 'common.operations.notEq',
    initialValue: true,
    input: false,
    enumValues: [
      { display: 'common.boolean.true', value: true },
      { display: 'common.boolean.false', value: false },
    ],
  },
};

export function requiredInputValidate(value: undefined | null | string | string[]) {
  if (value === undefined || value === null || value.length === 0 || value === '') {
    return 'common.fieldValidationError.required';
  }

  // Default
  return null;
}

export const stringOperations: FilterDefinitionOperationsModel<string> = {
  eq: {
    display: 'common.operations.eq',
    initialValue: '',
    input: true,
    inputType: 'string',
    caseInsensitiveEnabled: true,
  },
  notEq: {
    display: 'common.operations.notEq',
    initialValue: '',
    input: true,
    inputType: 'string',
    caseInsensitiveEnabled: true,
  },
  contains: {
    display: 'common.operations.contains',
    initialValue: '',
    input: true,
    inputType: 'string',
    inputValidate: requiredInputValidate,
    caseInsensitiveEnabled: true,
  },
  notContains: {
    display: 'common.operations.notContains',
    initialValue: '',
    input: true,
    inputType: 'string',
    inputValidate: requiredInputValidate,
    caseInsensitiveEnabled: true,
  },
  startsWith: {
    display: 'common.operations.startsWith',
    initialValue: '',
    input: true,
    inputType: 'string',
    inputValidate: requiredInputValidate,
    caseInsensitiveEnabled: true,
  },
  notStartsWith: {
    display: 'common.operations.notStartsWith',
    initialValue: '',
    input: true,
    inputType: 'string',
    inputValidate: requiredInputValidate,
    caseInsensitiveEnabled: true,
  },
  endsWith: {
    display: 'common.operations.endsWith',
    initialValue: '',
    input: true,
    inputType: 'string',
    inputValidate: requiredInputValidate,
    caseInsensitiveEnabled: true,
  },
  notEndsWith: {
    display: 'common.operations.notEndsWith',
    initialValue: '',
    input: true,
    inputType: 'string',
    inputValidate: requiredInputValidate,
    caseInsensitiveEnabled: true,
  },
  in: {
    display: 'common.operations.in',
    initialValue: [],
    input: true,
    inputType: 'string',
    multipleValues: true,
    inputValidate: requiredInputValidate,
    caseInsensitiveEnabled: true,
  },
  notIn: {
    display: 'common.operations.notIn',
    initialValue: [],
    input: true,
    inputType: 'string',
    multipleValues: true,
    inputValidate: requiredInputValidate,
    caseInsensitiveEnabled: true,
  },
  isNull: {
    display: 'common.operations.isNull',
    input: false,
    initialValue: true,
  },
  isNotNull: {
    display: 'common.operations.isNotNull',
    input: false,
    initialValue: true,
  },
};

export const dateOperations: FilterDefinitionOperationsModel<Date> = {
  eq: {
    display: 'common.operations.eq',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  notEq: {
    display: 'common.operations.notEq',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  gte: {
    display: 'common.operations.gte',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  notGte: {
    display: 'common.operations.notGte',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  lte: {
    display: 'common.operations.lte',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  notLte: {
    display: 'common.operations.notLte',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  gt: {
    display: 'common.operations.gt',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  notGt: {
    display: 'common.operations.notGt',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  lt: {
    display: 'common.operations.lt',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  notLt: {
    display: 'common.operations.notLt',
    initialValue: null,
    input: true,
    inputType: 'date',
    inputValidate: requiredInputValidate,
  },
  isNull: {
    display: 'common.operations.isNull',
    input: false,
    initialValue: true,
  },
  isNotNull: {
    display: 'common.operations.isNotNull',
    input: false,
    initialValue: true,
  },
};

export const intOperations: FilterDefinitionOperationsModel<number> = {
  eq: {
    display: 'common.operations.eq',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  notEq: {
    display: 'common.operations.notEq',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  gte: {
    display: 'common.operations.gte',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  notGte: {
    display: 'common.operations.notGte',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  lte: {
    display: 'common.operations.lte',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  notLte: {
    display: 'common.operations.notLte',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  gt: {
    display: 'common.operations.gt',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  notGt: {
    display: 'common.operations.notGt',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  lt: {
    display: 'common.operations.lt',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  notLt: {
    display: 'common.operations.notLt',
    initialValue: 0,
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
  },
  in: {
    display: 'common.operations.in',
    initialValue: [],
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
    multipleValues: true,
  },
  notIn: {
    display: 'common.operations.notIn',
    initialValue: [],
    input: true,
    inputType: 'number',
    inputValidate: requiredInputValidate,
    multipleValues: true,
  },
  isNull: {
    display: 'common.operations.isNull',
    input: false,
    initialValue: true,
  },
  isNotNull: {
    display: 'common.operations.isNotNull',
    input: false,
    initialValue: true,
  },
};

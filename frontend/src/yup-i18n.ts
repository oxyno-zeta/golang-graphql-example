import * as yup from 'yup';

yup.setLocale({
  mixed: {
    default: () => ({ key: 'common.fieldValidationError.default' }),
    required: () => ({ key: 'common.fieldValidationError.required' }),
    defined: () => ({ key: 'common.fieldValidationError.defined' }),
    notOneOf: ({ values }) => ({ key: 'common.fieldValidationError.notOneOf', values: { values } }),
    oneOf: ({ values }) => ({ key: 'common.fieldValidationError.oneOf', values: { values } }),
  },
  array: {
    length: ({ length }) => ({ key: 'common.fieldValidationError.array.length', values: { length } }),
    max: ({ max }) => ({ key: 'common.fieldValidationError.array.max', values: { max } }),
    min: ({ min }) => ({ key: 'common.fieldValidationError.array.min', values: { min } }),
  },
  boolean: {
    isValue: () => ({ key: 'common.fieldValidationError.boolean.isValue' }),
  },
  date: {
    max: ({ max }) => ({ key: 'common.fieldValidationError.date.max', values: { max } }),
    min: ({ min }) => ({ key: 'common.fieldValidationError.date.min', values: { min } }),
  },
  object: {
    noUnknown: () => ({ key: 'common.fieldValidationError.object.noUnknown' }),
  },
  string: {
    email: () => ({ key: 'common.fieldValidationError.string.email' }),
    length: ({ length }) => ({ key: 'common.fieldValidationError.string.length', values: { length } }),
    max: ({ max }) => ({ key: 'common.fieldValidationError.string.max', values: { max } }),
    min: ({ min }) => ({ key: 'common.fieldValidationError.string.min', values: { min } }),
    lowercase: () => ({ key: 'common.fieldValidationError.string.lowercase' }),
    uppercase: () => ({ key: 'common.fieldValidationError.string.lowercase' }),
    matches: ({ regex }) => ({ key: 'common.fieldValidationError.string.matches', values: { regex } }),
    trim: () => ({ key: 'common.fieldValidationError.string.trim' }),
    uuid: () => ({ key: 'common.fieldValidationError.string.uuid' }),
    url: () => ({ key: 'common.fieldValidationError.string.url' }),
  },
  number: {
    max: ({ max }) => ({ key: 'common.fieldValidationError.number.max', values: { max } }),
    min: ({ min }) => ({ key: 'common.fieldValidationError.number.min', values: { min } }),
    integer: () => ({ key: 'common.fieldValidationError.number.integer' }),
    positive: () => ({ key: 'common.fieldValidationError.number.positive' }),
    negative: () => ({ key: 'common.fieldValidationError.number.negative' }),
    lessThan: ({ less }) => ({ key: 'common.fieldValidationError.number.lessThan', values: { less } }),
    moreThan: ({ more }) => ({ key: 'common.fieldValidationError.number.moreThan', values: { more } }),
  },
});

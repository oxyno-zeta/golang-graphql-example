export type PageInfoModel = {
  hasNextPage: boolean;
  hasPreviousPage: boolean;
  startCursor?: string;
  endCursor?: string;
};

export type PaginationInputModel = {
  first?: number;
  last?: number;
  before?: string;
  after?: string;
};

export type SortOrderModel = 'ASC' | 'DESC' | undefined;

export type SortOrderFieldModel = {
  field: string;
  display: string;
};

export type StringFilterModel = {
  eq?: string;
  notEq?: string;
  contains?: string;
  notContains?: string;
  startsWith?: string;
  notStartsWith?: string;
  endsWith?: string;
  notEndsWith?: string;
  in?: string[];
  notIn?: string[];
  isNull?: boolean;
  isNotNull?: boolean;
};

// Note: This is called "IntFilter" but
// it supports float !
export type IntFilterModel = {
  eq?: number;
  notEq?: number;
  gte?: number;
  notGte?: number;
  lte?: number;
  notLte?: number;
  gt?: number;
  notGt?: number;
  lt?: number;
  notLt?: number;
  in?: number[];
  notIn?: number[];
  isNull?: boolean;
  isNotNull?: boolean;
};

export type BooleanFilterModel = {
  eq?: boolean;
  notEq?: boolean;
};

export type DateFilterModel = {
  eq?: string;
  notEq?: string;
  gte?: string;
  notGte?: string;
  lte?: string;
  notLte?: string;
  gt?: string;
  notGt?: string;
  lt?: string;
  notLt?: string;
  isNull?: boolean;
  isNotNull?: boolean;
  // For the moment, "In" and "NotIn" aren't supported
  // Because I don't know how to manage it in the GUI for the moment
};

export type FilterDefinitionFieldObjectMetadataModel<T> = {
  display: string;
  description?: string;
  operations: FilterDefinitionOperationsModel<T>;
};

export type FilterDefinitionOperationsModel<T> = Record<string, FilterOperationMetadataModel<T>>;

export type FilterDefinitionEnumObjectModel<T> = {
  value: T;
  display: string;
  description?: string;
};

export type FilterOperationMetadataModel<T> = {
  display: string;
  description?: string;
  initialValue?: T | T[] | undefined | null | boolean;
  input?: boolean;
  inputType?: string;
  inputValidate?: (value: undefined | null | string | string[]) => string | null | undefined;
  // Put that flag with "input" for a multi value input
  multipleValues?: boolean;
  enumValues?: FilterDefinitionEnumObjectModel<T>[];
  // Put that flag with "enumValues" for a multi select enum
  multipleSelect?: boolean;
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type FilterDefinitionFieldsModel = Record<string, FilterDefinitionFieldObjectMetadataModel<any>>;

export type YupTranslateErrorModel = {
  key: string;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  values?: any;
};

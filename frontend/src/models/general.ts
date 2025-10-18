export interface PageInfoModel {
  hasNextPage: boolean;
  hasPreviousPage: boolean;
  startCursor?: string;
  endCursor?: string;
}

export interface EdgeModel<T> {
  cursor: string;
  node: T;
}

export interface ConnectionModel<T> {
  edges?: EdgeModel<T>[];
  pageInfo: PageInfoModel;
}

export interface PaginationInputModel {
  first?: number;
  last?: number;
  before?: string;
  after?: string;
}

export type SortOrderModel = 'ASC' | 'DESC' | undefined;
export type SortOrderObjectModel<Keys extends string> = Partial<Record<Keys, SortOrderModel>>;
export const SortOrderAsc: SortOrderModel = 'ASC';
export const SortOrderDesc: SortOrderModel = 'DESC';
export const SortQueryParamName = 'sorts';

export interface SortOrderFieldModel {
  field: string;
  display: string;
}

export const FilterQueryParamName = 'filter';

export interface StringFilterModel {
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
  caseInsensitive?: boolean;
}

// Note: This is called "IntFilter" but
// it supports float !
export interface IntFilterModel {
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
}

export interface BooleanFilterModel {
  eq?: boolean;
  notEq?: boolean;
}

export interface DateFilterModel {
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
}

export interface FilterDefinitionFieldObjectMetadataModel<T> {
  display: string;
  description?: string;
  operations: FilterDefinitionOperationsModel<T>;
}

export type FilterDefinitionOperationsModel<T> = Record<string, FilterOperationMetadataModel<T>>;

export interface FilterDefinitionEnumObjectModel<T> {
  value: T;
  display: string;
  description?: string;
}

export interface FilterOperationMetadataModel<T> {
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
  caseInsensitiveEnabled?: boolean;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type FilterDefinitionFieldsModel = Record<string, FilterDefinitionFieldObjectMetadataModel<any>>;

export interface YupTranslateErrorModel {
  key: string;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  values?: any;
}

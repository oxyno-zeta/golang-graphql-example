/* eslint-disable @typescript-eslint/no-explicit-any */
export interface PredefinedFilter {
  filter: FilterValueObject;
  display: string;
  description?: string;
}

export interface LineOrGroup {
  type: string;
  key: string;
  // value?: undefined | null | any;
  initialValue: FieldInitialValueObject | BuilderInitialValueObject;
}

export interface BuilderInitialValueObject {
  items: LineOrGroup[];
  group: string;
}

export interface FieldInitialValueObject {
  field: string;
  operation: string;
  value: any;
}

export type FilterValueObject = {
  AND?: FilterValueObject[];
  OR?: FilterValueObject[];
} & Record<string, FieldOperationValueObject>;

export type FieldOperationValueObject = Record<string, any>;

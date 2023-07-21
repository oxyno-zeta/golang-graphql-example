import {
  FilterValueObject,
  BuilderInitialValueObject,
  LineOrGroup,
  FieldInitialValueObject,
  FieldOperationValueObject,
} from './types';

export function generateKey(prefix: string) {
  return `${prefix}-${(Math.random() + 1).toString(36).substring(2)}`;
}

// This is copied to avoid interaction with fields validation.
export function requiredInputValidate(value: undefined | null | string) {
  if (value === undefined || value === null || value === '') {
    return 'common.fieldValidationError.required';
  }

  // Default
  return null;
}

export function buildFilterBuilderInitialItems(
  initialValue: undefined | null | FilterValueObject,
  keyPrefix?: string,
): BuilderInitialValueObject {
  // Initialize result
  const res: BuilderInitialValueObject = {
    group: 'AND',
    items: [{ type: 'line', key: `${keyPrefix}root`, initialValue: buildFieldInitialValue(undefined)[0] }],
  };

  // Check if initial value is set
  if (initialValue === undefined || initialValue === null) {
    return res;
  }

  // Get initial value keys
  const keys = Object.keys(initialValue);

  // Check if it isn't an empty object
  if (keys.length === 0) {
    return res;
  }

  // Check if AND or OR is the only key
  if (keys.length === 1 && (keys[0] === 'OR' || keys[0] === 'AND')) {
    // Update group
    const [group] = keys;
    res.group = group;

    // Check if group have a value
    // Note: This is security for Typescript bypass cases or API coming filters that cannot be tested by TS.
    if (initialValue[group]) {
      res.items = (initialValue[group] as FilterValueObject[]).map((it, index) => {
        // Build key
        const key = keyPrefix + group + index;

        const value = buildFilterBuilderInitialItems(it, key);

        // Check if returned builder object is limited to 1 object
        // Optimize it
        if (value.items.length === 1) {
          const [v] = value.items;
          return v;
        }

        // Return a new group
        return {
          type: 'group',
          key,
          initialValue: value,
        };
      });

      // Check if result have only 1 value
      if (res.items.length === 1) {
        const [v] = res.items;
        // Check that value is a group
        if (v.type === 'group') {
          return v.initialValue as BuilderInitialValueObject;
        }
      }
    }

    return res;
  }

  // Update lines
  res.items = keys.reduce((arr: LineOrGroup[], key: string, index: number) => {
    const v: LineOrGroup[] = [];

    // Build line key
    const lineKey = keyPrefix + key + index;

    if (key === 'AND' || key === 'OR') {
      v.push({
        type: 'group',
        key: lineKey,
        initialValue: buildFilterBuilderInitialItems({ [key]: initialValue[key] } as FilterValueObject, lineKey),
      });
    } else {
      v.push(
        ...buildFieldInitialValue({ [key]: initialValue[key] }).map((i, indX) => ({
          type: 'line',
          key: lineKey + indX,
          initialValue: i,
        })),
      );
    }

    // Return
    return [...arr, ...v];
  }, []);

  // Default
  return res;
}

export function buildFieldInitialValue(
  input: undefined | null | Record<string, FieldOperationValueObject>,
): FieldInitialValueObject[] {
  // Initialize empty
  const empty = [{ field: '', operation: '', value: undefined }];

  // Check if input exists
  if (input === null || input === undefined) {
    return empty;
  }

  // Get input keys
  const keys = Object.keys(input);
  // Check size of keys
  if (keys.length === 0) {
    return empty;
  }

  // Create result
  const res: FieldInitialValueObject[] = [];

  // Loop over keys
  keys.forEach((key) => {
    // Get field data
    const fieldData = input[key];

    // Get operations
    const operations = Object.keys(fieldData);
    // Check if operations exists or the only one is case insensitive
    if (operations.length === 0 || (operations.length === 1 && operations[0] === 'caseInsensitive')) {
      res.push({
        field: key,
        operation: '',
        value: undefined,
      });
      // Stop
      return;
    }

    // Loop over operations
    operations.forEach((operation) => {
      if (operation === 'caseInsensitive') {
        // Stop
        return;
      }

      // Save
      res.push({
        field: key,
        operation,
        value: fieldData[operation],
      });
    });
  });

  return res;
}

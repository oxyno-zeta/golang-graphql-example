import { StringFilterModel } from '../../../models/general';

export function onMainSearchChangeContains<T>(
  filterKey: keyof Omit<T, 'AND' | 'OR'>,
  newValue: string,
  oldValue: string,
  setFilter: (f: (input: T) => T) => void,
) {
  return onMainSearchChangeGeneric<T>(filterKey, 'contains', newValue, oldValue, setFilter);
}

export function getMainSearchInitialValueContains<T>(filterKey: keyof Omit<T, 'AND' | 'OR'>, filter: T) {
  return getMainSearchInitialValueGeneric<T, string>(filterKey, 'contains', '', filter);
}

/* eslint-disable @typescript-eslint/no-explicit-any */
export function getMainSearchInitialValueGeneric<T, R>(
  filterKey: keyof Omit<T, 'AND' | 'OR'>,
  filterOperation: string,
  defaultValue: R,
  filter: any,
) {
  // Check if it is at root level
  if (filter && filter[filterKey] && (filter[filterKey] as any)[filterOperation]) {
    return (filter[filterKey] as any)[filterOperation] as R;
  }

  // Check if the AND case exists
  if (filter && filter.AND) {
    const v = filter.AND.find((it: any) => it[filterKey] && it[filterKey][filterOperation]);
    // Check if v exists
    if (v) {
      return (v[filterKey] as any)[filterOperation] as R;
    }
  }

  // Default
  return defaultValue;
}

export function onMainSearchChangeGeneric<T>(
  filterKey: keyof Omit<T, 'AND' | 'OR'>,
  filterOperation: string,
  newValue: string,
  oldValue: string,
  setFilter: (f: (input: T) => T) => void,
) {
  // Call set filter
  setFilter((initialFilter: T) => {
    // Create a deep copy in order to force a reload of graphql query.
    // Otherwise, it is ignored
    const filterCopy = JSON.parse(JSON.stringify(initialFilter));
    // Check if [filterKey] [filterOperation] is at root level
    if (
      filterCopy &&
      filterCopy[filterKey] &&
      filterCopy[filterKey][filterOperation] &&
      filterCopy[filterKey][filterOperation] === oldValue
    ) {
      // Flush filter
      delete filterCopy[filterKey];
      // Check if it isn't a flush case
      if (newValue !== '') {
        // Set filter
        filterCopy[filterKey] = { [filterOperation]: newValue, caseInsensitive: true } as StringFilterModel;
      }

      // Save
      return { ...filterCopy };
    }

    // Check if filter exists and AND also
    if (filterCopy && filterCopy.AND) {
      // Check if it is a clean case
      if (newValue === '') {
        // Filter on all elements and keep only elements that aren't equal to the main search
        const newAnd = filterCopy.AND.filter(
          (it: any) =>
            !(it[filterKey] && it[filterKey][filterOperation] && it[filterKey][filterOperation] === oldValue),
        );

        // Check if and arrays are different
        if (newAnd.length !== filterCopy.AND.length) {
          // Replace array
          filterCopy.AND = newAnd;
          // Save
          return { ...filterCopy };
        }
      } else {
        // Search if filter have been provided
        const item = filterCopy.AND.find((it: any) => {
          if (it[filterKey] && it[filterKey][filterOperation] && it[filterKey][filterOperation] === oldValue) {
            return it;
          }

          return null;
        });

        // Check if item exists
        if (item) {
          item[filterKey] = { [filterOperation]: newValue, caseInsensitive: true };
          return { ...filterCopy };
        }

        filterCopy.AND.push({ [filterKey]: { [filterOperation]: newValue, caseInsensitive: true } });
        return { ...filterCopy };
      }
    }

    // Check if it a clean
    if (newValue === '') {
      // Stop here as we haven't found anything before
      // Return the initial filter as nothing changed
      return filterCopy;
    }

    if (filterCopy && Object.keys(filterCopy).length >= 1) {
      filterCopy.AND = [{ [filterKey]: { [filterOperation]: newValue, caseInsensitive: true } }];
      return { ...filterCopy };
    }

    return { [filterKey]: { [filterOperation]: newValue, caseInsensitive: true } };
  });
}

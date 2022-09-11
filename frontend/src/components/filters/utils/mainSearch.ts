import { StringFilterModel } from '../../../models/general';

/* eslint-disable @typescript-eslint/no-explicit-any */
// eslint-disable-next-line import/prefer-default-export
export function onMainSearchChangeDefault<T>(
  filterKey: keyof Omit<T, 'AND' | 'OR'>,
  newValue: string,
  oldValue: string,
  setFilter: (f: (input: T) => T) => void,
) {
  // Call set filter
  setFilter((initialFilter: T) => {
    // Create a deep copy in order to force a reload of graphql query.
    // Otherwise, it is ignored
    const filterCopy = JSON.parse(JSON.stringify(initialFilter));
    // Check if [filterKey] contains is at root level
    if (
      filterCopy &&
      filterCopy[filterKey] &&
      filterCopy[filterKey].contains &&
      filterCopy[filterKey].contains === oldValue
    ) {
      // Flush filter
      delete filterCopy[filterKey];
      // Check if it isn't a flush case
      if (newValue !== '') {
        // Set filter
        filterCopy[filterKey] = { contains: newValue } as StringFilterModel;
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
          (it: any) => !(it[filterKey] && it[filterKey].contains && it[filterKey].contains === oldValue),
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
          if (it[filterKey] && it[filterKey].contains && it[filterKey].contains === oldValue) {
            return it;
          }

          return null;
        });

        // Check if item exists
        if (item) {
          item[filterKey] = { contains: newValue };
          return { ...filterCopy };
        }

        filterCopy.AND.push({ [filterKey]: { contains: newValue } });
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
      filterCopy.AND = [{ [filterKey]: { contains: newValue } }];
      return { ...filterCopy };
    }

    return { [filterKey]: { contains: newValue } };
  });
}

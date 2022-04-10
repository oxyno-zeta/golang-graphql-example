import { TodoFilterModel } from '../../models/todos';

// eslint-disable-next-line import/prefer-default-export
export function onMainSearchChange(
  newValue: string,
  oldValue: string,
  setFilter: (f: (input: TodoFilterModel) => TodoFilterModel) => void,
) {
  // Call set filter
  setFilter((initialFilter) => {
    // Create a deep copy in order to force a reload of graphql query.
    // Otherwise, it is ignored
    const filterCopy: TodoFilterModel = JSON.parse(JSON.stringify(initialFilter));
    // Check if text contains is at root level
    if (filterCopy && filterCopy.text && filterCopy.text.contains && filterCopy.text.contains === oldValue) {
      // Flush filter
      delete filterCopy.text;
      // Check if it isn't a flush case
      if (newValue !== '') {
        // Set filter
        filterCopy.text = { contains: newValue };
      }

      // Save
      return { ...filterCopy };
    }

    // Check if filter exists and AND also
    if (filterCopy && filterCopy.AND) {
      // Check if it is a clean case
      if (newValue === '') {
        // Filter on all elements and keep only elements that aren't equal to the main search
        const newAnd = filterCopy.AND.filter((it) => !(it.text && it.text.contains && it.text.contains === oldValue));

        // Check if and arrays are different
        if (newAnd.length !== filterCopy.AND.length) {
          // Replace array
          filterCopy.AND = newAnd;
          // Save
          return { ...filterCopy };
        }
      } else {
        // Search if filter have been provided
        const item = filterCopy.AND.find((it) => {
          if (it.text && it.text.contains && it.text.contains === oldValue) {
            return it;
          }

          return null;
        });

        // Check if item exists
        if (item) {
          item.text = { contains: newValue };
          return { ...filterCopy };
        }

        filterCopy.AND.push({ text: { contains: newValue } });
        return { ...filterCopy };
      }
    }

    // Check if it a clean
    if (newValue === '') {
      // Stop here as we haven't found anything before
      // Return the initial filter as nothing changed
      return initialFilter;
    }

    if (filterCopy && Object.keys(filterCopy).length >= 1) {
      filterCopy.AND = [{ text: { contains: newValue } }];
      return { ...filterCopy };
    }

    return { text: { contains: newValue } };
  });
}

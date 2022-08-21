import { GridSortModel, GridColDef } from '@mui/x-data-grid';
import { SortOrderAsc, SortOrderDesc, SortOrderModel } from '../../../models/general';

export function buildMUIXSort(sort: Record<string, SortOrderModel>, columns: GridColDef[]): GridSortModel {
  // Get all keys from sort object
  const keys = Object.keys(sort);

  // Build fields
  const fields = columns.filter((it) => it.sortable).map((it) => it.field);

  // Initialize result
  const res: GridSortModel = [];

  // Loop over keys
  for (let i = 0; i < keys.length; i += 1) {
    // Save key
    const key = keys[i];

    // Check if key have a value
    if (!sort[key]) {
      // Ignore that field
      continue; // eslint-disable-line
    }

    // Check if key isn't in supported fields
    if (!fields.includes(key)) {
      // Stop here => sort is on non supported fields
      return [];
    }

    // Save sort
    res.push({ field: key, sort: sort[key] === SortOrderAsc ? 'asc' : 'desc' });
  }

  // Limit to 1 because free version
  if (res.length > 1) {
    // In this case, as more fields are managed, ignore the sort here
    return [];
  }

  return res;
}

export function setMUIXSortBuilder(setSort: (data: Record<string, SortOrderModel>) => void) {
  return (input: GridSortModel) => {
    // Init sort object
    const sort: Record<string, SortOrderModel> = {};
    // Build sort object
    input.forEach((it) => {
      sort[it.field] = it.sort === 'asc' ? SortOrderAsc : SortOrderDesc;
    });

    // Save
    setSort(sort);
  };
}

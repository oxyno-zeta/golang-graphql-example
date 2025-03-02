import { GridSortModel, GridColDef, GridSortItem } from '@mui/x-data-grid';
import { SortOrderAsc, SortOrderDesc, SortOrderModel } from '../../../models/general';

export function buildMUIXSort(sorts: Record<string, SortOrderModel>[], columns: GridColDef[]): GridSortModel {
  // Limit to 1 because free version
  if (sorts.length > 1) {
    // In this case, as more fields are managed, ignore the sort here
    return [];
  }

  // Build fields
  const fields = columns.filter((it) => it.sortable).map((it) => it.field);

  // Loop over sorts
  const res: GridSortModel = sorts.reduce((accu, sort) => {
    // Get first key that have a value and is a supported field
    const key = Object.keys(sort)
      // Check if key isn't in supported fields
      .filter((it) => fields.includes(it))
      // Get first key that have a value
      .find((it) => !!sort[it]);
    // Check if such a key haven't been found
    if (!key) {
      // Nothing to save
      return accu;
    }

    // Save sort
    accu.push({ field: key, sort: sort[key] === SortOrderAsc ? 'asc' : 'desc' });
    // Default
    return accu;
  }, [] as GridSortItem[]);

  // Limit to 1 because free version
  // That will also manage the case where the same field is put 2 times into the sorts list
  if (res.length > 1) {
    // In this case, as more fields are managed, ignore the sort here
    return [];
  }

  return res;
}

export function setMUIXSortBuilder(setSorts: (data: Record<string, SortOrderModel>[]) => void) {
  return (input: GridSortModel) => {
    // Build sort list
    const sorts: Record<string, SortOrderModel>[] = input.map((it) => ({
      [it.field]: it.sort === 'asc' ? SortOrderAsc : SortOrderDesc,
    }));

    // Save
    setSorts(sorts);
  };
}

// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import { GridColDef, GridSortModel } from '@mui/x-data-grid';
import { SortOrderModel } from '../../../models/general';
import { buildMUIXSort, setMUIXSortBuilder } from './index';

describe('buildMUIXSort', () => {
  describe('two sortable columns', () => {
    const columns: GridColDef[] = [
      { field: 'fake', sortable: true },
      { field: 'field', sortable: true },
    ];

    test('should return an empty result when input is an empty list', () => {
      const res = buildMUIXSort([], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is a empty object list', () => {
      const res = buildMUIXSort([{}], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is a simple item with not supported field', () => {
      const res = buildMUIXSort([{ fake2: 'ASC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is 2 simple item with not supported field', () => {
      const res = buildMUIXSort([{ fake3: 'ASC' }, { fake2: 'ASC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is a simple item with supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }], columns);
      const expected: GridSortModel = [{ field: 'fake', sort: 'asc' }];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 2 items with same supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }, { fake: 'DESC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 2 items with different supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }, { field: 'DESC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 1 item with supported field and empty object', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }, {}], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 1 item with 2 supported fields (second field must be ignored)', () => {
      const res = buildMUIXSort([{ fake: 'ASC', field: 'ASC' }], columns);
      const expected: GridSortModel = [{ field: 'fake', sort: 'asc' }];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 1 item with supported field and 1 item with empty supported field object', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }, { field: undefined }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 1 item with empty supported field and valued supported field', () => {
      const res = buildMUIXSort([{ field: undefined, fake: 'ASC' }], columns);
      const expected: GridSortModel = [{ field: 'fake', sort: 'asc' }];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 1 item with valued supported field and empty supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC', field: undefined }], columns);
      const expected: GridSortModel = [{ field: 'fake', sort: 'asc' }];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 1 item with not supported field and valued supported field', () => {
      const res = buildMUIXSort([{ fake2: 'DESC', fake: 'ASC' }], columns);
      const expected: GridSortModel = [{ field: 'fake', sort: 'asc' }];
      expect(res).toEqual(expected);
    });
  });

  describe('simple sortable column', () => {
    const columns: GridColDef[] = [{ field: 'fake', sortable: true }];

    test('should return an empty result when input is an empty list', () => {
      const res = buildMUIXSort([], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is a empty object list', () => {
      const res = buildMUIXSort([{}], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is a simple item with not supported field', () => {
      const res = buildMUIXSort([{ fake2: 'ASC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is 2 simple item with not supported field', () => {
      const res = buildMUIXSort([{ fake3: 'ASC' }, { fake2: 'ASC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is a simple item with supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }], columns);
      const expected: GridSortModel = [{ field: 'fake', sort: 'asc' }];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 2 items with same supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }, { fake: 'DESC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should be ok when input is 1 item with supported field and empty object', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }, {}], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
  });

  describe('simple non sortable column', () => {
    const columns: GridColDef[] = [{ field: 'fake' }];

    test('should return an empty result when input is an empty list', () => {
      const res = buildMUIXSort([], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is 1 empty object list', () => {
      const res = buildMUIXSort([{}], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is 1 simple item with not supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is 2 simple item with not supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }, { fake2: 'ASC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
  });

  describe('empty columns', () => {
    const columns: GridColDef[] = [];

    test('should return an empty result when input is an empty list', () => {
      const res = buildMUIXSort([], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is 1 empty object list', () => {
      const res = buildMUIXSort([{}], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is 1 simple item with not supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
    test('should return an empty result when input is 2 simple item with not supported field', () => {
      const res = buildMUIXSort([{ fake: 'ASC' }, { fake2: 'ASC' }], columns);
      const expected: GridSortModel = [];
      expect(res).toEqual(expected);
    });
  });
});

describe('setMUIXSortBuilder', () => {
  test('should be ok when sort is empty', () => {
    let res: Record<string, SortOrderModel>[] = [];
    const expected: Record<string, SortOrderModel>[] = [];

    setMUIXSortBuilder((data) => {
      res = data;
    })([]);

    expect(res).toEqual(expected);
  });
  test('should be ok when sort have 1 item', () => {
    let res: Record<string, SortOrderModel>[] = [];
    const expected: Record<string, SortOrderModel>[] = [{ fake: 'ASC' }];

    setMUIXSortBuilder((data) => {
      res = data;
    })([{ field: 'fake', sort: 'asc' }]);

    expect(res).toEqual(expected);
  });
  test('should be ok when sort have 2 items', () => {
    let res: Record<string, SortOrderModel>[] = [];
    const expected: Record<string, SortOrderModel>[] = [{ fake: 'ASC' }, { fake2: 'DESC' }];

    setMUIXSortBuilder((data) => {
      res = data;
    })([
      { field: 'fake', sort: 'asc' },
      { field: 'fake2', sort: 'desc' },
    ]);

    expect(res).toEqual(expected);
  });
});

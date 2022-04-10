import { mockRandomForEach } from 'jest-mock-random';
import { buildFieldInitialValue, buildFilterBuilderInitialItems } from './utils';
import { BuilderInitialValueObject } from './types';

describe('buildFilterBuilderInitialItems', () => {
  // Random mock
  mockRandomForEach([0.1]);
  const key = '3llllllllm';

  test('should return an empty result when input is undefined', () => {
    const res = buildFilterBuilderInitialItems(undefined);
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: '', operation: '', value: undefined },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should return an empty result when input is null', () => {
    const res = buildFilterBuilderInitialItems(null);
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: '', operation: '', value: undefined },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should return an empty result when input is an AND array an empty array', () => {
    const res = buildFilterBuilderInitialItems({ AND: [] });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [],
    };
    expect(res).toEqual(expected);
  });
  test('should return an empty result when input is an OR array an empty array', () => {
    const res = buildFilterBuilderInitialItems({ OR: [] });
    const expected: BuilderInitialValueObject = {
      group: 'OR',
      items: [],
    };
    expect(res).toEqual(expected);
  });
  test('should return an empty result when input is an empty object', () => {
    const res = buildFilterBuilderInitialItems({});
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: '', operation: '', value: undefined },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple AND array with 1 value', () => {
    const res = buildFilterBuilderInitialItems({
      AND: [{ f1: { eq: 'val1' } }],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple OR array with 1 value', () => {
    const res = buildFilterBuilderInitialItems({
      OR: [{ f1: { eq: 'val1' } }],
    });
    const expected: BuilderInitialValueObject = {
      group: 'OR',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple AND array with multiple values', () => {
    const res = buildFilterBuilderInitialItems({
      AND: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple OR array with multiple values', () => {
    const res = buildFilterBuilderInitialItems({
      OR: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
    });
    const expected: BuilderInitialValueObject = {
      group: 'OR',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple OR array with multiple values and fields', () => {
    const res = buildFilterBuilderInitialItems({
      OR: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
      f3: { notEq2: 'val2' },
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'group',
          initialValue: {
            group: 'OR',
            items: [
              {
                key,
                type: 'line',
                initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
              },
              {
                key,
                type: 'line',
                initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
              },
            ],
          },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f3', operation: 'notEq2', value: 'val2' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok when input is a simple AND array with nested AND without any other field', () => {
    const res = buildFilterBuilderInitialItems({
      AND: [{ AND: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }] }],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok when input is a simple AND array nested OR without any other field', () => {
    const res = buildFilterBuilderInitialItems({
      AND: [{ OR: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }] }],
    });
    const expected: BuilderInitialValueObject = {
      group: 'OR',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok when input is a simple AND array nested OR and AND on 2 objects without any other field', () => {
    const res = buildFilterBuilderInitialItems({
      AND: [
        { OR: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }] },
        { AND: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }] },
      ],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'group',
          initialValue: {
            group: 'OR',
            items: [
              {
                key,
                type: 'line',
                initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
              },
              {
                key,
                type: 'line',
                initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
              },
            ],
          },
        },
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              {
                key,
                type: 'line',
                initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
              },
              {
                key,
                type: 'line',
                initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
              },
            ],
          },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok when input is a simple AND array nested OR and AND on same object without any other field', () => {
    const res = buildFilterBuilderInitialItems({
      AND: [
        {
          OR: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
          AND: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
        },
      ],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'group',
          initialValue: {
            group: 'OR',
            items: [
              {
                key,
                type: 'line',
                initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
              },
              {
                key,
                type: 'line',
                initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
              },
            ],
          },
        },
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              {
                key,
                type: 'line',
                initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
              },
              {
                key,
                type: 'line',
                initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
              },
            ],
          },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok when input is a simple OR array nested OR and AND on same object without any other field', () => {
    const res = buildFilterBuilderInitialItems({
      OR: [
        {
          OR: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
          AND: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
        },
      ],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'group',
          initialValue: {
            group: 'OR',
            items: [
              {
                key,
                type: 'line',
                initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
              },
              {
                key,
                type: 'line',
                initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
              },
            ],
          },
        },
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              {
                key,
                type: 'line',
                initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
              },
              {
                key,
                type: 'line',
                initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
              },
            ],
          },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok when input is a OR and AND array nested OR and AND on same object without any other field', () => {
    const res = buildFilterBuilderInitialItems({
      OR: [
        {
          OR: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
          AND: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
        },
      ],
      AND: [
        {
          OR: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
          AND: [{ f1: { eq: 'val1' } }, { f2: { notEq: 'val1' } }],
        },
      ],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              {
                key,
                type: 'group',
                initialValue: {
                  group: 'OR',
                  items: [
                    {
                      key,
                      type: 'line',
                      initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
                    },
                    {
                      key,
                      type: 'line',
                      initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
                    },
                  ],
                },
              },
              {
                key,
                type: 'group',
                initialValue: {
                  group: 'AND',
                  items: [
                    {
                      key,
                      type: 'line',
                      initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
                    },
                    {
                      key,
                      type: 'line',
                      initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
                    },
                  ],
                },
              },
            ],
          },
        },
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              {
                key,
                type: 'group',
                initialValue: {
                  group: 'OR',
                  items: [
                    {
                      key,
                      type: 'line',
                      initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
                    },
                    {
                      key,
                      type: 'line',
                      initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
                    },
                  ],
                },
              },
              {
                key,
                type: 'group',
                initialValue: {
                  group: 'AND',
                  items: [
                    {
                      key,
                      type: 'line',
                      initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
                    },
                    {
                      key,
                      type: 'line',
                      initialValue: { field: 'f2', operation: 'notEq', value: 'val1' },
                    },
                  ],
                },
              },
            ],
          },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple object with 1 field with 1 operation', () => {
    const res = buildFilterBuilderInitialItems({
      f1: { eq: 'val1' },
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple object with 2 fields with 1 operation', () => {
    const res = buildFilterBuilderInitialItems({
      f1: { eq: 'val1' },
      f2: { eq: 'val2' },
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f2', operation: 'eq', value: 'val2' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple object with 1 field with multiple operations', () => {
    const res = buildFilterBuilderInitialItems({
      f1: { eq: 'val1', notEq: 'val2' },
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'notEq', value: 'val2' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a simple object with 2 field with multiple operations', () => {
    const res = buildFilterBuilderInitialItems({
      f1: { eq: 'val1', notEq: 'val2' },
      f2: { eq2: 'val12', notEq2: 'val22' },
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'eq', value: 'val1' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f1', operation: 'notEq', value: 'val2' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f2', operation: 'eq2', value: 'val12' },
        },
        {
          key,
          type: 'line',
          initialValue: { field: 'f2', operation: 'notEq2', value: 'val22' },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a OR array with 2 field with multiple operations', () => {
    const res = buildFilterBuilderInitialItems({
      OR: [{ f1: { eq: 'val1', notEq2: 'valnotEq2' } }, { f2: { notEq: 'val1' }, f3: { isNull: true, eq: 'valEqf3' } }],
    });
    const expected: BuilderInitialValueObject = {
      group: 'OR',
      items: [
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              { key, type: 'line', initialValue: { field: 'f1', operation: 'eq', value: 'val1' } },
              { key, type: 'line', initialValue: { field: 'f1', operation: 'notEq2', value: 'valnotEq2' } },
            ],
          },
        },
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              { key, type: 'line', initialValue: { field: 'f2', operation: 'notEq', value: 'val1' } },
              { key, type: 'line', initialValue: { field: 'f3', operation: 'isNull', value: true } },
              { key, type: 'line', initialValue: { field: 'f3', operation: 'eq', value: 'valEqf3' } },
            ],
          },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a AND array with 2 field with multiple operations', () => {
    const res = buildFilterBuilderInitialItems({
      AND: [
        { f1: { eq: 'val1', notEq2: 'valnotEq2' } },
        { f2: { notEq: 'val1' }, f3: { isNull: true, eq: 'valEqf3' } },
      ],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              { key, type: 'line', initialValue: { field: 'f1', operation: 'eq', value: 'val1' } },
              { key, type: 'line', initialValue: { field: 'f1', operation: 'notEq2', value: 'valnotEq2' } },
            ],
          },
        },
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              { key, type: 'line', initialValue: { field: 'f2', operation: 'notEq', value: 'val1' } },
              { key, type: 'line', initialValue: { field: 'f3', operation: 'isNull', value: true } },
              { key, type: 'line', initialValue: { field: 'f3', operation: 'eq', value: 'valEqf3' } },
            ],
          },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
  test('should be ok result when input is a AND/OR array with 2 field with multiple operations', () => {
    const res = buildFilterBuilderInitialItems({
      OR: [{ f1: { eq: 'val1', notEq2: 'valnotEq2' } }, { f2: { notEq: 'val1' }, f3: { isNull: true, eq: 'valEqf3' } }],
      AND: [
        { f1: { eq: 'val1', notEq2: 'valnotEq2' } },
        { f2: { notEq: 'val1' }, f3: { isNull: true, eq: 'valEqf3' } },
      ],
    });
    const expected: BuilderInitialValueObject = {
      group: 'AND',
      items: [
        {
          key,
          type: 'group',
          initialValue: {
            group: 'OR',
            items: [
              {
                key,
                type: 'group',
                initialValue: {
                  group: 'AND',
                  items: [
                    { key, type: 'line', initialValue: { field: 'f1', operation: 'eq', value: 'val1' } },
                    { key, type: 'line', initialValue: { field: 'f1', operation: 'notEq2', value: 'valnotEq2' } },
                  ],
                },
              },
              {
                key,
                type: 'group',
                initialValue: {
                  group: 'AND',
                  items: [
                    { key, type: 'line', initialValue: { field: 'f2', operation: 'notEq', value: 'val1' } },
                    { key, type: 'line', initialValue: { field: 'f3', operation: 'isNull', value: true } },
                    { key, type: 'line', initialValue: { field: 'f3', operation: 'eq', value: 'valEqf3' } },
                  ],
                },
              },
            ],
          },
        },
        {
          key,
          type: 'group',
          initialValue: {
            group: 'AND',
            items: [
              {
                key,
                type: 'group',
                initialValue: {
                  group: 'AND',
                  items: [
                    { key, type: 'line', initialValue: { field: 'f1', operation: 'eq', value: 'val1' } },
                    { key, type: 'line', initialValue: { field: 'f1', operation: 'notEq2', value: 'valnotEq2' } },
                  ],
                },
              },
              {
                key,
                type: 'group',
                initialValue: {
                  group: 'AND',
                  items: [
                    { key, type: 'line', initialValue: { field: 'f2', operation: 'notEq', value: 'val1' } },
                    { key, type: 'line', initialValue: { field: 'f3', operation: 'isNull', value: true } },
                    { key, type: 'line', initialValue: { field: 'f3', operation: 'eq', value: 'valEqf3' } },
                  ],
                },
              },
            ],
          },
        },
      ],
    };
    expect(res).toEqual(expected);
  });
});

describe('buildFieldInitialValue', () => {
  test('should be ok result when input is undefined', () => {
    const res = buildFieldInitialValue(undefined);
    expect(res).toEqual([{ field: '', operation: '', value: undefined }]);
  });
  test('should be ok result when input is null', () => {
    const res = buildFieldInitialValue(null);
    expect(res).toEqual([{ field: '', operation: '', value: undefined }]);
  });
  test('should be ok result when input is an empty object', () => {
    const res = buildFieldInitialValue({});
    expect(res).toEqual([{ field: '', operation: '', value: undefined }]);
  });
  test('should be ok result when input is an empty second object', () => {
    const res = buildFieldInitialValue({ key: {} });
    expect(res).toEqual([{ field: 'key', operation: '', value: undefined }]);
  });
  test('should be ok result when input is normal (string)', () => {
    const res = buildFieldInitialValue({ key: { key2: 'string' } });
    expect(res).toEqual([{ field: 'key', operation: 'key2', value: 'string' }]);
  });
  test('should be ok result when input is normal (string[])', () => {
    const res = buildFieldInitialValue({ key: { key2: ['string'] } });
    expect(res).toEqual([{ field: 'key', operation: 'key2', value: ['string'] }]);
  });
  test('should be ok result when input is normal ([boolean])', () => {
    const res = buildFieldInitialValue({ key: { key2: [true] } });
    expect(res).toEqual([{ field: 'key', operation: 'key2', value: [true] }]);
  });
  test('should be ok result when input is normal (boolean)', () => {
    const res = buildFieldInitialValue({ key: { key2: true } });
    expect(res).toEqual([{ field: 'key', operation: 'key2', value: true }]);
  });
  test('should be ok result when input is normal with 2 operations (string and boolean)', () => {
    const res = buildFieldInitialValue({ key: { key2: true, eq: 'fake' } });
    expect(res).toEqual([
      { field: 'key', operation: 'key2', value: true },
      { field: 'key', operation: 'eq', value: 'fake' },
    ]);
  });
});

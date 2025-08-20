// Copied from https://github.com/sindresorhus/array-move/blob/main/index.js
// eslint-disable-next-line import-x/prefer-default-export
export function arrayMoveItem<T>(array: T[], fromIndex: number, toIndex: number) {
  const startIndex = fromIndex < 0 ? array.length + fromIndex : fromIndex;

  if (startIndex >= 0 && startIndex < array.length) {
    const endIndex = toIndex < 0 ? array.length + toIndex : toIndex;

    const [item] = array.splice(fromIndex, 1);
    array.splice(endIndex, 0, item);
  }
}

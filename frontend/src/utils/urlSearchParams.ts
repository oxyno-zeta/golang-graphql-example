import type { URLSearchParams } from 'url';
import type { URLSearchParamsInit } from 'react-router';

export function getAllSearchParams(searchParams: URLSearchParams): Record<string, string> {
  // Initial params
  const params: Record<string, string> = {};

  // Loop over params
  searchParams.forEach((value, name) => {
    params[name] = value;
  });

  return params;
}

export function addJSONObjectSearchParam(
  key: string,
  data: Record<string, any>, // eslint-disable-line @typescript-eslint/no-explicit-any
  searchParams: URLSearchParams,
) {
  // Stringify data
  const objStr = JSON.stringify(data);

  // Save data
  searchParams.set(key, objStr);

  return searchParams;
}

export function setJSONObjectSearchParam(
  key: string,
  data: Record<string, any>, // eslint-disable-line @typescript-eslint/no-explicit-any
  searchParams: URLSearchParams,
  setSearchParams: (
    nextInit: URLSearchParamsInit,
    navigateOptions?:
      | {
          replace?: boolean | undefined;
          state?: any; // eslint-disable-line @typescript-eslint/no-explicit-any
        }
      | undefined,
  ) => void,
) {
  // Stringify data
  const objStr = JSON.stringify(data);

  // Save data
  searchParams.set(key, objStr);

  // Save search params
  setSearchParams(searchParams as URLSearchParamsInit);
}

export function getJSONObjectFromSearchParam<T>(key: string, init: T, searchParams: URLSearchParams): T {
  // Get key
  const objStr = searchParams.get(key);

  // Check if object string is null or empty string
  if (objStr === null || objStr === '') {
    // Return init
    return init;
  }

  // Try to parse object
  try {
    // Parse
    const obj = JSON.parse(objStr || '');

    // Return
    return obj;
  } catch {
    // Return init
    return init;
  }
}

export function deleteAndSetSearchParam(
  key: string,
  searchParams: URLSearchParams,
  setSearchParams: (
    nextInit: URLSearchParamsInit,
    navigateOptions?:
      | {
          replace?: boolean | undefined;
          state?: any; // eslint-disable-line @typescript-eslint/no-explicit-any
        }
      | undefined,
  ) => void,
) {
  deleteAndSetSearchParams([key], searchParams, setSearchParams);
}

export function deleteAndSetSearchParams(
  keys: string[],
  searchParams: URLSearchParams,
  setSearchParams: (
    nextInit: URLSearchParamsInit,
    navigateOptions?:
      | {
          replace?: boolean | undefined;
          state?: any; // eslint-disable-line @typescript-eslint/no-explicit-any
        }
      | undefined,
  ) => void,
) {
  // Loop over keys
  keys.forEach((key) => {
    // Delete param
    searchParams.delete(key);
  });

  // Clean all
  setSearchParams(searchParams as URLSearchParamsInit);
}

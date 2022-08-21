import { URLSearchParams } from 'url';
import { URLSearchParamsInit } from 'react-router-dom';

// eslint-disable-next-line import/prefer-default-export
export function getAllSearchParams(searchParams: URLSearchParams): Record<string, string> {
  // Initial params
  const params: Record<string, string> = {};

  // Loop over params
  searchParams.forEach((value, name) => {
    params[name] = value;
  });

  return params;
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
  setSearchParams(getAllSearchParams(searchParams));
}

export function getJSONObjectFromSearchParam<T>(
  key: string,
  init: T,
  searchParams: URLSearchParams,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
): T {
  // Get key
  const objStr = searchParams.get(key);

  // Check if object string is null or empty string
  if (objStr === null && objStr === '') {
    // Return init
    return init;
  }

  // Try to parse object
  try {
    // Parse
    const obj = JSON.parse(objStr || '');

    // Check if it is an empty object
    if (Object.keys(obj).length === 0) {
      // Return init
      return init;
    }

    // Return
    return obj;
  } catch (e) {
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
  setSearchParams(getAllSearchParams(searchParams));
}

import { URLSearchParamsInit } from 'react-router-dom';
import { URLSearchParams } from 'url';
import { PaginationInputModel } from '../models/general';
import { getAllSearchParams } from './urlSearchParams';

export function cleanPaginationSearchParams(searchParams: URLSearchParams): URLSearchParams {
  // Delete pagination params
  searchParams.delete('first');
  searchParams.delete('last');
  searchParams.delete('before');
  searchParams.delete('after');

  return searchParams;
}

export function cleanAndSetCleanedPagination(
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
  // Delete pagination params
  cleanPaginationSearchParams(searchParams);

  // Clean all
  setSearchParams(getAllSearchParams(searchParams));
}

export function getPaginationFromSearchParams(
  initPagination: PaginationInputModel,
  maxPagination: number,
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
): PaginationInputModel {
  // Get first
  const firstStr = searchParams.get('first') || '';
  // Get last
  const lastStr = searchParams.get('last') || '';
  // Get before
  const before = searchParams.get('before') || '';
  // Get after
  const after = searchParams.get('after') || '';

  // Check pagination is empty
  if (firstStr === '' && lastStr === '' && before === '' && after === '') {
    // Return init
    return initPagination;
  }

  // Check if all fields are correctly used
  if (!((lastStr !== '' && before !== '') || (firstStr !== '' && after !== ''))) {
    // => Clean all
    // => Return init pagination
    cleanAndSetCleanedPagination(searchParams, setSearchParams);

    return initPagination;
  }

  // Try to parse first
  if (firstStr !== '' && after !== '') {
    try {
      // Parsed first
      let first = parseInt(firstStr as string, 10);
      // Check if value is greater than max pagination or lower than 0
      if (first > maxPagination || first <= 0) {
        first = maxPagination;
      }

      return {
        first,
        after,
      };
    } catch (e) {
      // Cannot be parsed
      // => Clean all
      // => Return init pagination
      cleanAndSetCleanedPagination(searchParams, setSearchParams);

      return initPagination;
    }
  }

  // Try to parse last
  if (lastStr !== '' && before !== '') {
    try {
      // Parsed last
      let last = parseInt(lastStr as string, 10);
      // Check if value is greater than max pagination or lower than 0
      if (last > maxPagination || last <= 0) {
        last = maxPagination;
      }

      return {
        last,
        before,
      };
    } catch (e) {
      // Cannot be parsed
      // => Clean all
      // => Return init pagination
      cleanAndSetCleanedPagination(searchParams, setSearchParams);

      return initPagination;
    }
  }

  // Default case
  return initPagination;
}

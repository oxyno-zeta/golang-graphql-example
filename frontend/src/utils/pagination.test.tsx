// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import { URLSearchParams } from 'url';
import { cleanAndSetCleanedPagination, cleanPaginationSearchParams, getPaginationFromSearchParams } from './pagination';

describe('utils/pagination', () => {
  describe('getPaginationFromSearchParams', () => {
    it('should return init pagination when nothing is found', () => {
      const input = new URLSearchParams();
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 1 });
    });

    it('should return init when wrong fields and params should be cleaned (1)', () => {
      const input = new URLSearchParams();
      input.set('last', '1');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = {};

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 1 });
      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should return init when wrong fields and params should be cleaned (2)', () => {
      const input = new URLSearchParams();
      input.set('first', '1');
      input.set('before', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = {};

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 1 });
      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should return init when only first is present', () => {
      const input = new URLSearchParams();
      input.set('first', '2');
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = {};

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 1 });
      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should return init when only last is present', () => {
      const input = new URLSearchParams();
      input.set('last', '2');
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = {};

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 1 });
      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should return init when number param is incorrect (1)', () => {
      const input = new URLSearchParams();
      input.set('first', 'fail');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = {};

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 1 });
      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should return init when number param is incorrect (2)', () => {
      const input = new URLSearchParams();
      input.set('last', 'fail');
      input.set('before', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = {};

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 1 });
      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should return a valid limited pagination (first case) (1)', () => {
      const input = new URLSearchParams();
      input.set('first', '100');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 5, after: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid limited pagination (first case) (2)', () => {
      const input = new URLSearchParams();
      input.set('first', '-100');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 5, after: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid limited pagination (last case) (1)', () => {
      const input = new URLSearchParams();
      input.set('last', '100');
      input.set('before', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ last: 5, before: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid limited pagination (last case) (2)', () => {
      const input = new URLSearchParams();
      input.set('last', '-100');
      input.set('before', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ last: 5, before: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid pagination (with first and after)', () => {
      const input = new URLSearchParams();
      input.set('first', '2');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 2, after: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid pagination (with last and before)', () => {
      const input = new URLSearchParams();
      input.set('last', '2');
      input.set('before', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ last: 2, before: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid pagination with first/after as primary choice when all are present', () => {
      const input = new URLSearchParams();
      input.set('last', '2');
      input.set('before', 'fake');
      input.set('first', '2');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 2, after: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid pagination with partial first/after ignored (1)', () => {
      const input = new URLSearchParams();
      input.set('last', '2');
      input.set('before', 'fake');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ last: 2, before: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid pagination with partial first/after ignored (2)', () => {
      const input = new URLSearchParams();
      input.set('last', '2');
      input.set('before', 'fake');
      input.set('first', '2');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ last: 2, before: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid pagination with partial last/before ignored (1)', () => {
      const input = new URLSearchParams();
      input.set('last', '2');
      input.set('first', '2');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 2, after: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });

    it('should return a valid pagination with partial last/before ignored (2)', () => {
      const input = new URLSearchParams();
      input.set('before', 'fake');
      input.set('first', '2');
      input.set('after', 'fake');
      const fn = jest.fn().mockImplementation((i) => i);

      const res = getPaginationFromSearchParams({ first: 1 }, 5, input, fn);

      expect(res).toEqual({ first: 2, after: 'fake' });
      expect(fn).not.toHaveBeenCalled();
    });
  });

  describe('cleanAndSetCleanedPagination', () => {
    it('should not do anything if empty', () => {
      const input = new URLSearchParams();
      const fn = jest.fn().mockImplementation((i) => i);

      cleanAndSetCleanedPagination(input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith({});
    });

    it('should clean pagination and let other params', () => {
      const input = new URLSearchParams();
      input.set('key1', 'v1');
      const fn = jest.fn().mockImplementation((i) => i);

      cleanAndSetCleanedPagination(input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith({ key1: 'v1' });
    });
  });

  describe('cleanPaginationSearchParams', () => {
    it('should ignore all params is not present', () => {
      const input = new URLSearchParams();
      const expectedRes = new URLSearchParams();

      const res = cleanPaginationSearchParams(input);

      expect(expectedRes).toEqual(res);
    });

    it('should ignore multiple params not present', () => {
      const input = new URLSearchParams();
      const expectedRes = new URLSearchParams();

      input.set('first', '1');
      input.set('last', '1');

      const res = cleanPaginationSearchParams(input);

      expect(expectedRes).toEqual(res);
    });

    it('should be ok to remove all params', () => {
      const input = new URLSearchParams();
      const expectedRes = new URLSearchParams();

      input.set('first', '1');
      input.set('last', '1');
      input.set('before', '1');
      input.set('after', '1');

      const res = cleanPaginationSearchParams(input);

      expect(expectedRes).toEqual(res);
    });
  });
});

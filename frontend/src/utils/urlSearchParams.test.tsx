// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import { URLSearchParams } from 'url';
import {
  addJSONObjectSearchParam,
  deleteAndSetSearchParam,
  deleteAndSetSearchParams,
  getAllSearchParams,
  getJSONObjectFromSearchParam,
  setJSONObjectSearchParam,
} from './urlSearchParams';

describe('utils/urlSearchParams', () => {
  describe('deleteAndSetSearchParam', () => {
    it('should be ok to delete one not found key', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      input.set('stay', 'stay');
      expectedRes.set('stay', 'stay');

      deleteAndSetSearchParam(key, input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should be ok to delete one key', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      input.set('stay', 'stay');
      expectedRes.set('stay', 'stay');

      input.set(key, 'fake');

      deleteAndSetSearchParam(key, input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });
  });

  describe('deleteAndSetSearchParams', () => {
    it('should be ok to delete one not found key', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      input.set('stay', 'stay');
      expectedRes.set('stay', 'stay');

      deleteAndSetSearchParams([key], input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should be ok to delete two not found keys', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const key2 = 'key2';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      input.set('stay', 'stay');
      expectedRes.set('stay', 'stay');

      deleteAndSetSearchParams([key, key2], input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should be ok to delete one key', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      input.set('stay', 'stay');
      expectedRes.set('stay', 'stay');

      input.set(key, 'fake');

      deleteAndSetSearchParams([key], input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should be ok to delete two keys', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const key2 = 'key2';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      input.set('stay', 'stay');
      expectedRes.set('stay', 'stay');

      input.set(key, 'fake');
      input.set(key2, 'fake');

      deleteAndSetSearchParams([key, key2], input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should be ok to delete one key and ignore one', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const key2 = 'key2';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      input.set('stay', 'stay');
      expectedRes.set('stay', 'stay');

      input.set(key, 'fake');

      deleteAndSetSearchParams([key, key2], input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });
  });

  describe('getJSONObjectFromSearchParam', () => {
    it('should return the init value when nothing is found in search params', () => {
      const input = new URLSearchParams();
      const key = 'key1';

      const res = getJSONObjectFromSearchParam(key, {}, input);

      expect(res).toEqual({});
    });

    it('should return the object value in search params', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      input.set(key, '{"k":1}');

      const res = getJSONObjectFromSearchParam(key, {}, input);

      expect(res).toEqual({ k: 1 });
    });

    it('should return the init value when not an object in search params', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      input.set(key, 'fail');

      const res = getJSONObjectFromSearchParam(key, { k: 1 }, input);

      expect(res).toEqual({ k: 1 });
    });
  });

  describe('setJSONObjectSearchParam', () => {
    it('should be ok with empty object', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      expectedRes.set(key, '{}');

      // Call
      setJSONObjectSearchParam(key, {}, input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should be ok with simple object', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      expectedRes.set(key, '{"k1":1}');

      // Call
      setJSONObjectSearchParam(key, { k1: 1 }, input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });

    it('should be ok with complex object', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      const fn = jest.fn().mockImplementation((i) => i);
      const expectedRes = new URLSearchParams();
      expectedRes.set(key, '{"k1":{"k2":"fake"}}');

      // Call
      setJSONObjectSearchParam(key, { k1: { k2: 'fake' } }, input, fn);

      expect(fn).toHaveBeenCalledTimes(1);
      expect(fn).toHaveBeenCalledWith(expectedRes);
    });
  });

  describe('addJSONObjectSearchParam', () => {
    it('should be ok with empty object', () => {
      const input = new URLSearchParams();
      const key = 'key1';

      // Call
      const res = addJSONObjectSearchParam(key, {}, input);

      expect(res.get(key)).toEqual('{}');
    });

    it('should be ok with simple object', () => {
      const input = new URLSearchParams();
      const key = 'key1';

      // Call
      const res = addJSONObjectSearchParam(key, { k1: 1 }, input);

      expect(res.get(key)).toEqual('{"k1":1}');
    });

    it('should be ok with complex object', () => {
      const input = new URLSearchParams();
      const key = 'key1';

      // Call
      const res = addJSONObjectSearchParam(key, { k1: { k2: 'fake' } }, input);

      expect(res.get(key)).toEqual('{"k1":{"k2":"fake"}}');
    });

    it('should be ok to override key', () => {
      const input = new URLSearchParams();
      const key = 'key1';
      input.set(key, 'fake');

      // Call
      const res = addJSONObjectSearchParam(key, { k1: { k2: 'fake' } }, input);

      expect(res.get(key)).toEqual('{"k1":{"k2":"fake"}}');
    });
  });

  describe('getAllSearchParams', () => {
    it('should be ok with empty search params', () => {
      const input = new URLSearchParams();

      // Call
      const res = getAllSearchParams(input);

      expect(res).toEqual({});
    });

    it('should be ok with 1 search params', () => {
      const input = new URLSearchParams();
      input.set('fake1', 'value1');

      // Call
      const res = getAllSearchParams(input);

      expect(res).toEqual({ fake1: 'value1' });
    });

    it('should be ok with mutilple search params', () => {
      const input = new URLSearchParams();
      input.set('fake1', 'value1');
      input.set('fake2', 'value2');

      // Call
      const res = getAllSearchParams(input);

      expect(res).toEqual({ fake1: 'value1', fake2: 'value2' });
    });
  });
});

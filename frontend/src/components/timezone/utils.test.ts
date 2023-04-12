// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import { getTimeZone } from './utils';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

describe('timezone/utils', () => {
  describe('getTimeZone', () => {
    it('should be ok with Europe/London zone', () => {
      const res = getTimeZone('Europe/London');
      expect(res).toMatch(/UTC|UTC\+01:00/);
    });

    it('should be ok with Europe/Paris zone', () => {
      const res = getTimeZone('Europe/Paris');
      expect(res).toMatch(/UTC\+01:00|UTC\+02:00/);
    });

    it('should be ok with Europe/Copenhagen zone', () => {
      const res = getTimeZone('Europe/Copenhagen');
      expect(res).toMatch(/UTC-01:00|UTC/);
    });
  });
});

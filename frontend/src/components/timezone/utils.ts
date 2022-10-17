import dayjs from 'dayjs';

/* eslint-disable @typescript-eslint/no-namespace */
// Hack because not supported by Typescript....
// https://github.com/microsoft/TypeScript/issues/49231
declare namespace Intl {
  type Key = 'calendar' | 'collation' | 'currency' | 'numberingSystem' | 'timeZone' | 'unit';

  function supportedValuesOf(input: Key): string[];
}

export const availableTimezones = Intl.supportedValuesOf('timeZone');

/**
 * Get client side timezone.
 *
 * @returns {UTC(+|-)HH:mm} - Where `HH` is 2 digits hours and `mm` 2 digits minutes.
 * @example
 * // From Indian/Reunion with UTC+4
 * // 'UTC+04:00'
 * getTimeZone()
 */
export function getTimeZone(tz: string) {
  const timezoneOffset = dayjs().tz(tz).utcOffset();

  // Check if offset is zero
  if (timezoneOffset === 0) {
    return 'UTC';
  }

  const offset = Math.abs(timezoneOffset);
  const offsetOperator = timezoneOffset < 0 ? '-' : '+';
  const offsetHours = (offset / 60).toString().padStart(2, '0');
  const offsetMinutes = (offset % 60).toString().padStart(2, '0');

  return `UTC${offsetOperator}${offsetHours}:${offsetMinutes}`;
}

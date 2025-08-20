import dayjs, { type ConfigType } from 'dayjs';

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

export function getDayjsTz(input: ConfigType, timezone: string) {
  return dayjs(input).tz(timezone);
}

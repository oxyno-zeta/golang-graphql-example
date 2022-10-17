import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import type { ConfigType } from 'dayjs';

/* eslint-disable @typescript-eslint/no-explicit-any */
// eslint-disable-next-line import/prefer-default-export
export class AdapterDayjsTZ extends AdapterDayjs {
  constructor({ locale, formats, instance }: any) {
    super({ locale, formats, instance });
    // Override dayjs with tz dayjs directly
    this.dayjs = (input: ConfigType) => instance(input).tz();
  }
}

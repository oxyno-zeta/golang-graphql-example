import { createContext } from 'react';

export interface TimezoneContextModel {
  setTimezone: (v: string) => void;
  getTimezone: () => string;
}

export default createContext<TimezoneContextModel>({
  setTimezone: () => {},
  getTimezone: () => '',
});

import { createContext } from 'react';

export default createContext({
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  setTimezone: (timezone: string) => {}, // This is just an empty function, ThemeProvider will set the real one
  getTimezone: () => '' as string, // This is just an empty function, ThemeProvider will set the real one
});

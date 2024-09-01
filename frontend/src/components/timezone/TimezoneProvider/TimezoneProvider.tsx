import React, { ReactNode, useState, useMemo, useContext } from 'react';
import Cookies from 'universal-cookie';
import dayjs from 'dayjs';
import TimezoneContext from '../../../contexts/TimezoneContext';
import ConfigContext from '../../../contexts/ConfigContext';

export interface Props {
  readonly children: ReactNode;
}

const cookieName = 'selected-timezone';

function TimezoneProvider({ children }: Props) {
  // Get config from context
  const cfg = useContext(ConfigContext);
  // Get cookies object
  const cookies = useMemo(() => new Cookies(), []);
  // Get stored selected timezone
  const storedSelectedTimezone = cookies.get(cookieName);
  // Compute initial value
  let initVal = storedSelectedTimezone;
  if (!initVal) {
    initVal = dayjs.tz.guess();
  } else {
    // Save as dayjs default tz
    dayjs.tz.setDefault(initVal);
  }

  // State for timezone
  const [selectedTimezone, setSelectedTimezone] = useState<string>(initVal);

  // Expand
  const { configCookieDomain } = cfg;

  // Create color mode context
  const contextValue = useMemo(() => {
    // Set cookie
    const setCookie = (input: string) => {
      cookies.set(cookieName, input, {
        path: '/',
        maxAge: 31536000, // 1 year
        domain: configCookieDomain,
      });
    };

    return {
      getTimezone: () => selectedTimezone,
      setTimezone: (input: string) => {
        setSelectedTimezone(() => {
          // Save in storage
          setCookie(input);

          // Save as dayjs default tz
          dayjs.tz.setDefault(input);

          return input;
        });
      },
    };
  }, [selectedTimezone, configCookieDomain, cookies]);

  return <TimezoneContext.Provider value={contextValue}>{children}</TimezoneContext.Provider>;
}

export default TimezoneProvider;

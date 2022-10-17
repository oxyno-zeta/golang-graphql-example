import { useContext } from 'react';
import dayjs, { ConfigType } from 'dayjs';
import TimezoneContext from '../../../contexts/TimezoneContext';

function useDayjsTz(input: ConfigType) {
  // Get timezone context
  const timezoneCtx = useContext(TimezoneContext);

  return dayjs(input).tz(timezoneCtx.getTimezone());
}

export default useDayjsTz;

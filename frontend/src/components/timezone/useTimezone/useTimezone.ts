import { useContext } from 'react';
import TimezoneContext from '../../../contexts/TimezoneContext';

function useTimezone() {
  // Get timezone context
  const timezoneCtx = useContext(TimezoneContext);

  return timezoneCtx.getTimezone();
}

export default useTimezone;

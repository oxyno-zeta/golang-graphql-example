import { type AxiosError } from 'axios';
import { WithTraceError } from '~components/ClientProvider';

// eslint-disable-next-line import-x/prefer-default-export
export function fromAxiosErrorToWithTraceError(err: AxiosError) {
  // Get request id & trace id
  let requestId = '';
  let traceId = '';

  // Get headers
  if (err.response && err.response.headers) {
    requestId = (err.response.headers['X-Request-ID'] as string | null | undefined) || '';
    traceId = (err.response.headers['X-Trace-ID'] as string | null | undefined) || '';
  }

  return new WithTraceError(err, requestId, traceId);
}

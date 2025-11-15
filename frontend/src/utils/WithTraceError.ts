import { type ApolloLink, type ErrorLike } from '@apollo/client';
import { type AxiosError } from 'axios';

export default class WithTraceError extends Error {
  static is(error: unknown): error is WithTraceError {
    return isBranded(error, 'TraceError');
  }

  readonly name: string;

  readonly requestId: string;

  readonly traceId: string;

  readonly err: Error;

  constructor(err: Error, requestId: string, traceId: string) {
    super(err.message);
    this.requestId = requestId;
    this.traceId = traceId;
    this.err = err;
    this.name = 'TraceError';
  }
}

function isBranded(error: unknown, name: string) {
  // @ts-expect-error error.name isn't managed by typescript as check...
  return typeof error === 'object' && error !== null && Object.hasOwn(error, 'name') && error.name === name;
}

const xRequestId = 'X-Request-ID';
const xCorrelationId = 'X-Correlation-ID';
const xTraceId = 'X-Trace-ID';

export function fromAxiosErrorToWithTraceError(err: AxiosError) {
  // Get request id & trace id
  let requestId = '';
  let traceId = '';

  // Get headers
  if (err.response && err.response.headers) {
    requestId =
      ((err.response.headers[xRequestId] || err.response.headers[xCorrelationId]) as string | null | undefined) || '';
    traceId = (err.response.headers[xTraceId] as string | null | undefined) || '';
  }

  return new WithTraceError(err, requestId, traceId);
}

export function fromApolloContextErrorToWithTraceError(
  error: ErrorLike,
  context: Readonly<ApolloLink.OperationContext>,
) {
  // Get request id & trace id
  let requestId = '';
  let traceId = '';

  // Get headers
  if (context && context.response && context.response.headers) {
    requestId = context.response.headers.get(xCorrelationId) || context.response.headers.get(xRequestId);
    traceId = context.response.headers.get(xTraceId);
  }

  return new WithTraceError(error, requestId, traceId);
}

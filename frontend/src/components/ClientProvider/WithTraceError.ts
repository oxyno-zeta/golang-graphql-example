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

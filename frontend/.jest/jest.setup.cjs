// https://reactjs.org/blog/2022/03/08/react-18-upgrade-guide.html#configuring-your-testing-environment
globalThis.IS_REACT_ACT_ENVIRONMENT = true;

// This is a polyfill to have Response, festch, ... in jsdom
// See issue: https://github.com/jsdom/jsdom/issues/1724#issuecomment-720727999
require('whatwg-fetch');

// See here: https://github.com/jsdom/jsdom/issues/2524#issuecomment-897707183
// Implementation coming from here: https://github.com/remix-run/react-router/blob/94d4351290267f1a7ed93f8639a588e7a901d24f/packages/react-router/__tests__/setup.ts#L4
if (!globalThis.TextEncoder || !globalThis.TextDecoder) {
  const { TextDecoder, TextEncoder } = require('node:util');
  globalThis.TextEncoder = TextEncoder;
  globalThis.TextDecoder = TextDecoder;
}

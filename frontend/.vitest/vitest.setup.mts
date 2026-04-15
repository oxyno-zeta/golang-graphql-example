import '@testing-library/jest-dom';
import * as emotionSerializer from '@emotion/jest/serializer';
import { expect } from 'vitest';

expect.addSnapshotSerializer(emotionSerializer);

// React 19 act() environment
globalThis.IS_REACT_ACT_ENVIRONMENT = true;

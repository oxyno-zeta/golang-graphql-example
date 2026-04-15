import { readFileSync } from 'node:fs';
import path from 'node:path';

import { defineConfig, mergeConfig } from 'vitest/config';

import viteConfig from './vite.config.mts';

// Derive resolve.alias automatically from tsconfig compilerOptions.paths.
// Each entry like "~foo/*": ["./src/foo/*"] becomes "~foo" → "<root>/src/foo".
const { compilerOptions } = JSON.parse(readFileSync('./tsconfig.json', 'utf-8'));
const alias = Object.fromEntries(
  Object.entries<string[]>(compilerOptions.paths ?? {}).map(([key, [value]]) => [
    key.replace('/*', ''),
    path.resolve(__dirname, value.replace('/*', '')),
  ]),
);

export default mergeConfig(
  viteConfig,
  defineConfig({
    resolve: { alias },
    test: {
      environment: 'jsdom',
      globals: true,
      setupFiles: ['./.vitest/vitest.setup.mts'],
      reporters: ['verbose', 'junit'],
      outputFile: {
        junit: 'junit.xml',
      },
      env: {
        TZ: 'UTC',
      },
      coverage: {
        provider: 'v8',
        reportsDirectory: 'reports',
        reporter: ['text', 'clover', 'json', 'lcov', 'cobertura'],
      },
    },
  }),
);

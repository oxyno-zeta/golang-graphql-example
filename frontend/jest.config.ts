import type { Config } from '@jest/types';
import { pathsToModuleNameMapper } from 'ts-jest';

const { compilerOptions } = require('./tsconfig.json');

// Sync object
const config: Config.InitialOptions = {
  testEnvironment: 'jsdom',
  verbose: true,
  transform: {
    '^.+\\.tsx?$': 'ts-jest',
  },
  moduleNameMapper: pathsToModuleNameMapper(compilerOptions.paths),
  modulePaths: ['<rootDir>'],
  setupFiles: ['<rootDir>/.jest/jest.setup.js'],
  globalSetup: '<rootDir>/.jest/jest.global-setup.js',
  snapshotSerializers: ['@emotion/jest/serializer'],
  reporters: [
    'default',
    // 'jest-junit' to enable GitLab unit test report integration
    [
      'jest-junit',
      {
        outputName: 'junit.xml',
      },
    ],
  ],
  coverageDirectory: 'reports',
  coverageReporters: [
    // 'text' to let GitLab grab coverage from stdout
    'text',
    'clover',
    'json',
    'lcov',
    // 'cobertura' to enable GitLab test coverage visualization
    'cobertura',
  ],
};

export default config;

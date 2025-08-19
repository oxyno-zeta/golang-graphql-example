import type { Config } from '@jest/types';
import { pathsToModuleNameMapper } from 'ts-jest';
import fs from 'fs';

const tsconfigStr = fs.readFileSync('./tsconfig.json');
const { compilerOptions } = JSON.parse(tsconfigStr.toString());

// Delete assets
delete compilerOptions.paths['~assets/*'];

const tsModuleNameMapper = pathsToModuleNameMapper(compilerOptions.paths);

// Sync object
const config: Config.InitialOptions = {
  testEnvironment: 'jsdom',
  verbose: true,
  transform: {
    '^.+\\.tsx?$': 'ts-jest',
  },
  moduleNameMapper: {
    ...tsModuleNameMapper,
    '.+\\.(css|styl|less|sass|scss)': 'identity-obj-proxy',
    '.+\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)':
      '<rootDir>/__mocks__/fileMock.js',
  },
  modulePaths: ['<rootDir>'],
  setupFiles: ['<rootDir>/.jest/jest.setup.cjs'],
  globalSetup: '<rootDir>/.jest/jest.global-setup.cjs',
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

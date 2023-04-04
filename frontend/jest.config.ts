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
  setupFiles: ['<rootDir>/jest.setup.js'],
};

export default config;

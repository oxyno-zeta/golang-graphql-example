// @ts-check

import pluginJs from '@eslint/js';
import { importX as eslintPluginImportX } from 'eslint-plugin-import-x';
import jsxA11y from 'eslint-plugin-jsx-a11y';
import eslintReact from 'eslint-plugin-react';
import hooksPlugin from 'eslint-plugin-react-hooks';
import tsEslint from 'typescript-eslint';
import { defineConfig } from 'eslint/config';
import eslintConfigPrettier from 'eslint-config-prettier/flat';
import eslintPluginUnicorn from 'eslint-plugin-unicorn';
import unusedImports from 'eslint-plugin-unused-imports';
import {
  configs as airbnbConfigs,
  plugins as airbnbPlugins,
  rules as airbnbRules,
} from 'eslint-config-airbnb-extended';
// For more info, see https://github.com/storybookjs/eslint-plugin-storybook#configuration-flat-config-format
import storybook from 'eslint-plugin-storybook';

import tsParser from '@typescript-eslint/parser';

export default defineConfig(
  /**
   * ESLint recommended rules
   */
  {
    name: 'js/config',
    ...pluginJs.configs.recommended,
  },

  /**
   * Stylistic Plugin
   */
  airbnbPlugins.stylistic,

  /**
   * Eslint import X
   */
  eslintPluginImportX.flatConfigs.recommended,
  eslintPluginImportX.flatConfigs.typescript,

  /**
   * ESLint Typescript recommended
   * @see
   */
  {
    files: ['**/*.{js,mjs,cjs,jsx,mjsx,ts,tsx,mtsx}'],
    languageOptions: {
      parser: tsParser,
      ecmaVersion: 'latest',
      sourceType: 'module',
      parserOptions: {
        ecmaFeatures: { jsx: true },
        projectService: true,
        tsconfigRootDir: import.meta.dirname,
      },
    },
  },
  ...tsEslint.configs.recommended,

  /**
   * From Airbnb
   */
  ...airbnbConfigs.base.recommended,

  /**
   * React
   */
  eslintReact.configs.flat.all,
  {
    files: ['**/*.{ts,tsx}'],
    plugins: {
      'react-hooks': hooksPlugin,
    },
    rules: {
      ...hooksPlugin.configs.recommended.rules,
    },
  },

  /**
   * Static AST checker for accessibility rules on JSX elements.
   * @see https://github.com/jsx-eslint/eslint-plugin-jsx-a11y
   */
  jsxA11y.flatConfigs.recommended,

  /**
   * Airbnb react
   */
  ...airbnbConfigs.react.recommended,

  /**
   * Unicorn
   */
  {
    plugins: {
      unicorn: eslintPluginUnicorn,
    },
    // rules: {
    //   'unicorn/better-regex': 'error',
    //   'unicorn/…': 'error',
    // },
  },

  /**
   * Airbnb Typescript
   */
  ...airbnbConfigs.base.typescript,
  ...airbnbConfigs.react.typescript,
  {
    files: ['**/*.test.ts{x,}'],
    rules: {
      '@typescript-eslint/no-unsafe-return': 0,
    },
  },
  // Inspired from : airbnbRules.typescript.typescriptEslintStrict
  {
    name: 'airbnb/config/typescript/typescript-eslint/strict',
    files: ['**/*.ts', '**/*.tsx', '**/*.mts', '**/*.d.ts', '**/*.cts'],
    rules: {
      // https://typescript-eslint.io/rules/consistent-type-imports
      '@typescript-eslint/consistent-type-imports': [
        'error',
        {
          disallowTypeAnnotations: true,
          fixStyle: 'inline-type-imports',
          prefer: 'type-imports',
        },
      ],
      // https://typescript-eslint.io/rules/ban-ts-comment
      '@typescript-eslint/ban-ts-comment': [
        'error',
        {
          minimumDescriptionLength: 3,
          'ts-check': false,
          'ts-expect-error': 'allow-with-description',
          'ts-ignore': true,
          'ts-nocheck': true,
        },
      ],
      // https://typescript-eslint.io/rules/consistent-type-exports
      '@typescript-eslint/consistent-type-exports': [
        'error',
        {
          fixMixedExportsWithInlineTypeSpecifier: false,
        },
      ],
      // https://typescript-eslint.io/rules/no-duplicate-type-constituents
      '@typescript-eslint/no-duplicate-type-constituents': [
        'error',
        {
          ignoreIntersections: false,
          ignoreUnions: false,
        },
      ],
      // https://typescript-eslint.io/rules/no-explicit-any
      '@typescript-eslint/no-explicit-any': [
        'error',
        {
          fixToUnknown: false,
          ignoreRestArgs: false,
        },
      ],
    },
  },

  /**
   * Eslint Unused imports
   */
  {
    plugins: {
      'unused-imports': unusedImports,
    },
    rules: {
      'no-unused-vars': 'off',
      '@typescript-eslint/no-unused-vars': 'off',
      'unused-imports/no-unused-imports': 'error',
      'unused-imports/no-unused-vars': [
        'warn',
        {
          vars: 'all',
          varsIgnorePattern: '^_',
          args: 'after-used',
          argsIgnorePattern: '^_',
        },
      ],
    },
  },

  /**
   * Storybook
   */
  storybook.configs['flat/recommended'],

  /**
   * Custom
   */
  {
    rules: {
      'no-continue': 0,
      '@typescript-eslint/no-non-null-assertion': 0,
      'react-hooks/exhaustive-deps': 2,
      'import-x/no-named-as-default-member': 0,
      'import-x/no-rename-default': 0,
      'react/no-object-type-as-default-prop': 0,
      'no-confusing-arrow': 0,
      'react/jsx-filename-extension': [1, { extensions: ['.js', '.jsx', '.tsx', '.ts'] }],
      'react/jsx-props-no-spreading': 0,
      'react/forbid-component-props': 0,
      'import-x/no-extraneous-dependencies': [
        'error',
        { devDependencies: ['**/*.test.{j,t}s*', '**/*.spec.{j,t}s*', '**/*.stories.{j,t}s*'] },
      ],
      'import-x/no-unresolved': 'error',
      'no-console': [0],
      '@typescript-eslint/explicit-module-boundary-types': ['off'],
      '@typescript-eslint/no-use-before-define': [
        'error',
        {
          functions: false,
        },
      ],
      'no-param-reassign': [
        'error',
        {
          ignorePropertyModificationsFor: ['registration'],
        },
      ],
      'no-underscore-dangle': [
        'error',
        {
          allow: ['__WB_MANIFEST'],
        },
      ],
      'function-paren-newline': [0],
      indent: [
        2,
        2,
        {
          SwitchCase: 1,
        },
      ],
      'max-len': [
        2,
        120,
        {
          ignoreTrailingComments: true,
          ignoreStrings: true,
          ignoreUrls: true,
          ignoreTemplateLiterals: true,
          ignoreRegExpLiterals: true,
        },
      ],
      'react/jsx-no-useless-fragment': ['error', { allowExpressions: true }],
      'react/require-default-props': ['error', { functions: 'defaultArguments' }],
    },
  },
  {
    files: ['**/*.stories.tsx'],
    rules: {
      '@typescript-eslint/no-unused-vars': 0,
    },
  },

  /**
   * Prettier
   */
  eslintConfigPrettier,
);

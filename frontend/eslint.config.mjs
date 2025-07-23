// @ts-check

import pluginJs from '@eslint/js';
import eslintPluginImportX from 'eslint-plugin-import-x';
import jsxA11y from 'eslint-plugin-jsx-a11y';
import eslintReact from 'eslint-plugin-react';
import hooksPlugin from 'eslint-plugin-react-hooks';
import tsEslint from 'typescript-eslint';
import eslintPluginPrettierRecommended from 'eslint-plugin-prettier/recommended';
import eslintPluginUnicorn from 'eslint-plugin-unicorn';
import eslintConfigAirbnbReact from 'eslint-config-airbnb/rules/react';
import eslintConfigAirbnbReactA11y from 'eslint-config-airbnb/rules/react-a11y';
import eslintConfigAirbnbErrors from 'eslint-config-airbnb-base/rules/errors';
import eslintConfigAirbnbEs6 from 'eslint-config-airbnb-base/rules/es6';
import eslintConfigAirbnbStrict from 'eslint-config-airbnb-base/rules/strict';
// For more info, see https://github.com/storybookjs/eslint-plugin-storybook#configuration-flat-config-format
import storybook from 'eslint-plugin-storybook';

import tsParser from '@typescript-eslint/parser';

export default tsEslint.config(
  pluginJs.configs.recommended,
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
    settings: {
      'import-x/resolver': {
        typescript: true,
      },
    },
  },
  ...tsEslint.configs.recommended,

  /**
   * Eslint import X
   */
  eslintPluginImportX.flatConfigs.recommended,
  eslintPluginImportX.flatConfigs.typescript,

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
   * From Airbnb
   */
  { rules: eslintConfigAirbnbReact.rules },
  { rules: eslintConfigAirbnbReactA11y.rules },
  { rules: eslintConfigAirbnbErrors.rules },
  { rules: eslintConfigAirbnbEs6.rules },
  { rules: eslintConfigAirbnbStrict.rules },

  /**
   * Custom
   */
  {
    rules: {
      'react-hooks/exhaustive-deps': 2,
      'import-x/no-named-as-default-member': 0,
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
  eslintPluginPrettierRecommended,

  /**
   * Storybook
   */
  storybook.configs['flat/recommended'],
);

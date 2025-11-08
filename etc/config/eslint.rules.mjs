import typescriptEslint from '@typescript-eslint/eslint-plugin';
import typescriptEslintParser from '@typescript-eslint/parser';
import eslintConfigPrettier from 'eslint-config-prettier';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import simpleImportSort from 'eslint-plugin-simple-import-sort';
import unusedImports from 'eslint-plugin-unused-imports';
import globals from 'globals';

export const typescriptEslintRules = {
  '@typescript-eslint/await-thenable': 'error',
  '@typescript-eslint/consistent-type-definitions': 'error',
  '@typescript-eslint/dot-notation': 'error',
  '@typescript-eslint/explicit-function-return-type': [
    'error',
    {
      allowConciseArrowFunctionExpressionsStartingWithVoid: true,
      allowDirectConstAssertionInArrowFunctions: true,
      allowExpressions: false,
      allowHigherOrderFunctions: true,
      allowTypedFunctionExpressions: true,
    },
  ],
  '@typescript-eslint/explicit-member-accessibility': [
    'error',
    {
      accessibility: 'explicit',
    },
  ],
  '@typescript-eslint/explicit-module-boundary-types': [
    'error',
    {
      allowArgumentsExplicitlyTypedAsAny: true,
      allowDirectConstAssertionInArrowFunctions: true,
      allowHigherOrderFunctions: false,
      allowTypedFunctionExpressions: false,
    },
  ],
  '@typescript-eslint/member-ordering': 'off',
  '@typescript-eslint/no-empty-function': 'off',
  '@typescript-eslint/no-empty-interface': 'off',
  '@typescript-eslint/no-empty-object-type': 'error',
  '@typescript-eslint/no-explicit-any': 'error',
  '@typescript-eslint/no-extraneous-class': 'error',
  '@typescript-eslint/no-floating-promises': 'error',
  '@typescript-eslint/no-inferrable-types': [
    'error',
    {
      ignoreParameters: true,
    },
  ],
  '@typescript-eslint/no-misused-new': 'error',
  '@typescript-eslint/no-namespace': 'off',
  '@typescript-eslint/no-non-null-assertion': 'off',
  '@typescript-eslint/no-require-imports': 'error',
  '@typescript-eslint/no-restricted-types': 'error',
  '@typescript-eslint/no-shadow': [
    'error',
    {
      hoist: 'all',
    },
  ],
  '@typescript-eslint/no-this-alias': 'error',
  '@typescript-eslint/no-unsafe-argument': 'off',
  '@typescript-eslint/no-unsafe-assignment': 'off',
  '@typescript-eslint/no-unsafe-call': 'off',
  '@typescript-eslint/no-unsafe-declaration-merging': 'error',
  '@typescript-eslint/no-unsafe-function-type': 'error',
  '@typescript-eslint/no-unsafe-member-access': 'off',
  '@typescript-eslint/no-unused-expressions': 'error',
  '@typescript-eslint/no-unsafe-return': 'off',
  '@typescript-eslint/no-unused-vars': 'off',
  '@typescript-eslint/no-wrapper-object-types': 'error',
  '@typescript-eslint/prefer-for-of': 'error',
  '@typescript-eslint/prefer-function-type': 'error',
  '@typescript-eslint/prefer-namespace-keyword': 'error',
  '@typescript-eslint/promise-function-async': 'off',
  '@typescript-eslint/restrict-plus-operands': 'error',
  '@typescript-eslint/restrict-template-expressions': 'off',
  '@typescript-eslint/strict-boolean-expressions': [
    'error',
    {
      allowString: false,
      allowNumber: false,
      allowNullableObject: false,
      allowNullableBoolean: false,
      allowNullableString: false,
      allowNullableNumber: false,
      allowNullableEnum: false,
      allowAny: false,
      allowRuleToRunWithoutStrictNullChecksIKnowWhatIAmDoing: false,
    },
  ],
  '@typescript-eslint/triple-slash-reference': [
    'error',
    {
      lib: 'always',
      path: 'always',
      types: 'prefer-import',
    },
  ],
  '@typescript-eslint/typedef': [
    'error',
    {
      parameter: true,
      propertyDeclaration: true,
    },
  ],
  '@typescript-eslint/unbound-method': 'off',
  '@typescript-eslint/unified-signatures': 'error',
};

export const simpleImportSortRules = {
  'simple-import-sort/imports': 'error',
  'simple-import-sort/exports': 'error',
};

export const unusedImportsRules = {
  'unused-imports/no-unused-imports': 'error',
  'unused-imports/no-unused-vars': [
    'error',
    {
      args: 'all',
      argsIgnorePattern: '^_',
      caughtErrors: 'all',
      caughtErrorsIgnorePattern: '^_',
      varsIgnorePattern: '^_',
    },
  ],
};

export const esLintRules = {
  complexity: [
    'error',
    {
      max: 30,
    },
  ],
  'constructor-super': 'error',
  eqeqeq: ['error', 'smart'],
  'guard-for-in': 'error',
  'id-denylist': ['error', 'any', 'Number', 'number', 'String', 'string', 'Boolean', 'boolean', 'Undefined', 'undefined'],
  'id-match': 'error',
  'no-bitwise': 'error',
  'no-caller': 'error',
  'no-case-declarations': 'error',
  'no-cond-assign': 'error',
  'no-console': [
    'warn',
    {
      allow: [
        'warn',
        'dir',
        'timeLog',
        'assert',
        'clear',
        'count',
        'countReset',
        'group',
        'groupEnd',
        'table',
        'dirxml',
        'error',
        'groupCollapsed',
        'Console',
        'profile',
        'profileEnd',
        'timeStamp',
        'context',
      ],
    },
  ],
  'no-debugger': 'error',
  'no-duplicate-case': 'error',
  'no-duplicate-imports': 'error',
  'no-empty': [
    'error',
    {
      allowEmptyCatch: true,
    },
  ],
  'no-eval': 'error',
  'no-fallthrough': 'error',
  'no-implicit-coercion': 'error',
  'no-inner-declarations': 'off',
  'no-invalid-this': 'off',
  'no-irregular-whitespace': 'error',
  'no-multiple-empty-lines': [
    'error',
    {
      max: 1,
    },
  ],
  'no-new-wrappers': 'error',
  'no-redeclare': 'error',
  'no-restricted-imports': [
    'error',
    {
      patterns: ['**/index.js', '**/index.mjs'],
    },
  ],
  'no-return-await': 'error',
  'no-shadow': 'off',
  'no-throw-literal': 'error',
  'no-undef-init': 'error',
  'no-unused-expressions': 'off',
  'no-unused-labels': 'error',
  'no-unused-vars': 'off',
  'no-var': 'error',
  'one-var': ['error', 'never'],
  'padding-line-between-statements': [
    'off',
    {
      blankLine: 'always',
      next: 'return',
      prev: '*',
    },
  ],
  'prefer-const': 'error',
  'prefer-object-spread': 'error',
  'prefer-template': 'error',
  radix: 'error',
  'sort-imports': 'off', // off for simple-import-sort
  'use-isnan': 'error',
  yoda: 'error',
};

const rules = [
  // all typescript files
  {
    files: ['**/*.{ts,cts,mts}'],
    ignores: ['**/dist/**/*', '**/node_modules/**/*', '**/output/**/*'],
    languageOptions: {
      parser: typescriptEslintParser,
      parserOptions: {
        // https://typescript-eslint.io/packages/parser/#projectservice
        projectService: {
          allowDefaultProject: ['vitest.config.ts'],
        },
      },
    },
    plugins: {
      '@typescript-eslint': typescriptEslint,
      'unused-imports': unusedImports,
      'simple-import-sort': simpleImportSort,
    },
    rules: {
      ...esLintRules,
      ...simpleImportSortRules,
      ...typescriptEslintRules,
      ...unusedImportsRules,
    },
  },
  // all javascript files
  {
    files: ['**/*.{js,cjs,mjs}'],
    ignores: ['**/dist/**/*', '**/node_modules/**/*', '**/output/**/*'],
    plugins: {
      'unused-imports': unusedImports,
      'simple-import-sort': simpleImportSort,
    },
    rules: {
      ...esLintRules,
      ...simpleImportSortRules,
      ...unusedImportsRules,
    },
  },
  // all tsx files
  {
    files: ['**/*.tsx'],
    ignores: ['**/dist/**/*', '**/node_modules/**/*', '**/output/**/*'],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
      parser: typescriptEslintParser,
      parserOptions: {
        projectService: true,
      },
    },
    name: 'renderer',
    plugins: {
      '@typescript-eslint': typescriptEslint,
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
      'simple-import-sort': simpleImportSort,
      'unused-imports': unusedImports,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      'react-refresh/only-export-components': ['warn', { allowConstantExport: true }],
      //...esLintRules,
      ...simpleImportSortRules,
      ...typescriptEslintRules,
      ...unusedImportsRules,
    },
  },
  eslintConfigPrettier,
];

export { rules };

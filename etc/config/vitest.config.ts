import path from 'node:path';

import { coverageConfigDefaults, defineConfig } from 'vitest/config';

//
// this file is imported relative to each project folder
//

const isTestDebug = process.env['TEST_DEBUG'] === '1';
const packageName = process.env['npm_package_name'] ?? '';
const rootDir = process.env['ROOT_DIR'] ?? '';

const config = defineConfig({
  test: {
    bail: 1,
    coverage: {
      exclude: ['**/output/**', ...coverageConfigDefaults.exclude],
      reportsDirectory: path.join(rootDir, 'output', packageName, 'coverage'),
    },
    environment: 'node',
    globals: true,
    logHeapUsage: true,
    name: packageName,
    passWithNoTests: true,
    reporters: ['default', 'junit', 'html'],
    outputFile: {
      junit: path.join(rootDir, 'output', packageName, 'tests', 'junit.xml'),
      html: path.join(rootDir, 'output', packageName, 'tests', 'html', 'index.html'),
    },
  },
});

if (isTestDebug) {
  // to be able to debug tests in VSCode with breakpoints
  config.test!.testTimeout = 60 * 60 * 1000;
}

export default config;

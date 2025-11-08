import { defineConfig, externalizeDepsPlugin } from 'electron-vite';
import { Plugin } from 'rollup';
import { createFilter, TransformResult } from 'vite';

export default defineConfig({
  main: {
    build: {
      outDir: 'output/main',
      rollupOptions: {
        input: 'apps/main/src/main.ts',
        output: {
          entryFileNames: 'main.cjs',
          format: 'cjs',
        },
      },
      sourcemap: true,
    },
    environments: {},
    plugins: [externalizeDepsPlugin()],
  },
  preload: {
    build: {
      outDir: 'output/preload',
      rollupOptions: {
        input: 'apps/preload/src/preload.ts',
        output: {
          entryFileNames: 'preload.cjs',
          format: 'cjs',
        },
      },
      sourcemap: true,
    },
    plugins: [externalizeDepsPlugin()],
  },
  renderer: {
    root: 'apps/renderer',
    build: {
      outDir: 'output/renderer',
      rollupOptions: {
        input: 'apps/renderer/index.html',
      },
      sourcemap: true,
    },
    plugins: [removeUseClientPlugin()],
  },
});

// https://stackoverflow.com/a/78751258
function removeUseClientPlugin(): Plugin {
  const filter = createFilter(/.*\.(js|ts|cjs|cts|mjs|mts|jsx|tsx)$/);
  return {
    name: 'remove-use-client',
    transform(code: string, id: string): TransformResult | undefined {
      if (!filter(id)) {
        return undefined;
      }
      const newCode = code.replace(/['"]use client['"];\s*/g, '');
      return { code: newCode, map: null };
    },
  };
}

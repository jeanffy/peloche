import fs from 'node:fs';
import path from 'node:path';

import react from '@vitejs/plugin-react';
import { defineConfig, Plugin, ResolvedConfig, SassPreprocessorOptions } from 'vite';

// https://vite.dev/config
export default defineConfig({
  base: './',
  build: {
    outDir: 'dist',
  },
  css: {
    preprocessorOptions: {
      scss: {
        api: 'modern-compiler',
      } as SassPreprocessorOptions,
    },
  },
  // publicDir disabled here because vite does copy all public dir files, without any exclusion possible
  // exclusions are needed because of symbolic links that are needed for running with electron-vite or debugging (see filePathToRendererUrl)
  // we replace vite publicDir copy with a custom plugin
  publicDir: false,
  plugins: [react(), publicDirCopyPlugin()],
});

function publicDirCopyPlugin(): Plugin {
  let resolvedConfig: ResolvedConfig;
  return {
    name: 'public-dir-copy',
    apply: 'build',
    configResolved(config: ResolvedConfig): void {
      resolvedConfig = config;
    },
    writeBundle(): void {
      const foldersToCopy = ['locales'];
      for (const folderToCopy of foldersToCopy) {
        fs.cpSync(path.join(resolvedConfig.root, 'public', folderToCopy), path.join(resolvedConfig.root, resolvedConfig.build.outDir, folderToCopy), {
          recursive: true,
        });
      }
    },
  };
}

import { defineConfig } from '@farmfe/core';
import tsconfigPaths from 'vite-tsconfig-paths';
import react from '@vitejs/plugin-react';
import eslint from 'vite-plugin-eslint';
import { visualizer } from 'rollup-plugin-visualizer';
import preserveDirectives from 'rollup-preserve-directives';

export default defineConfig({
  compilation: {
    persistentCache: true,
    lazyCompilation: true,
    partialBundling: {
      targetConcurrentRequests: 15,
      targetMinSize: 200 * 1024, // 200 KB
      targetMaxSize: 500 * 1024,
    },
    sourcemap: 'all',
    input: {
      index: './index.html',
    },
    output: {
      path: 'build',
      publicPath: '/',
      targetEnv: 'browser-esnext',
      assetsFilename: '[entryName]-[resourceName].[hash].[ext]',
    },
    script: {
      plugins: [],
      target: 'esnext',
    },
  },
  server: {
    port: 3000,
    hmr: true,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  plugins: ['@farmfe/plugin-react'],
  vitePlugins: [
    tsconfigPaths(),
    eslint({
      emitWarning: true,
      // See issue: https://github.com/storybookjs/builder-vite/issues/367
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      // @ts-ignore
      exclude: [/virtual:/, /node_modules/],
    }),
    visualizer({
      template: 'treemap', // or sunburst
      open: false,
      gzipSize: true,
      brotliSize: true,
      filename: 'bundle-analyze.html', // will be saved in project's root
    }),
  ],
});

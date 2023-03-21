import { defineConfig } from 'vite';
import tsconfigPaths from 'vite-tsconfig-paths';
import react from '@vitejs/plugin-react';
import eslint from 'vite-plugin-eslint';

export default defineConfig({
  plugins: [
    react({
      // Exclude storybook stories
      exclude: /\.stories\.(t|j)sx?$/,
    }),
    tsconfigPaths(),
    eslint({ emitWarning: true }),
  ],
  build: {
    target: 'es2018',
    rollupOptions: {},
    sourcemap: true,
  },
  cacheDir: './.vite-cache',
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
});

import { defineConfig } from 'vite';
import tsconfigPaths from 'vite-tsconfig-paths';
import react from '@vitejs/plugin-react';
import eslint from 'vite-plugin-eslint';
import { visualizer } from 'rollup-plugin-visualizer';
import preserveDirectives from 'rollup-preserve-directives';
import UnpluginInjectPreload from 'unplugin-inject-preload/vite';
import { imagetools } from 'vite-imagetools';

export default defineConfig({
  plugins: [
    preserveDirectives(),
    react({
      // Exclude storybook stories
      exclude: /\.stories\.(t|j)sx?$/,
    }),
    tsconfigPaths(),
    // eslint({
    //   emitWarning: true,
    //   // See issue: https://github.com/storybookjs/builder-vite/issues/367
    //   // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    //   // @ts-ignore
    //   exclude: [/virtual:/, /node_modules/, /sb-preview/],
    // }),
    visualizer({
      template: 'treemap', // or sunburst
      open: false,
      gzipSize: true,
      brotliSize: true,
      filename: 'bundle-analyze.html', // will be saved in project's root
    }),
    UnpluginInjectPreload({
      files: [
        {
          outputMatch: /.*.(png|jpg|webp|avif)$/,
          attributes: {
            rel: 'preload',
            as: 'image',
          },
        },
      ],
      injectTo: 'head-prepend',
    }),
    imagetools({
      cache: {
        enabled: true,
        dir: './node_modules/.cache/imagetools',
        retention: 172800,
      },
    }),
  ],
  build: {
    target: 'es2018',
    rollupOptions: {
      output: {
        manualChunks: {
          muibase: ['@mui/material', '@emotion/react', '@emotion/styled'],
          muiheavy: ['@mui/lab', '@mui/x-data-grid', '@mui/x-date-pickers'],
          connectivity: ['axios', '@apollo/client', 'graphql'],
          translate: ['i18next', 'i18next-browser-languagedetector', 'i18next-http-backend', 'react-i18next'],
        },
      },
    },
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

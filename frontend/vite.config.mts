import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
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
          entryMatch: /.*\.woff2$/,
        },
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
    rolldownOptions: {
      output: {
        strictExecutionOrder: true,
        codeSplitting: {
          minSize: 400000, // 400KB
          maxSize: 500000, // 500KB
          groups: [
            {
              name: 'vendor',
              test: /node_modules/,
            },
          ],
        },
      },
    },
    sourcemap: true,
  },
  resolve: {
    tsconfigPaths: true,
    conditions: ['mui-modern', 'module', 'browser', 'development|production'],
  },
  assetsInclude: ['**/*.png', '**/*.jpg'],
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

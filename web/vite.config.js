import { fileURLToPath, URL } from 'node:url';

import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

// Define your custom build output directory
const customBuildPath = './static'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  // Set the custom build output directory
  build: {
    outDir: customBuildPath,
  },
  server: {
    proxy: {
      '/api': { // adjust this to target paths to be rerouted
        target: 'http://localhost:8082', // your Docker server address
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, '')
      },
    },
  }
});
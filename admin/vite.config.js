import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

customBuildPath = './static'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  base: "/admin",
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
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
})

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// target: 'http://10.0.33.50:23456',

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      '/api': {
        target: 'https://console.ismdeep.com',
        changeOrigin: true,
      }
    }
  }
})

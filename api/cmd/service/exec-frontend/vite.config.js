import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    port: 8080,
    proxy: {
      '/server': {
        target: 'http://localhost:8003',
        ws: true,
        secure: false,
        changeOrigin: true,
        rewrite: (p) => p.replace(/^\/server/, '')
      },
      '/creatorapi': {
        target: 'http://localhost:8005',
        ws: true,
        secure: false,
        changeOrigin: true,
        rewrite: (p) => p.replace(/^\/creatorapi/, '')
      },
    }
  },
  productionSourceMap: false
})

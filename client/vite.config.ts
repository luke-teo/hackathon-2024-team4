import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    cors: false,
    port: 5173,
    proxy: {
      "/api": "https://e5b9-152-165-116-45.ngrok-free.app",
    }
  },
})

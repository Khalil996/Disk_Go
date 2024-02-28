import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path"
import {fileURLToPath, URL} from 'node:url'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: {
        port: 5173,
        proxy: {
            '/api': {
                target: 'http://localhost:8888/',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '') // 不可以省略rewrite
            }
        }
    },
    // 这里是配置@为src目录，
    resolve: {
        alias: {
            // '@': path.resolve('./src'),
            '@': fileURLToPath(new URL('./src', import.meta.url))
        },
    },
})


import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {fileURLToPath, URL} from 'node:url'

export default defineConfig({
    plugins: [vue()],
    server: {
        open: false,
        port: 3000,
        proxy: {
            '/api': {
                target: 'http://localhost:8888/', //
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '')
            },
        },
    },
    // 这里是配置@为src目录，
    resolve: {
        alias: {
            // '@': path.resolve('./src'),
            '@': fileURLToPath(new URL('./src', import.meta.url))
        },
    },
})

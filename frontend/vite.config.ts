/// <reference types="vitest" />
import react from '@vitejs/plugin-react'
import { defineConfig, loadEnv } from 'vite'
import dotenv from 'dotenv'
import tsconfigPaths from 'vite-tsconfig-paths'
import { resolve } from 'path'
import dynamicImport from 'vite-plugin-dynamic-import'
import Config from './config.json'
import { nodePolyfills } from 'vite-plugin-node-polyfills'

dotenv.config()

export default ({ mode }) => {
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) }

  // import.meta.env.VITE_PORT available here with: process.env.VITE_PORT

  return defineConfig({
    base: '/',
    plugins: [
      react({
        jsxImportSource: '@emotion/react',
        jsxRuntime: 'automatic',
        babel: {
          plugins: ['@emotion/babel-plugin'],
        },
      }),
      tsconfigPaths({ root: '.' }),
      dynamicImport({}),
      // nodePolyfills(),
    ],
    server: {
      host: Config.DOMAIN,
      port: Number(Config.FRONTEND_PORT) || 5143,
      strictPort: true,
      // hmr: {
      //   protocol: 'wss',
      //   clientPort: 9443,
      // },
      https: {
        key: '../certificates/localhost-key.pem',
        cert: '../certificates/localhost.pem',
      },
    },
    optimizeDeps: {
      exclude: ['react-hook-form'],
    },
    define: {
      'process.env.NODE_ENV': `"${mode}"`,
    },
    esbuild: {
      logOverride: { 'this-is-undefined-in-esm': 'silent' },
    },
    build: {
      minify: 'terser',
      commonjsOptions: {
        transformMixedEsModules: true,
      },
      terserOptions: {
        compress: {
          drop_console: true,
          drop_debugger: true,
        },
      },
      outDir: './build',
      rollupOptions: {
        input: {
          main: resolve(__dirname, 'index.html'),
          // nested: resolve(__dirname, 'nested/index.html')
        },
        external: ['src/index.tsx'],
      },
      dynamicImportVarsOptions: {
        exclude: [],
      },
    },
  })
}

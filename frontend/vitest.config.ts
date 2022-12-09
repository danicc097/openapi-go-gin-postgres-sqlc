import path from 'path'
import react from '@vitejs/plugin-react'
import dotenv from 'dotenv'
import tsconfigPaths from 'vite-tsconfig-paths'
import { resolve } from 'path'
import dynamicImport from 'vite-plugin-dynamic-import'
import Config from './config.json'
import { defineConfig } from 'vitest/config'

const r = (p: string) => resolve(__dirname, p)

const alias: Record<string, string> = {
  '~~': r('.'),
  '~~/': r('./'),
  '@@': r('.'),
  '@@/': r('./'),
  assets: r('./assets'),
  public: r('./public'),
  'public/': r('./public/'),
  '@': path.resolve(__dirname, './src'),
  '@roles': path.resolve(__dirname, './roles.json'),
  '@scopes': path.resolve(__dirname, './scopes.json'),
  '@config': path.resolve(__dirname, './config.json'),
  '@operationAuth': path.resolve(__dirname, './operationAuth.gen.json'),
}

export default defineConfig({
  // esbuild: {
  //   tsconfigRaw: {},
  // },
  // resolve: {
  //   alias,
  // },
  test: {
    deps: {
      inline: ['framer-motion'],
    },
    globals: true,
    environmentOptions: {
      jsdom: {
        console: true,
      },
    },
    environment: 'jsdom',
    setupFiles: './src/setupTests.ts',
    coverage: {
      reporter: ['text', 'html'],
      exclude: ['node_modules/', 'src/setupTests.ts'],
    },
    // transformMode: {
    //   web: [/\.[jt]sx$/],
    // },
  },
})

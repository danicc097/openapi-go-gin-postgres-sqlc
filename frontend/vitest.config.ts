import path from 'path'
import react from '@vitejs/plugin-react'
import dotenv from 'dotenv'
import tsconfigPaths from 'vite-tsconfig-paths'
import { resolve } from 'path'
import dynamicImport from 'vite-plugin-dynamic-import'
import Config from './config.json'
import { defineConfig } from 'vitest/config'
import viteConfig from './vite.config'
import { mergeConfig } from 'vite'

const r = (p: string) => resolve(__dirname, p)

export default mergeConfig(
  viteConfig,
  defineConfig({
    // esbuild: {
    //   tsconfigRaw: {},
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
        provider: 'c8',
        reporter: ['text', 'html'],
        exclude: ['node_modules/', 'src/setupTests.ts'],
      },
      // `vitest typecheck`, not run in watch (https://github.com/vitest-dev/vitest/issues/2299)
      typecheck: {
        ignoreSourceErrors: true,
      },
      css: false,
      // transformMode: {
      //   web: [/\.[jt]sx$/],
      // },
    },
  }),
)

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

const alias: Record<string, string> = {
  '~': r('src'),
  src: r('./src'),
  '~~': r('.'),
  '~~/': r('./'),
  '@@': r('.'),
  '@@/': r('./'),
  assets: r('./assets'),
  public: r('./public'),
  'public/': r('./public/'),
  '@': r('./src'),
  '@operationAuth': r('./operationAuth.gen.json'),
}

export default mergeConfig(
  viteConfig,
  defineConfig({
    // esbuild: {
    //   tsconfigRaw: {},
    // },
    resolve: {
      alias,
    },
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
      css: false,
      // transformMode: {
      //   web: [/\.[jt]sx$/],
      // },
    },
  }),
)

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
import { nodePolyfills } from 'vite-plugin-node-polyfills'

const r = (p: string) => resolve(__dirname, p)

// needed for absolute path resolution
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
}

export default defineConfig((env) =>
  mergeConfig(
    viteConfig(env),
    defineConfig({
      // esbuild: {
      //   tsconfigRaw: {},
      // },

      resolve: {
        alias,
      },
      define: {
        'process.env.TESTING': true,
      },
      test: {
        env: {
          TESTING: 'true',
        },
        testTimeout: 50_000,
        // reporters: ['verbose'],
        // outputFile: './reporter-output/result',
        isolate: false, // if we dont depend on global state can be false to make tests faster
        deps: {
          optimizer: {
            web: {
              include: [
                'framer-motion',
                '@tabler/icons',
                '@tabler/icons-react',
                // breaks all imports
                //'@mantine/core',
              ],
              // exclude: ['@mantine/core'],
              enabled: true,
              esbuildOptions: {
                treeShaking: false,
              },
            },
          },
        },
        globals: true,
        environmentOptions: {
          jsdom: {
            console: true,
          },
        },
        environment: 'jsdom',
        setupFiles: './src/setupTests.tsx',
        coverage: {
          provider: 'v8',
          reporter: ['text', 'html'],
          exclude: ['node_modules/', 'src/setupTests.tsx'],
        },
        // `vitest typecheck`, not run in watch (https://github.com/vitest-dev/vitest/issues/2299)
        typecheck: {
          exclude: ['**/node_modules/**/*'],
          include: ['src/**/*.test-d.ts*'],
          ignoreSourceErrors: true,
          tsconfig: `${__dirname}/tsconfig.json`,
        },
        css: false,
        // transformMode: {
        //   web: [/\.[jt]sx$/],
        // },
      },
    }),
  ),
)

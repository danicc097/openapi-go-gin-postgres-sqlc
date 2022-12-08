import { defineConfig } from 'vitest/config'
import path from 'path'

export default defineConfig({
  resolve: {
    alias: {
      // '@': path.resolve(__dirname, './src'),
    },
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
      reporter: ['text', 'html'],
      exclude: ['node_modules/', 'src/setupTests.ts'],
    },
    // transformMode: {
    //   web: [/\.[jt]sx$/],
    // },
  },
})

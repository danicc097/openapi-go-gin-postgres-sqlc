import { defineConfig } from 'orval'

// for custom client see https://github.com/anymaniax/orval/blob/master/samples/react-query/custom-client/src/api/mutator/custom-client.ts#L1
export default defineConfig({
  main: {
    output: {
      mode: 'split',
      target: './src/gen/main.ts',
      schemas: './src/gen/model',
      client: 'react-query',
      mock: true,
      tsconfig: './tsconfig.json',
      override: { useDates: true },
    },
    input: {
      target: '../openapi.yaml',
    },
    hooks: {
      afterAllFilesWrite: 'prettier --write',
    },
  },
})

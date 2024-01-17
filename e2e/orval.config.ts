import { defineConfig } from 'orval'

// for custom client see https://github.com/anymaniax/orval/blob/master/samples/react-query/custom-client/src/api/mutator/custom-client.ts#L1
export default defineConfig({
  main: {
    output: {
      mode: 'tags-split',
      target: './client/gen/main.ts',
      schemas: './client/gen/model',
      client: 'axios-functions',
      mock: false,
      tsconfig: './tsconfig.json',
      // for extreme cases can also override the core package itself https://github.com/anymaniax/orval/tree/master/packages/core
      override: {
        // TODO Axios converter if using useDates
        // https://orval.dev/reference/configuration/output#usedates
        useDates: true,
        mutator: { path: 'client/api/mutator.ts', name: 'customInstance' },
      },
    },
    input: {
      target: '../openapi.exploded.yaml',
      // validation: true, // https://github.com/IBM/openapi-validator/#configuration via .validaterc
    },
    hooks: {
      afterAllFilesWrite: 'prettier --write',
    },
  },
})

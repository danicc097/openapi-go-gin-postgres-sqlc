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
      // for extreme cases can also override the core package itself https://github.com/anymaniax/orval/tree/master/packages/core
      override: {
        useDates: true,
        query: {
          signal: true, // generation of abort signal
          useQuery: true,
          useInfinite: true, // https://tanstack.com/query/v4/docs/guides/infinite-queries
          useInfiniteQueryParam: 'nextId',
          options: {
            // QueryObserverOptions passed directly to all generated queries
            staleTime: 1000 * 3600,
          },
        },
      },
    },
    input: {
      target: '../openapi.yaml',
      // validation: true, // https://github.com/IBM/openapi-validator/#configuration via .validaterc
    },
    hooks: {
      afterAllFilesWrite: 'prettier --write',
    },
  },
})

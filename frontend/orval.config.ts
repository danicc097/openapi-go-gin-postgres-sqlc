import { defineConfig } from 'orval'
import { faker } from '@faker-js/faker'
import _, { uniqueId } from 'lodash'
import { reactQueryDefaultAppOptions } from './src/react-query'

// for custom client see https://github.com/anymaniax/orval/blob/master/samples/react-query/custom-client/src/api/mutator/custom-client.ts#L1
export default defineConfig({
  main: {
    output: {
      mode: 'tags-split',
      target: './src/gen/main.ts',
      schemas: './src/gen/model',
      client: 'react-query',
      mock: true, // until changes are made to deeply nested generation
      tsconfig: './tsconfig.json',
      // for extreme cases can also override the core package itself https://github.com/anymaniax/orval/tree/master/packages/core
      override: {
        // TODO Axios converter if using useDates
        // https://orval.dev/reference/configuration/output#usedates
        useDates: true,
        query: {
          signal: true, // generation of abort signal
          useQuery: true,
          useInfinite: true, // https://tanstack.com/query/v4/docs/guides/infinite-queries
          // useInfiniteQueryParam: 'nextId',
          options: reactQueryDefaultAppOptions.queries,
        },
        mock: {
          format: {
            date: () => faker.date.past(),
            'date-time': () => faker.date.past(),
          },
          properties: {
            // will use basic string replace to get BrandedTypes.
            // userID: () => faker.datatype.uuid(),
            email: () => faker.internet.email(),
            metadata: () => ({
              key: faker.color.hsl(),
            }),
          },
          required: true,
        },
        mutator: { path: 'src/api/mutator.ts', name: 'customInstance' },
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

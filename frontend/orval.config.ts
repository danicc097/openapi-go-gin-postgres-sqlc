import { defineConfig } from 'orval'
import { faker } from '@faker-js/faker'
import * as _ from 'lodash'

// relative paths only
import { reactQueryDefaultAppOptions } from './src/react-query.default'

// for custom client see https://github.com/anymaniax/orval/blob/master/samples/react-query/custom-client/src/api/mutator/custom-client.ts#L1
export default defineConfig({
  main: {
    output: {
      mock: true,
      mode: 'tags-split',
      target: './src/gen/main.ts',
      schemas: './src/gen/model',
      client: 'react-query',
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
          options: reactQueryDefaultAppOptions.queries,
          // FIXME: leads to issues with /events, /oidc and /project where it assumes there's a cursor param
          useInfiniteQueryParam: 'cursor', // same param for all app paginated queries.
        },
        operations: {
          ..._.fromPairs(
            ['GetProjectWorkitems', 'MyProviderLogin', 'Events'].map((operation) => [
              operation,
              {
                query: {
                  useQuery: true,
                  useInfinite: false,
                },
              },
            ]),
          ),
        },
        mock: {
          delay: 200,
          format: {
            date: () => faker.date.past(),
            'date-time': () => faker.date.past(),
          },
          properties: {
            // will use basic string replace to get BrandedTypes.
            // userID: () => faker.string.uuid(),
            email: () => faker.internet.email(),
            metadata: () => ({
              key: faker.string.sample(),
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
    // required for orval types gen right after
    // hooks: {
    //   afterAllFilesWrite: 'prettier --write',
    // },
  },
})

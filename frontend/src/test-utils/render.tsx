import { render as testingLibraryRender } from '@testing-library/react'
import { MantineProvider } from '@mantine/core'
import { queryClient } from 'src/react-query'
import { persister } from 'src/idb'
import { PersistQueryClientProvider } from '@tanstack/react-query-persist-client'
import { QueryClientProvider } from '@tanstack/react-query'
import { VirtuosoMockContext } from 'react-virtuoso'

export function render(ui: React.ReactNode) {
  return testingLibraryRender(<>{ui}</>, {
    wrapper: ({ children }: { children: React.ReactNode }) => (
      <VirtuosoMockContext.Provider value={{ viewportHeight: Infinity, itemHeight: 100 }}>
        <QueryClientProvider client={queryClient} /**persistOptions={{ persister }} */>
          <MantineProvider>{children}</MantineProvider>
        </QueryClientProvider>
      </VirtuosoMockContext.Provider>
    ),
  })
}

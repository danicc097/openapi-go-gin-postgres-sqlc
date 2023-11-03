import { render as testingLibraryRender } from '@testing-library/react'
import { MantineProvider } from '@mantine/core'
import { queryClient } from 'src/react-query'
import { persister } from 'src/idb'
import { PersistQueryClientProvider } from '@tanstack/react-query-persist-client'

export function render(ui: React.ReactNode) {
  return testingLibraryRender(<>{ui}</>, {
    wrapper: ({ children }: { children: React.ReactNode }) => (
      <PersistQueryClientProvider client={queryClient} persistOptions={{ persister }}>
        <MantineProvider>{children}</MantineProvider>
      </PersistQueryClientProvider>
    ),
  })
}

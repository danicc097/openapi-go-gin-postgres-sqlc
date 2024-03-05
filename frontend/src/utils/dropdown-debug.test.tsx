import { act, fireEvent, render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { Combobox, Group, InputBase, MantineProvider, Select, Text } from '@mantine/core'
import { vitest } from 'vitest'
import { SelectOptionComponentDebug } from 'src/utils/dropdown-debug'
import { VirtuosoMockContext } from 'react-virtuoso'

describe.skip('Select component', () => {
  it('calls onChange handler when an option is selected', async () => {
    const handleChange = vitest.fn()
    const options = [
      { value: '1', label: 'Option 1' },
      { value: '2', label: 'Option 2' },
      { value: '3', label: 'Option 3' },
    ]

    render(
      <MantineProvider>
        <Select data-testid="select" data={options} onChange={handleChange} />
      </MantineProvider>,
    )

    await waitFor(async () => {
      const dropdownMenu = screen.getByTestId('select')

      await userEvent.click(screen.getByText('Option 2'))
    })

    expect(handleChange).toHaveBeenCalledTimes(1)
    expect(handleChange).toHaveBeenCalledWith('2', { value: '2', label: 'Option 2' })
  })
})

describe.skip('Combobox component', () => {
  it('calls onChange handler when an option is selected', async () => {
    const handleChange = vitest.fn()

    render(
      <VirtuosoMockContext.Provider value={{ viewportHeight: Infinity, itemHeight: 100 }}>
        <MantineProvider>
          <SelectOptionComponentDebug onSubmit={handleChange} defaultOption="Bananas" />
        </MantineProvider>
      </VirtuosoMockContext.Provider>,
    )

    const getOptionBananas = () => screen.getByRole('option', { name: 'Bananas', hidden: true }) as HTMLOptionElement
    const getOptionBroccoli = () => screen.getByRole('option', { name: 'Broccoli', hidden: true }) as HTMLOptionElement

    expect(getOptionBananas().getAttribute('aria-selected')).toBe('true') // mantine does not set selected

    await waitFor(async () => {
      const dropdownMenu = screen.getByTestId('select')
      await userEvent.click(screen.getByRole('option', { name: 'Broccoli', hidden: true }))
    })
    expect(getOptionBroccoli().getAttribute('aria-selected')).toBe('true')

    expect(handleChange).toHaveBeenCalledTimes(1)
    expect(handleChange).toHaveBeenCalledWith('Broccoli')
  })
})

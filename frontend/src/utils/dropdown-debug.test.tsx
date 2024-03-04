import { fireEvent, render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { Combobox, Group, InputBase, MantineProvider, Select, Text } from '@mantine/core'
import { vitest } from 'vitest'
import { SelectOptionComponentDebug } from 'src/utils/dropdown-debug'

it('should correctly set default option', async () => {
  render(
    <select data-testid="select state">
      <option>CA</option>
      <option>NY</option>
      <option selected>MA</option>
    </select>,
  )
  const optionMA = screen.getByRole('option', { name: 'MA' }) as HTMLOptionElement
  const optionCA = screen.getByRole('option', { name: 'CA' }) as HTMLOptionElement
  const combobox = screen.getAllByTestId('select state')[0] as HTMLSelectElement
  // await userEvent.click(combobox, { pointerState: await userEvent.pointer({ target: combobox }) })
  expect(optionCA.selected).toBe(false)
  expect(optionMA.selected).toBe(true)
  await userEvent.selectOptions(combobox, 'CA')
  expect(optionCA.selected).toBe(true)
  expect(optionMA.selected).toBe(false)
})

describe('Select component', () => {
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

describe('Combobox component', () => {
  it('calls onChange handler when an option is selected', async () => {
    const handleChange = vitest.fn()

    // TODO:
    render(
      <MantineProvider>
        <SelectOptionComponentDebug onSubmit={handleChange} />
      </MantineProvider>,
    )

    await waitFor(async () => {
      const dropdownMenu = screen.getByTestId('select')

      await userEvent.click(screen.getByText('Broccoli'))
    })

    expect(handleChange).toHaveBeenCalledTimes(1)
    expect(handleChange).toHaveBeenCalledWith('Broccoli')
  })
})

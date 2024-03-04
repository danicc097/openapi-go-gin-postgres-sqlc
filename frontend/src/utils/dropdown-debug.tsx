import { useState } from 'react'
import { Combobox, Group, Input, InputBase, Text, useCombobox } from '@mantine/core'
import { Virtuoso, VirtuosoMockContext } from 'react-virtuoso'
import { ArrayElement } from 'src/types/utils'

interface Item {
  emoji: string
  value: string
  description: string
}

const options: Item[] = [
  { emoji: '🍎', value: 'Apples', description: 'Crisp and refreshing fruit' },
  { emoji: '🍌', value: 'Bananas', description: 'Naturally sweet and potassium-rich fruit' },
  { emoji: '🥦', value: 'Broccoli', description: 'Nutrient-packed green vegetable' },
  { emoji: '🥕', value: 'Carrots', description: 'Crunchy and vitamin-rich root vegetable' },
  { emoji: '🍫', value: 'Chocolate', description: 'Indulgent and decadent treat' },
]

function SelectOption({ emoji, value, description }: Item) {
  return (
    <Group>
      <Text fz={20}>{emoji}</Text>
      <div>
        <Text fz="sm" fw={500}>
          {value}
        </Text>
        <Text fz="xs" opacity={0.6}>
          {description}
        </Text>
      </div>
    </Group>
  )
}

interface Props {
  onSubmit: (val: string) => void
  defaultOption?: string
}

export function SelectOptionComponentDebug({ onSubmit, defaultOption }: Props) {
  const combobox = useCombobox({
    onDropdownClose: () => combobox.resetSelectedOption(),
  })

  const [value, setValue] = useState<string | null>(defaultOption ?? null)
  const selectedOption = options.find((item) => item.value === value)

  return (
    <Combobox
      store={combobox}
      withinPortal={true}
      onOptionSubmit={(val) => {
        onSubmit(val)
        setValue(val)
        combobox.closeDropdown()
      }}
    >
      <Combobox.Target>
        <InputBase
          data-testid="select"
          component="button"
          type="button"
          pointer
          rightSection={<Combobox.Chevron />}
          onClick={() => combobox.toggleDropdown()}
          rightSectionPointerEvents="none"
          multiline
        >
          {selectedOption ? <SelectOption {...selectedOption} /> : <Input.Placeholder>Pick value</Input.Placeholder>}
        </InputBase>
      </Combobox.Target>

      <Combobox.Dropdown>
        <Combobox.Options>
          <Virtuoso
            style={{ height: '200px' }}
            totalCount={options.length}
            itemContent={(index) => (
              <Combobox.Option
                aria-selected={options[index]!.value === value}
                value={options[index]!.value}
                key={value}
                aria-label={`${options[index]!.value}`}
              >
                {options[index]?.value}
              </Combobox.Option>
            )}
          ></Virtuoso>
        </Combobox.Options>
      </Combobox.Dropdown>
    </Combobox>
  )
}

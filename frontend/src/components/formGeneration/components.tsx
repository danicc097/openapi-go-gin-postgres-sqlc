import { ColorInput } from '@mantine/core'
import { colorBlindPalette } from 'src/utils/colors'
import { InputOptions } from 'src/utils/formGeneration'

export const colorSwatchComponentInputOption: InputOptions<string, unknown>['component'] = (
  <ColorInput disallowInput withPicker={false} swatches={colorBlindPalette} styles={{ root: { width: '100%' } }} />
)

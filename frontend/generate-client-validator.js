import { join } from 'path'
import { generate } from 'openapi-typescript-validator'
import { fileURLToPath } from 'url'
import path from 'path'

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

generate({
  schemaFile: join(__dirname, '../openapi.yaml'),
  schemaType: 'yaml',
  directory: join(__dirname, 'src/client-validator/gen'),
  formatOptions: { mode: 'full' },
  prettierOptions: {
    printWidth: 120,
    semi: false,
    singleQuote: true,
    tabWidth: 2,
    trailingComma: 'all',
    parser: 'typescript',
  },
})

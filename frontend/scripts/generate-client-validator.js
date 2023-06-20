import { join } from 'path'
import { generate } from 'openapi-typescript-validator'
import { fileURLToPath } from 'url'
import path from 'path'
import $RefParser from 'json-schema-ref-parser'
import fs from 'fs'

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)
const root = path.resolve(__dirname, '..')

generate({
  schemaFile: join(root, '../openapi.yaml'),
  schemaType: 'yaml',
  directory: join(root, 'src/client-validator/gen'),
  prettierOptions: {
    printWidth: 120,
    semi: false,
    singleQuote: true,
    tabWidth: 2,
    trailingComma: 'all',
    parser: 'typescript',
  },
  addFormats: true,
  formatOptions: { formats: ['int64', 'int32', 'binary', 'date-time'] },
  ajvOptions: { strict: false, allErrors: true },
})

const jsonSchemaFilePath = join(root, 'src/client-validator/gen/schema.json')
const outputFilePath = join(root, 'src/client-validator/gen/dereferenced-schema.json')

$RefParser
  .dereference(jsonSchemaFilePath)
  .then((schema) => {
    const schemaString = JSON.stringify(schema, null, 2)

    fs.writeFile(outputFilePath, schemaString, (err) => {
      if (err) {
        console.error('Error saving dereferenced schema:', err)
      } else {
        console.log('Dereferenced schema saved successfully.')
      }
    })
  })
  .catch((e) => console.error(e))

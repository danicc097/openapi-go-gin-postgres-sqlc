import { join } from 'path'
import { generate } from 'openapi-typescript-validator'
import { fileURLToPath } from 'url'
import path from 'path'
import $RefParser from 'json-schema-ref-parser'
import fs from 'fs'

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)
const root = path.resolve(__dirname, '..')

await generate({
  schemaFile: join(root, '../openapi.exploded.yaml'),
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
  formatOptions: { formats: ['int64', 'int32', 'binary', 'date-time', 'date'] },
  ajvOptions: { strict: false, allErrors: true },
})

const jsonSchemaFilePath = join(root, 'src/client-validator/gen/schema.json')

const schema = readSchemaFromFile(jsonSchemaFilePath)
modifyDateFormats(schema)
modifyStringValidation(schema)
deleteNestedKeys(schema, '$schema') // breaks ajv since it tries to resolve it
saveSchemaToFile(schema, jsonSchemaFilePath)

const outputFilePath = join(root, 'src/client-validator/gen/dereferenced-schema.json')

$RefParser
  .dereference(jsonSchemaFilePath)
  .then((schema) => {
    saveSchemaToFile(schema, outputFilePath)
  })
  .catch((e) => console.error(e))

function saveSchemaToFile(schema, filePath) {
  const schemaString = JSON.stringify(schema, null, 2)

  fs.writeFileSync(filePath, schemaString, (err) => {
    if (err) {
      throw new Error(`Error saving schema: ${err}`)
    }
  })
}

function modifyDateFormats(schema) {
  if (schema.type === 'array') {
    if (schema.items) {
      if (Array.isArray(schema.items)) {
        for (const itemSchema of schema.items) {
          modifyDateFormats(itemSchema)
          if (itemSchema.format === 'date-time' || itemSchema.format === 'date') {
            itemSchema.type = ['object']
            itemSchema.type = itemSchema.type.filter((type) => type !== 'string')
          }
        }
      } else {
        modifyDateFormats(schema.items)
        if (schema.items.format === 'date-time' || schema.items.format === 'date') {
          schema.items.type = ['object']
          schema.items.type = schema.items.type.filter((type) => type !== 'string')
        }
      }
    }
  } else if (schema.type === 'object' && schema.properties) {
    for (const propertyKey in schema.properties) {
      if (schema.properties.hasOwnProperty(propertyKey)) {
        const propertySchema = schema.properties[propertyKey]
        modifyDateFormats(propertySchema)
        if (propertySchema.format === 'date-time' || propertySchema.format === 'date') {
          propertySchema.type = ['object']
          propertySchema.type = propertySchema.type.filter((type) => type !== 'string')
        }
      }
    }
  }

  if (schema.definitions) {
    for (const definitionKey in schema.definitions) {
      if (schema.definitions.hasOwnProperty(definitionKey)) {
        const definitionSchema = schema.definitions[definitionKey]
        modifyDateFormats(definitionSchema)
      }
    }
  }
}

function modifyStringValidation(schema) {
  if (schema.type === 'array') {
    if (schema.items) {
      if (Array.isArray(schema.items)) {
        for (const itemSchema of schema.items) {
          modifyStringValidation(itemSchema)
          addMinLengthToRequiredStringTypes(itemSchema, schema.required)
        }
      } else {
        modifyStringValidation(schema.items)
        addMinLengthToRequiredStringTypes(schema.items, schema.required)
      }
    }
  } else if (schema.type === 'object' && schema.properties) {
    for (const propertyKey in schema.properties) {
      if (schema.properties.hasOwnProperty(propertyKey)) {
        const propertySchema = schema.properties[propertyKey]
        modifyStringValidation(propertySchema)
        addMinLengthToRequiredStringTypes(propertySchema, schema.required)
      }
    }
  }

  if (schema.definitions) {
    for (const definitionKey in schema.definitions) {
      if (schema.definitions.hasOwnProperty(definitionKey)) {
        const definitionSchema = schema.definitions[definitionKey]
        modifyStringValidation(definitionSchema)
      }
    }
  }

  if (schema.required && Array.isArray(schema.required)) {
    const requiredStringProperties = schema.required.filter((property) => {
      const propertySchema = schema.properties[property]
      return (
        propertySchema &&
        (propertySchema.type === 'string' ||
          (propertySchema.type === 'array' && propertySchema.items?.type === 'string'))
      )
    })

    requiredStringProperties.forEach((property) => {
      const propertySchema = schema.properties[property]
      if (!propertySchema.minLength) {
        propertySchema.minLength = 1
      }
    })
  }
}

function addMinLengthToRequiredStringTypes(schema, parentRequired) {
  if (
    parentRequired &&
    parentRequired.includes(schema.title) &&
    (schema.type === 'string' || (schema.type === 'array' && schema.items?.type === 'string'))
  ) {
    schema.minLength = 1
  }
}

function readSchemaFromFile(filePath) {
  const absolutePath = path.resolve(filePath)

  try {
    const fileContent = fs.readFileSync(absolutePath, 'utf-8')
    const schema = JSON.parse(fileContent)
    return schema
  } catch (error) {
    console.error(`Error reading schema from file: ${absolutePath}`)
    console.error(error)
    return null
  }
}

function deleteNestedKeys(obj, key) {
  for (const prop in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, prop)) {
      if (typeof obj[prop] === 'object') {
        deleteNestedKeys(obj[prop], key)
      } else if (prop === key) {
        delete obj[prop]
      }
    }
  }
}

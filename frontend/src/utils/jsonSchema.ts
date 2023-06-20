interface JsonSchemaField {
  type?: string | string[]
  format?: string
  items?: JsonSchemaField
  properties?: { [key: string]: JsonSchemaField }
  required?: string[]
}

type SchemaType = {
  type: string
  required: boolean
  isArray?: boolean
}

type FormGeneratorFields = {
  [key: string]: {
    type: string
    required: boolean
    isArray?: boolean
  }
}

export function extractFieldTypes(schema: JsonSchemaField): FormGeneratorFields {
  const result: { [key: string]: SchemaType } = {}

  function traverseSchema(field: JsonSchemaField, currentPath: string, parentArray = false, parent?: JsonSchemaField) {
    const requiredFields = new Set(field.required || [])

    const isFieldRequired = !requiredFields.has(currentPath)
    if (currentPath !== '') {
      if (field.format) {
        result[currentPath] = { type: field.format, required: isFieldRequired, isArray: parentArray }
      } else if (field.type) {
        const type = filterType(field)
        const isArray = Array.isArray(field.type) || parentArray
        result[currentPath] = { type, required: isFieldRequired, isArray }
      }
    }

    if (field.properties) {
      for (const key in field.properties) {
        const newPath = currentPath ? `${currentPath}.${key}` : key
        traverseSchema(field.properties[key], newPath, false, field)
      }
    }

    if (field.items && field.items.type === 'object') {
      const newPath = currentPath
      traverseSchema(field.items, newPath, true, field)
      result[newPath] = { type: 'arrayOfObject', required: isFieldRequired, isArray: true }
    }

    if (parent?.type === 'array' && field.items?.type && !field.items.properties) {
      const newPath = currentPath
      const isArray = true
      const type = filterType(field.items)
      result[newPath] = { type, required: isFieldRequired, isArray }
    }
  }

  traverseSchema(schema, '', false)

  return result
}
function filterType(field: JsonSchemaField) {
  // TODO return format if exists
  return Array.isArray(field.type) ? field.type.filter((type) => type !== 'null')[0] : field.type
}

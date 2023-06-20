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

export function extractFieldTypes(schema: JsonSchemaField): {
  [key: string]: { type: string; required: boolean; isArray?: boolean }
} {
  const result: { [key: string]: SchemaType } = {}

  function traverseSchema(field: JsonSchemaField, currentPath: string, parentRequired: boolean, parentArray = false) {
    const requiredFields = new Set(field.required || [])
    const isParentRequired = parentRequired || requiredFields.size > 0

    if (currentPath !== '') {
      if (field.type && (typeof field.type === 'string' || Array.isArray(field.type))) {
        const type = Array.isArray(field.type) ? field.type[0] : field.type
        const isArray = Array.isArray(field.type) || parentArray

        const isFieldRequired = isParentRequired && !requiredFields.has(currentPath)
        result[currentPath] = { type, required: isFieldRequired, isArray }
      }

      if (field.format) {
        const isFieldRequired = isParentRequired && !requiredFields.has(currentPath)
        result[currentPath] = { type: field.format, required: isFieldRequired, isArray: parentArray }
      }
    }

    if (field.properties) {
      for (const key in field.properties) {
        const newPath = currentPath ? `${currentPath}.${key}` : key
        traverseSchema(field.properties[key], newPath, isParentRequired)
      }
    }

    if (field.items) {
      const newPath = currentPath
      traverseSchema(field.items, newPath, isParentRequired, true)
    }
  }

  traverseSchema(schema, '', false)

  return result
}

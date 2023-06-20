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
      if (field.format) {
        const isFieldRequired = isParentRequired && !requiredFields.has(currentPath)
        result[currentPath] = { type: field.format, required: isFieldRequired, isArray: parentArray }
      } else if (field.type) {
        const types = Array.isArray(field.type) ? field.type.filter((type) => type !== 'null') : [field.type]
        const type = types.filter((t) => t !== 'null')[0]
        const isArray = Array.isArray(field.type) || parentArray
        const isFieldRequired = isParentRequired && !requiredFields.has(currentPath)
        result[currentPath] = { type, required: isFieldRequired, isArray }
      }
    }

    if (field.properties) {
      for (const key in field.properties) {
        const newPath = currentPath ? `${currentPath}.${key}` : key
        traverseSchema(field.properties[key], newPath, isParentRequired)
      }
    }

    if (field.items && field.items.type === 'object') {
      const newPath = currentPath
      traverseSchema(field.items, newPath, isParentRequired, true)
      result[newPath] = { type: 'arrayOfObject', required: isParentRequired, isArray: true }
    }
  }

  traverseSchema(schema, '', false)

  return result
}

import JSON_SCHEMA from '../src/client-validator/gen/dereferenced-schema.json' assert { type: 'json' }

const projects = JSON_SCHEMA.definitions.Project.enum

const workItemResponses = {
  demo: JSON_SCHEMA.definitions.RestDemoWorkItemsResponse,
  demo_two: JSON_SCHEMA.definitions.RestDemoTwoWorkItemsResponse,
}

const schemaFields = Object.entries(workItemResponses).reduce((acc, [project, schema]) => {
  if (!projects.includes(project))
    throw new Error(`Project '${project}' does not exist. Existing projects: ${projects}`)

  acc[project] = parseSchemaFields(schema)

  return acc
}, {})

// instead of a global config mess we can have per-user config in local storage for visualization.
// and get rid of current project.boardConfig fields (will be used for other actually useful generic settings)
console.log(schemaFields)

// adapted from src/utils/jsonSchema.ts
function parseSchemaFields(schema) {
  const schemaFields = {}

  function traverseSchema(obj, path = []) {
    const isArrayOfObj = extractType(obj) === 'object' && !!obj.type?.includes('array') && obj.items?.properties
    const isObj = obj.properties && extractType(obj) === 'object'

    if (isArrayOfObj) {
      extract(obj?.items?.properties)
    } else if (isObj) {
      extract(obj.properties)
    }

    function extract(properties) {
      for (const key in properties) {
        const newPath = [...path, key]
        const property = properties[key]
        if (!property) {
          continue
        }

        schemaFields[newPath.join('.')] = null
        traverseSchema(property, newPath)
      }
    }
  }

  traverseSchema(schema)

  return Object.keys(schemaFields)
}

function extractType(obj) {
  const type = _type(obj.type)
  if (type === 'array') {
    if (obj?.items?.type === 'object') {
      return 'object'
    } else {
      return _type(obj?.items?.type)
    }
  }

  return obj.format ?? type

  function _type(x) {
    return Array.isArray(x) ? x.filter((t) => t !== 'null')[0] : x
  }
}

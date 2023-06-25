import type { GetKeys, RecursiveKeyOf, RecursiveKeyOfArray, TypeOf } from 'src/types/utils'
import DynamicForm, { constructFormKey } from 'src/utils/formGeneration'
import { parseSchemaFields, type JsonSchemaField, type SchemaField } from 'src/utils/jsonSchema'
import { describe, expect, test } from 'vitest'
import { getByTestId, render, screen, renderHook } from '@testing-library/react'
import '@testing-library/jest-dom'
import dayjs from 'dayjs'
import { useForm } from '@mantine/form'

const formInitialValues = {
  base: {
    items: [
      { items: ['0001', '0002'], name: 'item-1' },
      { items: ['0011', '0012'], name: 'item-2' },
    ],
    closed: dayjs('2023-03-24T20:42:00.000Z').toDate(),
    targetDate: dayjs('2023-02-22').toDate(),
    description: 'some text',
    kanbanStepID: 1,
    teamID: 1,
    // title: {},
    workItemTypeID: 1,
  },
  // TODO: need to check runtime type, else all fails catastrophically.
  // it should update the form but show callout error saying ignoring bad type in `formField`, in this case `tagIDs.1`
  // tagIDs: [1, 'fsfefes'], // {"invalidParams":{"name":"tagIDs.1","reason":"must be integer"} and we can set invalid manually via component id (which will be `input-tagIDs.1` )
  demoProject: {
    lastMessageAt: dayjs('2023-03-24T20:42:00.000Z').toDate(),
    line: '3e3e2',
    ref: '312321',
    workItemID: 1,
  },
  tagIDs: [0, 1, 2],
  members: [
    { role: 'preparer', userID: 'user 1' },
    { role: 'preparer', userID: 'user 2' },
  ],
} as TestTypes.RestDemoWorkItemCreateRequest

const schemaFields: Record<GetKeys<TestTypes.RestDemoWorkItemCreateRequest>, SchemaField> = {
  base: { isArray: false, required: true, type: 'object' },
  'base.closed': { type: 'date-time', required: true, isArray: false },
  'base.description': { type: 'string', required: true, isArray: false },
  'base.kanbanStepID': { type: 'integer', required: true, isArray: false },
  'base.metadata': { type: 'integer', required: true, isArray: true },
  'base.targetDate': { type: 'date', required: true, isArray: false },
  'base.teamID': { type: 'integer', required: true, isArray: false },
  'base.items': { type: 'object', required: true, isArray: true },
  'base.items.name': { type: 'string', required: true, isArray: false },
  'base.items.items': { type: 'string', required: true, isArray: true },
  'base.workItemTypeID': { type: 'integer', required: true, isArray: false },
  demoProject: { isArray: false, required: true, type: 'object' },
  'demoProject.lastMessageAt': { type: 'date-time', required: true, isArray: false },
  'demoProject.line': { type: 'string', required: true, isArray: false },
  'demoProject.ref': { type: 'string', required: true, isArray: false },
  'demoProject.reopened': { type: 'boolean', required: true, isArray: false },
  'demoProject.workItemID': { type: 'integer', required: true, isArray: false },
  members: { type: 'object', required: true, isArray: true },
  'members.role': { type: 'string', required: true, isArray: false },
  'members.userID': { type: 'string', required: true, isArray: false },
  tagIDs: { type: 'integer', required: true, isArray: true },
}

const schema = {
  properties: {
    base: {
      properties: {
        closed: {
          format: 'date-time',
          type: ['string', 'null'],
        },
        description: {
          type: 'string',
        },
        kanbanStepID: {
          type: 'integer',
        },
        metadata: {
          items: {
            minimum: 0,
            type: 'integer',
          },
          type: ['array', 'null'],
        },
        targetDate: {
          format: 'date',
          type: 'string',
        },
        teamID: {
          type: 'integer',
        },
        items: {
          items: {
            properties: {
              items: {
                items: {
                  type: 'string',
                },
                type: ['array', 'null'],
              },
              name: {
                type: 'string',
                $schema: 'http://json-schema.org/draft-04/schema#',
              },
            },
            required: ['items', 'name'],
            type: 'object',
            $schema: 'http://json-schema.org/draft-04/schema#',
          },
          type: ['array', 'null'],
        },
        workItemTypeID: {
          type: 'integer',
        },
      },
      required: [
        'items',
        'description',
        'workItemTypeID',
        'metadata',
        'teamID',
        'kanbanStepID',
        'closed',
        'targetDate',
      ],
      type: 'object',
      $schema: 'http://json-schema.org/draft-04/schema#',
    },
    demoProject: {
      properties: {
        lastMessageAt: {
          format: 'date-time',
          type: 'string',
        },
        line: {
          type: 'string',
        },
        ref: {
          pattern: '^[0-9]{8}$',
          type: 'string',
        },
        reopened: {
          type: 'boolean',
        },
        workItemID: {
          type: 'integer',
        },
      },
      required: ['workItemID', 'ref', 'line', 'lastMessageAt', 'reopened'],
      type: 'object',
      $schema: 'http://json-schema.org/draft-04/schema#',
    },
    members: {
      items: {
        properties: {
          role: {
            title: 'WorkItem role',
            type: 'string',
            'x-generated': '-',
            enum: ['preparer', 'reviewer'],
            description: "represents a database 'work_item_role'",
            $schema: 'http://json-schema.org/draft-04/schema#',
          },
          userID: {
            type: 'string',
            $schema: 'http://json-schema.org/draft-04/schema#',
          },
        },
        required: ['userID', 'role'],
        type: 'object',
        $schema: 'http://json-schema.org/draft-04/schema#',
      },
      type: ['array', 'null'],
    },
    tagIDs: {
      items: {
        type: 'integer',
      },
      type: ['array', 'null'],
    },
  },
  required: ['demoProject', 'base', 'tagIDs', 'members'],
  type: 'object',
  'x-postgen-struct': 'RestDemoWorkItemCreateRequest',
  $schema: 'http://json-schema.org/draft-04/schema#',
} as JsonSchemaField

describe('parseSchemaFields', () => {
  test('should extract field types correctly from a JSON schema', () => {
    /**

    form generator will use these keys. to generate multiple forms when is array we just check
    if parent (split by . and keep up to len-2) isArray (members) or the child itself isArray (tagIDs)

    it doesnt seem to be easy to get typed keys for these when arrays are involved.
    */

    expect(parseSchemaFields(schema)).toEqual(schemaFields)
  })

  test('should render form fields and buttons', () => {
    const { result } = renderHook(() =>
      useForm({
        initialValues: formInitialValues,
      }),
    )

    render(
      <DynamicForm<TestTypes.RestDemoWorkItemCreateRequest>
        name="demoWorkItemCreateForm"
        schemaFields={schemaFields}
        form={result.current}
        options={{
          defaultValue: {
            'demoProject.line': '534543523',
            members: [{ role: 'preparer', userID: 'c446259c-1083-4212-98fe-bd080c41e7d7' }],
          },
        }}
      />,
    )

    // Recursive function to test data-testid attributes
    const testField = (fieldKey, parentFormField = '') => {
      const formField = constructFormKey(fieldKey, parentFormField)
      const field = schemaFields[fieldKey]

      if (field.isArray) {
        const addButtonTestId =
          fieldKey === parentFormField
            ? `form-field-add-button-${formField}`
            : `form-field-add-button-${parentFormField}.${fieldKey}`
        console.log('addButtonTestId:', addButtonTestId)
        const addButtonElement = screen.queryByTestId(addButtonTestId)
        console.log('addButtonElement:', addButtonElement)

        if (addButtonElement) {
          expect(addButtonElement).toBeInTheDocument()

          // Test data-testid attribute for each array element and remove button
          const arrayElements = formInitialValues[fieldKey]
          arrayElements.forEach((_, index) => {
            const arrayElementFormField = `${formField}.${index}`
            testField(fieldKey, arrayElementFormField)

            const removeButtonTestId = `form-field-remove-button-${arrayElementFormField}`
            console.log('removeButtonTestId:', removeButtonTestId)
            const removeButtonElement = screen.queryByTestId(removeButtonTestId)
            console.log('removeButtonElement:', removeButtonElement)

            expect(removeButtonElement).toBeInTheDocument()
          })
        } else {
          console.log(`Add button not found for field: ${formField}`)
        }
      } else if (field.type === 'object') {
        // Skip if field type is "object" and not an array
        return
      } else {
        const dataTestId = `form-field-${formField}`
        console.log('dataTestId:', dataTestId)
        const fieldElement = screen.getByTestId(dataTestId)
        console.log('fieldElement:', fieldElement)

        expect(fieldElement).toBeInTheDocument()

        // Test data-testid attribute for nested fields
        Object.keys(field).forEach((subFieldKey) => {
          const nestedFormField = constructFormKey(subFieldKey, formField)
          testField(subFieldKey, nestedFormField)
        })
      }
    }

    // Start testing with top-level fields
    Object.keys(formInitialValues).forEach((fieldKey) => {
      testField(fieldKey)
    })
  })
})

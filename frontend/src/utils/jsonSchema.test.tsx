import { type FieldPath, type PathValue } from 'react-hook-form'
import type { RecursiveKeyOf, RecursiveKeyOfArray } from 'src/types/utils'
import DynamicForm from 'src/utils/formGeneration'
import { parseSchemaFields, type JsonSchemaField, type SchemaField } from 'src/utils/jsonSchema'
import { describe, expect, test } from 'vitest'
import { getByTestId, render, screen, renderHook } from '@testing-library/react'
import '@testing-library/jest-dom'
import dayjs from 'dayjs'
import { useForm } from '@mantine/form'

type RestDemoWorkItemCreateRequestFormField =
  // hack to use 'members.role' instead of 'members.??.role'
  | FieldPath<TestTypes.RestDemoWorkItemCreateRequest>
  | RecursiveKeyOfArray<TestTypes.RestDemoWorkItemCreateRequest['members'], 'members'>

const schemaFields: Record<RestDemoWorkItemCreateRequestFormField, SchemaField> = {
  base: { isArray: false, required: true, type: 'object' },
  'base.closed': { type: 'date-time', required: true, isArray: false },
  'base.description': { type: 'string', required: true, isArray: false },
  'base.kanbanStepID': { type: 'integer', required: true, isArray: false },
  'base.metadata': { type: 'integer', required: true, isArray: true },
  'base.targetDate': { type: 'date-time', required: true, isArray: false },
  'base.teamID': { type: 'integer', required: true, isArray: false },
  'base.packs': { type: 'object', required: true, isArray: true },
  // FIXME: FieldPath allows 'base.packs.<i>.name' only.
  // need to remove TupleKeys from allowed path indexing
  // maybe just bring that code over and update
  // NOTE: WORKAROUND: or maybe could have convention to use .0 indexing and  update parseSchemaFields to include that `.0` those (or simply replaceAll `.0` with "" and leave parseSchemaFields as is).
  // although in that case we get no keys autocomplete for nested array elements
  'base.packs.name': { type: 'string', required: true, isArray: false },
  'base.packs.items': { type: 'string', required: true, isArray: true },
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

type a = PathValue<TestTypes.RestDemoWorkItemCreateRequest, 'base.packs'>

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
          format: 'date-time',
          type: 'string',
        },
        teamID: {
          type: 'integer',
        },
        packs: {
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
        'title',
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
        initialValues: {
          base: {
            closed: dayjs('20-2-2020').toDate(),
            targetDate: dayjs('20-2-2020').toDate(),
            description: 'some text',
            kanbanStepID: 1,
            teamID: 1,
            packs: {},
            workItemTypeID: 1,
          },
          // TODO: test validation error callout for generated form
          // tagIDs: [1, 'fsfefes'], // {"invalidParams":{"name":"tagIDs.1","reason":"must be integer"} and we can set invalid manually via component id (which will be `input-tagIDs.1` )
          demoProject: {
            lastMessageAt: dayjs().toDate(),
            line: '3e3e2',
            ref: '312321',
            workItemID: 1,
          },
          tagIDs: [0, 1, 2],
          members: [
            { role: 'preparer', userID: 'user 1' },
            { role: 'reviewer', userID: 'user 2' },
          ],
        } as TestTypes.RestDemoWorkItemCreateRequest,
      }),
    )

    render(
      <DynamicForm
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

    // TODO: test matches. update initialValues to test out multiple nested
  })
})

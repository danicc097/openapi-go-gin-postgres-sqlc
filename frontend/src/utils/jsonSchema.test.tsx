import type { DeepPartial, GetKeys, RecursiveKeyOf, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import DynamicForm from 'src/utils/formGeneration'
import { parseSchemaFields, type JsonSchemaField, type SchemaField } from 'src/utils/jsonSchema'
import { describe, expect, test } from 'vitest'
import { getByTestId, render, screen, renderHook } from '@testing-library/react'
import '@testing-library/jest-dom'
import dayjs from 'dayjs'
import { entries, keys } from 'src/utils/object'
import { FormProvider, useForm } from 'react-hook-form'
import { ajvResolver } from '@hookform/resolvers/ajv'

const schema = {
  properties: {
    base: {
      properties: {
        closed: {
          format: 'date-time',
          type: ['object', 'null'],
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
          type: 'object',
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
          type: 'object',
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
} as JsonSchemaField

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
    { role: null, userID: 'a446259c-1083-4212-98fe-bd080c41e7d7' }, // with defaultValue of "member.role": {role: 'preparer'} it will fill null or undefined form values
    { role: 'reviewer', userID: 'b446259c-1083-4212-98fe-bd080c41e7d7' },
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
    const { result: form } = renderHook(() =>
      useForm<TestTypes.RestDemoWorkItemCreateRequest>({
        resolver: ajvResolver(schema as any, {
          strict: false,
          // formats: fullFormats,
        }),
        mode: 'onChange',
        defaultValues: formInitialValues ?? {},
        // shouldUnregister: true, // defaultValues will not be merged against submission result.
      }),
    )

    const formName = 'demoWorkItemCreateForm'

    const view = render(
      <FormProvider {...form.current}>
        <DynamicForm<TestTypes.RestDemoWorkItemCreateRequest, 'base.metadata'>
          formName={formName}
          schemaFields={schemaFields}
          options={{
            labels: {
              base: null,
              'base.closed': 'closed',
              'base.description': 'description',
              // 'base.metadata': 'metadata', // ignored -> not a key
              'base.kanbanStepID': 'kanbanStepID', // if using KanbanStep transformer, then "Kanban step", "Kanban step name", etc.
              'base.targetDate': 'targetDate',
              'demoProject.reopened': 'reopened',
              'base.teamID': 'teamID',
              'base.items': 'items',
              'base.items.name': 'name',
              'base.items.items': 'items',
              'base.workItemTypeID': 'workItemTypeID',
              demoProject: null,
              'demoProject.lastMessageAt': 'lastMessageAt',
              'demoProject.line': 'line',
              'demoProject.ref': 'ref',
              'demoProject.workItemID': 'workItemID',
              members: 'members',
              'members.role': 'role',
              'members.userID': 'User',
              tagIDs: 'tagIDs',
            },
            defaultValues: {
              'demoProject.line': '43121234',
              // FIXME: does not infer return type like selectOptions does. this should fail
              'members.role': 'prepadrer',
            },
            selectOptions: {},
          }}
        />
      </FormProvider>,
    )

    const ids = [
      'demoWorkItemCreateForm-base.description',
      'demoWorkItemCreateForm-base.kanbanStepID',
      'demoWorkItemCreateForm-base.teamID',
      'demoWorkItemCreateForm-base.items-add-button',
      'demoWorkItemCreateForm-base.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.0.name',
      'demoWorkItemCreateForm-base.items.0.items-add-button',
      'demoWorkItemCreateForm-base.items.0.items-0',
      'demoWorkItemCreateForm-base.items.0.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.0.items-1',
      'demoWorkItemCreateForm-base.items.0.items-remove-button-1',
      'demoWorkItemCreateForm-base.items-remove-button-1',
      'demoWorkItemCreateForm-base.items.1.name',
      'demoWorkItemCreateForm-base.items.1.items-add-button',
      'demoWorkItemCreateForm-base.items.1.items-0',
      'demoWorkItemCreateForm-base.items.1.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.1.items-1',
      'demoWorkItemCreateForm-base.items.1.items-remove-button-1',
      'demoWorkItemCreateForm-base.workItemTypeID',
      'demoWorkItemCreateForm-demoProject.line',
      'demoWorkItemCreateForm-demoProject.ref',
      'demoWorkItemCreateForm-demoProject.reopened',
      'demoWorkItemCreateForm-demoProject.workItemID',
      'demoWorkItemCreateForm-members-add-button',
      'demoWorkItemCreateForm-members-remove-button-0',
      'demoWorkItemCreateForm-members.0.role',
      'demoWorkItemCreateForm-members.0.userID',
      'demoWorkItemCreateForm-members-remove-button-1',
      'demoWorkItemCreateForm-members.1.role',
      'demoWorkItemCreateForm-members.1.userID',
      'demoWorkItemCreateForm-tagIDs-add-button',
      'demoWorkItemCreateForm-tagIDs-0',
      'demoWorkItemCreateForm-tagIDs-remove-button-0',
      'demoWorkItemCreateForm-tagIDs-1',
      'demoWorkItemCreateForm-tagIDs-remove-button-1',
      'demoWorkItemCreateForm-tagIDs-2',
      'demoWorkItemCreateForm-tagIDs-remove-button-2',
    ]

    ids.forEach((id) => {
      const el = document.getElementById(id)
      expect(el, `${id} not found`).toBeTruthy()
      expect(el).toBeInTheDocument()
    })

    const titleDataTestIds = [
      'base-title',
      'base.items-title',
      'base.items.0.items-title',
      'base.items.1.items-title',
      'demoProject-title',
      'members-title',
      'tagIDs-title',
    ]
    titleDataTestIds.forEach((id) => {
      expect(view.getByTestId(id)).toBeInTheDocument()
    })

    test('should update form with default values', () => {
      // defaultValues: {
      //   'demoProject.line': '43121234',
      //   members: [
      //     { role: 'preparer', userID: 'a446259c-1083-4212-98fe-bd080c41e7d7' },
      //     { role: 'reviewer', userID: 'b446259c-1083-4212-98fe-bd080c41e7d7' },
      //   ],
      // },
      // TODO: get input by id
    })
  })
})

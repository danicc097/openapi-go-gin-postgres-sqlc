import type { DeepPartial, GetKeys, RecursiveKeyOf, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import DynamicForm from 'src/utils/formGeneration'
import { parseSchemaFields, type SchemaField } from 'src/utils/jsonSchema'
import { describe, expect, test } from 'vitest'
import { getByTestId, render, screen, renderHook, fireEvent, act, getByText } from '@testing-library/react'
import '@testing-library/jest-dom'
import dayjs from 'dayjs'
import { entries, keys } from 'src/utils/object'
import { FormProvider, useForm } from 'react-hook-form'
import { ajvResolver } from '@hookform/resolvers/ajv'
import { fullFormats } from 'ajv-formats/dist/formats'
import { Group, Avatar, Space, Flex, MantineProvider } from '@mantine/core'
import { nameInitials } from 'src/utils/strings'
import { JSONSchemaType } from 'ajv'
import { selectOptionsBuilder } from 'src/utils/formGeneration.context'

const tags = [...Array(10)].map((x, i) => {
  return {
    name: `${i} tag`,
    color: `#${i}34236`,
    workItemTagID: i,
    projectID: 1,
    description: 'description',
  }
})

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
          type: 'object',
        },
        targetDate: {
          format: 'date',
          type: 'object',
        },
        teamID: {
          type: 'integer',
        },
        // purposely name them nested items to ensure correct recursion
        items: {
          items: {
            properties: {
              items: {
                items: {
                  type: 'string',
                },
                type: ['array', 'null'],
              },
              userId: {
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
            required: ['userId', 'items', 'name'],
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
    tagIDsMultiselect: {
      items: {
        type: 'integer',
      },
      type: ['array', 'null'],
    },
  },
  required: ['demoProject', 'base', 'tagIDsMultiselect', 'members'],
  type: 'object',
  'x-gen-struct': 'RestDemoWorkItemCreateRequest',
} as JSONSchemaType<true>

const formInitialValues = {
  base: {
    items: [
      { items: ['0001', '0002'], userId: [], name: 'item-1' },
      { items: ['0011', '0012'], userId: [], name: 'item-2' },
    ],
    closed: dayjs('2023-03-24T20:42:00.000Z').toDate(),
    // targetDate: dayjs('2023-02-22').toDate(),
    description: 'some text',
    kanbanStepID: 1,
    teamID: 1,
    metadata: {},
    workItemTypeID: 1,
    targetDate: dayjs('2024-03-24T20:42:00.000Z').toDate(),
  },
  demoProject: {
    lastMessageAt: dayjs('2023-03-24T20:42:00.000Z').toDate(),
    line: '3e3e2',
    ref: '124321', // should fail pattern validation
    workItemID: 1,
    reopened: true,
  },
  tagIDs: [0, 1, 2],
  tagIDsMultiselect: [0, 1, 2],
  members: [
    // with defaultValue of "member.role": {role: 'preparer'} it will fill null or undefined form values.
    // since userid exists and it's an initial value, it will show custom select card to work around https://github.com/mantinedev/mantine/issues/980
    // therefore its element input id does not exist
    { userID: 'a446259c-1083-4212-98fe-bd080c41e7d7' },
    // userid does not exist in selectOptions users -> will show input directly instead
    { role: 'reviewer', userID: 'b446259c-1083-4212-98fe-bd080c41e7d7' },
  ],
} as TestTypes.DemoWorkItemCreateRequest

const schemaFields: Record<GetKeys<TestTypes.DemoWorkItemCreateRequest>, SchemaField> = {
  base: { isArray: false, required: true, type: 'object' },
  'base.closed': { type: 'date-time', required: false, isArray: false },
  'base.description': { type: 'string', required: true, isArray: false },
  'base.kanbanStepID': { type: 'integer', required: true, isArray: false },
  'base.metadata': { type: 'object', required: true, isArray: false },
  'base.targetDate': { type: 'date', required: true, isArray: false },
  'base.teamID': { type: 'integer', required: true, isArray: false },
  'base.items': { type: 'object', required: false, isArray: true },
  'base.items.name': { type: 'string', required: true, isArray: false },
  'base.items.userId': { type: 'string', required: false, isArray: true },
  'base.items.items': { type: 'string', required: false, isArray: true },
  'base.workItemTypeID': { type: 'integer', required: true, isArray: false },
  demoProject: { isArray: false, required: true, type: 'object' },
  'demoProject.lastMessageAt': { type: 'date-time', required: true, isArray: false },
  'demoProject.line': { type: 'string', required: true, isArray: false },
  'demoProject.ref': { type: 'string', required: true, isArray: false },
  'demoProject.reopened': { type: 'boolean', required: true, isArray: false },
  'demoProject.workItemID': { type: 'integer', required: true, isArray: false },
  members: { type: 'object', required: false, isArray: true },
  'members.role': { type: 'string', required: true, isArray: false },
  'members.userID': { type: 'string', required: true, isArray: false },
  tagIDs: { type: 'integer', required: false, isArray: true },
  tagIDsMultiselect: { type: 'integer', required: false, isArray: true },
}

describe('form generation', () => {
  test('should extract field types correctly from a JSON schema', () => {
    expect(parseSchemaFields(schema)).toEqual(schemaFields)
  })

  test('should render form fields and buttons', async () => {
    /**
     * FIXME: no need for renderHook. test via ui.
     * https://react-hook-form.com/advanced-usage#TestingForm
     */
    const { result: form } = renderHook(() =>
      useForm<TestTypes.DemoWorkItemCreateRequest>({
        resolver: ajvResolver(schema as any, {
          strict: false,
          formats: fullFormats,
        }),
        mode: 'onChange',
        defaultValues: formInitialValues ?? {},
        // shouldUnregister: true, // defaultValues will not be merged against submission result.
      }),
    )

    const formName = 'demoWorkItemCreateForm'

    const { isDirty, isSubmitting, submitCount } = form.current.formState

    const view = render(
      <MantineProvider>
        <FormProvider {...form.current}>
          <DynamicForm<TestTypes.DemoWorkItemCreateRequest, 'base.metadata'>
            onSubmit={(e) => {
              e.preventDefault()
              form.current.handleSubmit(
                (data) => {
                  console.log({ data })
                },
                (errors) => {
                  console.log({ errors })
                },
              )(e)
            }}
            formName={formName}
            schemaFields={schemaFields}
            options={{
              renderOrderPriority: ['tagIDs', 'members'],
              labels: {
                base: 'base', // just title via renderTitle
                'base.closed': 'closed',
                'base.description': 'description',
                // 'base.metadata': 'metadata', // ignored -> not a key
                'base.kanbanStepID': 'kanbanStepID',
                'base.targetDate': 'targetDate',
                'demoProject.reopened': 'reopened',
                'base.teamID': 'teamID',
                'base.items': 'items',
                'base.items.name': 'name',
                'base.items.items': 'items',
                'base.items.userId': 'user',
                'base.workItemTypeID': 'workItemTypeID',
                demoProject: null, // won't render title
                'demoProject.lastMessageAt': 'lastMessageAt',
                'demoProject.line': 'line',
                'demoProject.ref': 'ref',
                'demoProject.workItemID': 'workItemID',
                members: 'members',
                'members.role': 'role',
                'members.userID': 'User',
                tagIDs: 'tagIDs',
                tagIDsMultiselect: 'tagIDsMultiselect',
              },
              defaultValues: {
                'demoProject.line': '43121234', // should be ignored since it's set
                'members.role': 'preparer',
              },
              selectOptions: {
                'members.userID': selectOptionsBuilder({
                  type: 'select',
                  values: [...Array(1)].map((x, i) => {
                    const user = {
                      username: '1',
                      email: '1@mail.com',
                      userID: 'a446259c-1083-4212-98fe-bd080c41e7d7',
                    }
                    return user
                  }),
                  optionTransformer(el) {
                    return (
                      <Group align="center">
                        <div style={{ display: 'flex', alignItems: 'center' }}>
                          <Avatar size={35} radius="xl" data-test-id="header-profile-avatar" alt={el?.username}>
                            {nameInitials(el?.email || '')}
                          </Avatar>
                          <Space p={5} />
                        </div>

                        <div style={{ marginLeft: 'auto' }}>{el?.email}</div>
                      </Group>
                    )
                  },
                  formValueTransformer(el) {
                    return el.userID
                  },
                  pillTransformer(el) {
                    return <>el.email</>
                  },
                }),
                tagIDsMultiselect: selectOptionsBuilder({
                  type: 'multiselect',
                  values: tags,
                  optionTransformer(el) {
                    return (
                      <Group align="center">
                        <Flex align={'center'}></Flex>
                        <div style={{ marginLeft: 'auto' }}>{el?.name}</div>
                      </Group>
                    )
                  },
                  formValueTransformer(el) {
                    return el.workItemTagID
                  },
                  pillTransformer(el) {
                    return <>{el.name} label</>
                  },
                }),
              },
            }}
          />
        </FormProvider>
      </MantineProvider>,
    )

    // console.log(view.container.innerHTML)

    const ids = [
      'demoWorkItemCreateForm',
      'demoWorkItemCreateForm-tagIDs-title',
      'demoWorkItemCreateForm-tagIDs-add-button',
      'demoWorkItemCreateForm-tagIDs-0',
      'demoWorkItemCreateForm-tagIDs-remove-button-0',
      'demoWorkItemCreateForm-tagIDs-1',
      'demoWorkItemCreateForm-tagIDs-remove-button-1',
      'demoWorkItemCreateForm-tagIDs-2',
      'demoWorkItemCreateForm-tagIDs-remove-button-2',
      'demoWorkItemCreateForm-tagIDsMultiselect', // multiselects dont have titles - using vanilla
      'demoWorkItemCreateForm-members-title',
      'demoWorkItemCreateForm-members-add-button',
      'demoWorkItemCreateForm-members-remove-button-0',
      'demoWorkItemCreateForm-members.0.role',
      'demoWorkItemCreateForm-members.0.userID',
      'demoWorkItemCreateForm-members-remove-button-1',
      'demoWorkItemCreateForm-members.1.role',
      'demoWorkItemCreateForm-members.1.userID',
      'demoWorkItemCreateForm-base-title',
      'demoWorkItemCreateForm-base.closed',
      'demoWorkItemCreateForm-base.description',
      'demoWorkItemCreateForm-base.kanbanStepID',
      'demoWorkItemCreateForm-base.targetDate',
      'demoWorkItemCreateForm-base.teamID',
      'demoWorkItemCreateForm-base.items-title',
      'demoWorkItemCreateForm-base.items-add-button',
      'demoWorkItemCreateForm-base.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.0.name',
      'demoWorkItemCreateForm-base.items.0.userId-title',
      'demoWorkItemCreateForm-base.items.0.userId-add-button',
      'demoWorkItemCreateForm-base.items.0.items-title',
      'demoWorkItemCreateForm-base.items.0.items-add-button',
      'demoWorkItemCreateForm-base.items.0.items-0',
      'demoWorkItemCreateForm-base.items.0.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.0.items-1',
      'demoWorkItemCreateForm-base.items.0.items-remove-button-1',
      'demoWorkItemCreateForm-base.items-remove-button-1',
      'demoWorkItemCreateForm-base.items.1.name',
      'demoWorkItemCreateForm-base.items.1.userId-title',
      'demoWorkItemCreateForm-base.items.1.userId-add-button',
      'demoWorkItemCreateForm-base.items.1.items-title',
      'demoWorkItemCreateForm-base.items.1.items-add-button',
      'demoWorkItemCreateForm-base.items.1.items-0',
      'demoWorkItemCreateForm-base.items.1.items-remove-button-0',
      'demoWorkItemCreateForm-base.items.1.items-1',
      'demoWorkItemCreateForm-base.items.1.items-remove-button-1',
      'demoWorkItemCreateForm-base.workItemTypeID',
      'demoWorkItemCreateForm-demoProject.lastMessageAt',
      'demoWorkItemCreateForm-demoProject.line',
      'demoWorkItemCreateForm-demoProject.ref',
      'demoWorkItemCreateForm-demoProject.reopened',
      'demoWorkItemCreateForm-demoProject.workItemID',
    ]

    const actualIds = [...document.querySelectorAll('[data-testid^="demoWorkItemCreateForm"]')].map((e) =>
      e.getAttribute('data-testid'),
    )
    // console.log({ actualIds })
    expect(actualIds.sort()).toEqual(ids.sort())

    // test should submit with default values if none changed

    // FIXME: dont check state, its not updated. call submit with mock onsubmit that returns data and check
    // that return value is what we expect.
    // https://react-hook-form.com/advanced-usage#TestingForm
    // but this is still not updated in test for some reason: formisbool: true  but 'demoProject.reopened' false
    expect(form.current.getValues('members.0.role')).toEqual('preparer') // was intentionally undefined
    // expect(form.current.getValues('demoProject.reopened')).toEqual(true)

    const formElement = screen.getByTestId(formName)
    fireEvent.submit(formElement)
    console.log({ formValues: form.current.getValues() })
    console.log(form.current.formState.errors)
    console.log(form.current.formState.isValid)
    expect(form.current.formState.errors).toEqual({})

    const checkbox = screen.getByTestId('demoWorkItemCreateForm-demoProject.reopened')
    console.log({ checkbox: checkbox.innerHTML })

    // TODO: test comboboxes options get rendered
  })
})

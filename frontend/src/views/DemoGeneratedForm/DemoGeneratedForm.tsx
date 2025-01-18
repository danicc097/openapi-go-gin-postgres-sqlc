import { Title, Button, Text, Flex, Group, Container } from '@mantine/core'
import DynamicForm from 'src/utils/formGeneration'
import { Topic, type CreateWorkItemTagRequest, type ModelsWorkItemTag } from 'src/gen/model'
import { FormProvider, useForm } from 'react-hook-form'
import { ajvResolver } from '@hookform/resolvers/ajv'
import dayjs from 'dayjs'
import _ from 'lodash'
import { colorBlindPalette } from 'src/utils/colors'
import { IconCircle, IconTag } from '@tabler/icons'
import { fullFormats } from 'ajv-formats/dist/formats'
import WorkItemRoleBadge from 'src/components/Badges/WorkItemRoleBadge'
import { WORK_ITEM_ROLES } from 'src/services/authorization'
import { useEffect, useState } from 'react'
import { useCalloutErrors } from 'src/components/Callout/useCalloutErrors'
import UserComboboxOption from 'src/components/Combobox/UserComboboxOption'
import { colorSwatchComponentInputOption } from 'src/components/formGeneration/components'
import { useGetPaginatedUsers } from 'src/gen/user/user'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { useFormSlice } from 'src/slices/form'
import { WorkItemTagID, ProjectID } from 'src/gen/entity-ids'
import { selectOptionsBuilder } from 'src/utils/formGeneration.context'
import { parseSchemaFields } from 'src/utils/jsonSchema'
import { JSONSchema4 } from 'json-schema'
import { apiPath } from 'src/services/apiPaths'
import qs from 'qs'

const schema: JSONSchema4 = {
  properties: {
    base: {
      properties: {
        closed: {
          format: 'date-time',
          type: ['object', 'null'],
        },
        description: {
          type: 'string',
          minLength: 1,
        },
        kanbanStepID: {
          type: 'integer',
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
                  minLength: 1,
                },
                type: ['array', 'null'],
              },
              userId: {
                items: {
                  type: 'string',
                  minLength: 1,
                },
                type: ['array', 'null'],
              },
              name: {
                type: 'string',
                minLength: 1,
              },
            },
            required: ['userId', 'name'],
            type: 'object',
          },
          type: ['array', 'null'],
        },
        workItemTypeID: {
          type: 'integer',
        },
      },
      required: ['items', 'description', 'workItemTypeID', 'teamID', 'kanbanStepID', 'closed', 'targetDate'],
      type: 'object',
    },
    demoProject: {
      properties: {
        lastMessageAt: {
          format: 'date-time',
          type: 'object',
        },
        line: {
          type: 'string',
          minLength: 1,
        },
        ref: {
          pattern: '^[0-9]{8}$',
          type: 'string',
          minLength: 1,
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
    },
    members: {
      items: {
        properties: {
          role: {
            title: 'WorkItem role',
            type: 'string',
            minLength: 1,
            'x-generated': '-',
            enum: ['preparer', 'reviewer'],
            description: "represents a database 'work_item_role'",
          },
          userID: {
            type: 'string',
            minLength: 1,
          },
        },
        required: ['userID', 'role'],
        type: 'object',
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
}
const uuids = [
  'fcd252dc-72a4-4514-bdd1-3cac573a5fac',
  '120cb364-2b18-49fb-b505-568834614c5d',
  'bdab07d6-c2b4-44b0-b6d0-e87f62037cc1',
  'ad52daf8-9bad-4671-b3f1-535178b0346e',
  '3e82b3a5-5757-4860-8bf7-2e7962534328',
  'd59d3a5c-b99f-40aa-9419-75a2bbb0fd52',
]

const tags = [...Array(1000)].map((x, i) => {
  const tag: ModelsWorkItemTag = {
    name: `tag #${i}`,
    color: _.sample(colorBlindPalette)!,
    workItemTagID: i as WorkItemTagID,
    projectID: 1 as ProjectID,
    description: `description for tag #${i}`,
  } // TODO: get workitem tags endpoint
  return tag
})

export default function DemoGeneratedForm() {
  const formInitialValues: TestTypes.DemoWorkItemCreateRequest = {
    base: {
      items: [
        {
          items: ['0001', '0002'],
          // TODO: use usersData below but leave some that dont exist.
          // TODO: if select or multiselect not found, it should show a warning callout
          // stating option was not found so its being ignored (persistent callout)
          userId: ['120cb364-2b18-49fb-b505-568834614c5d', 'fcd252dc-72a4-4514-bdd1-3cac573a5fac'],
          name: 'item-1',
        },
        { items: ['0011', '0012'], userId: ['baduserid'], name: 'item-2' },
      ],
      closed: dayjs('2023-03-24T20:42:00.000Z').toDate(),
      targetDate: dayjs('2023-02-22').toDate(),
      description: 'some text',
      kanbanStepID: 1,
      teamID: 1,
      metadata: {},
      // title: {},
      workItemTypeID: 1,
    },
    // TODO: formGeneration must not assume options do exist, else all fails catastrophically.
    // it's not just checking types...
    // 1. move callout errors state to zustand, and create callout warnings too
    // 2.(sol 1) if option not found for initial data, remove from form values
    // and show persistent callout _warning_ that X was deleted since it was not found.
    // it should update the form but show callout error saying ignoring bad type in `formField`, in this case `tagIDs.1`
    // 2. (sol 2 which wont work) leave form as is and validate on first render will not catch errors for options not found, if type is right...
    tagIDs: null,
    tagIDsMultiselect: [
      1,
      2,
      'badid' as any,
      'badid2' as any /** FIXME: show warning callout, it should not create it at all */,
    ],
    // tagIDs: [0, 5, 8],
    demoProject: {
      line: '',
      ref: '12341234',
      lastMessageAt: dayjs('2023-03-24T20:42:00.000Z').toDate(),
      // ref: '12341234', // will set defaultValue if unset
      workItemID: 1,
      reopened: true, // TODO: test it does work (requires no defaultValue being set on checkbox component)
    },
    members: [
      { userID: '81c662f2-c014-4681-a75a-4a56a0d87a93', role: null as any /** test defaultValues */ },
      { role: 'preparer', userID: 'bad userID' },
    ],
  }

  const { user } = useAuthenticatedUser()

  const [cursor, setCursor] = useState(new Date().toISOString())

  useEffect(() => {
    const sse = new EventSource(
      `${apiPath('/events')}?${qs.stringify(
        { projectName: 'demo', topics: [Topic.GlobalAlerts, Topic.TeamCreated, Topic.AppDebug] },
        { arrayFormat: 'repeat' },
      )}`,
      {
        withCredentials: true,
      },
    )
    function getRealtimeData(data) {
      console.log({ dataSSE: data })
    }
    sse.onmessage = (e) => getRealtimeData(JSON.parse(e.data))
    sse.onerror = (e) => {
      console.log({ errorSSE: e })
      sse.close()
    }
    return () => {
      sse.close()
    }
  }, [])

  // useStopInfiniteRenders(20)

  // watch out for queryKey slugs having dynamic values (like new Date() or anything generated)
  const { data: usersData } = useGetPaginatedUsers({ direction: 'desc', cursor, limit: 0, column: 'createdAt' })

  const form = useForm<TestTypes.DemoWorkItemCreateRequest>({
    resolver: ajvResolver(schema as any, {
      strict: false,
      formats: fullFormats,
    }),
    mode: 'all',
    reValidateMode: 'onChange',
    defaultValues: formInitialValues ?? {},
    // shouldUnregister: true, // defaultValues will not be merged against submission result.
  })
  const createWorkItemTagForm = useForm<CreateWorkItemTagRequest>({
    resolver: ajvResolver(schema as any, {
      strict: false,
      formats: fullFormats,
    }),
    mode: 'all',
    reValidateMode: 'onChange',
  })
  const { register, handleSubmit, control, formState } = form
  const errors = formState.errors
  const formSLice = useFormSlice()
  const [errorSet, seterrorSet] = useState(false)
  const { extractCalloutErrors, setCalloutErrors, calloutErrors, extractCalloutTitle } =
    useCalloutErrors('demoWorkItemCreateForm')

  // useEffect(() => {
  //   console.log('errors')
  //   console.log(errors)
  //   // if (Object.keys(errors).length > 0 && !errorSet) {
  //   // setCalloutErrors('Validation error')

  //   // console.log('errors')
  //   // console.log(errors)

  //   // setCalloutErrors('Validation error')
  //   // seterrorSet(true)
  //   // // console.log(formSLice.callout[formName])
  //   // // console.log(`form has errors: ${Object.keys(errors).length > 0}`)
  //   // }
  // }, [formState])

  // useEffect(() => {
  //   console.log(demoWorkItemCreateForm.values)
  //   try {
  //     DemoWorkItemCreateRequestDecoder.decode(demoWorkItemCreateForm.values)
  //   } catch (error) {
  //     console.error(JSON.stringify(error.validationErrors.errors))
  //   }
  // }, [demoWorkItemCreateForm])

  type ExcludedFormKeys = 'base.metadata' | 'tagIDs' | 'demoProject' | 'base'

  const users = usersData?.items

  const userIdSelectOption = selectOptionsBuilder({
    type: 'select',
    values: users ?? [],
    //  TODO: transformers can be reusable between forms. could simply become
    //  {
    //   type: "select"
    //   values: ...
    //   ...userIdFormTransformers
    // }
    optionTransformer(el) {
      return <UserComboboxOption user={el} />
    },
    formValueTransformer(el) {
      return el.userID
    },
    pillTransformer(el) {
      return <>{el.email}</>
    },
    searchValueTransformer(el) {
      return `${el.email} ${el.fullName} ${el.username}`
    },
  })

  return (
    <Container maw={600}>
      {/* <LandingPage /> */}

      <Title size={20}>This form has been automatically generated from an openapi spec</Title>
      <Button
        onClick={(e) => {
          try {
            console.log(errors?.demoProject)
            // const r = demoWorkItemCreateForm.validate()
            // console.log({ r })
            // DemoWorkItemCreateRequestDecoder.decode(demoWorkItemCreateForm.values)
          } catch (error) {
            console.error(JSON.stringify(error?.validationErrors?.errors))
          }
        }}
      >
        Validate form
      </Button>
      <FormProvider {...form}>
        <DynamicForm<TestTypes.DemoWorkItemCreateRequest, ExcludedFormKeys>
          onSubmit={(e) => {
            e.preventDefault()
            form.handleSubmit(
              (data) => {
                console.log({ data })
              },
              (errors) => {
                console.log({ errors })
              },
            )(e)
          }}
          formName="demoWorkItemCreateForm"
          // schemaFields will come from `parseSchemaFields(schema.RestDemo... OR  asConst(jsonSchema.definitions.<...>))`
          // using this hardcoded for testing purposes
          schemaFields={parseSchemaFields(schema)}
          options={{
            // labels are mandatory. Use null to exclude if needed.
            labels: {
              'base.closed': 'Closed',
              'base.description': 'Description',
              // 'base.metadata': 'metadata', // in excluded form keys
              'base.kanbanStepID': 'Kanban step', // if using KanbanStep transformer, then "Kanban step", "Kanban step name", etc.
              'base.targetDate': 'Target date',
              'demoProject.reopened': 'Reopened',
              'base.teamID': 'Team',
              'base.items': 'Items',
              'base.items.name': 'Name',
              'base.items.items': 'Items',
              'base.items.userId': 'User',
              'base.workItemTypeID': 'Type',
              'demoProject.lastMessageAt': 'Last message at',
              'demoProject.line': 'Line',
              'demoProject.ref': 'Ref',
              'demoProject.workItemID': 'Work item',
              members: 'Members',
              'members.role': 'Role',
              'members.userID': 'User',
              tagIDsMultiselect: 'Tags',
            },
            /** TODO: array of arrays|string to allow horizontal grouping instead renderOrderPriority */
            // no need to ensure all fields are present
            // renderLayout: [['demoProject.ref', 'demoProject.line'], 'members', ...],
            renderOrderPriority: ['tagIDsMultiselect', 'members'],
            accordion: {
              'base.items': {
                defaultOpen: true,
                title: formAccordionTitle('Items'),
              },
            },
            // TODO: should have `warnings` options funcs that receive the element
            // and returns a string[] of warnings.
            // can be used for adhoc warnings, e.g. this value may be too high, or
            // this user hasn't logged in >n months, this date is before today's date, etc.
            // TODO: these should be default values for nested array fields on creation, != formDefaultValues
            defaultValues: {
              'demoProject.ref': '11112222',
              'members.role': 'preparer',
            },
            fieldOptions: {
              'base.closed': {
                warningFn(el) {
                  return dayjs(el) > dayjs('01-01-2023') ? ['Date is higher than 01-01-2023'] : []
                },
              },
              'base.items': {
                warningFn(el) {
                  const warnings: string[] = []
                  if (el.name !== 'item-1') warnings.push('Item name is not "item-1"')

                  // FIXME: must initialize when adding new items via
                  // demoWorkItemCreateForm-base.items-add-button
                  // for array of objects in which the object has an array key
                  if (el.items?.includes('0001')) warnings.push('Nested items include "0001"')

                  return warnings
                },
              },
            },
            selectOptions: {
              'members.userID': userIdSelectOption,
              'base.items.userId': userIdSelectOption,
              tagIDsMultiselect: selectOptionsBuilder({
                type: 'multiselect',
                searchValueTransformer(el) {
                  return el.name
                },
                values: tags,
                optionTransformer(el) {
                  return (
                    <Group align="center">
                      <Flex align="center" gap={12} justify="center">
                        <IconCircle size={12} fill={el.color} />
                        <div>{el.name}</div>
                      </Flex>
                    </Group>
                  )
                },
                formValueTransformer: (el) => el.workItemTagID,
                pillTransformer: (el) => <div>{el.name}</div>,
                labelColor: (el) => el.color,
              }),
              'members.role': selectOptionsBuilder({
                type: 'select',
                values: WORK_ITEM_ROLES,
                optionTransformer: (el) => <WorkItemRoleBadge role={el} />,
                formValueTransformer: (el) => el,
                pillTransformer: (el) => <WorkItemRoleBadge role={el} />,
              }),
            },
            input: {
              'demoProject.line': {
                component: colorSwatchComponentInputOption,
              },
            },
            propsOverride: {
              'base.workItemTypeID': {
                description: 'This is some help text for a disabled field.',
                disabled: true,
              },
            },
          }} // satisfies DynamicFormOptions<TestTypes.DemoWorkItemCreateRequest, ExcludedFormKeys> // not needed anymore for some reason
        />
      </FormProvider>
    </Container>
  )
}
function formAccordionTitle(title: string): JSX.Element {
  return (
    <Flex align="center" gap={10}>
      <IconTag size={16} />
      <Text fw={700} size={'md'}>
        {title}
      </Text>
    </Flex>
  )
}

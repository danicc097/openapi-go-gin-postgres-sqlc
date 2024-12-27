import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import 'src/assets/css/pulsate.css'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
import { Button, Card, Popover, Space, Text, Textarea } from '@mantine/core'
import DynamicForm, { type SelectOptions, type DynamicFormOptions, InputOptions } from 'src/utils/formGeneration'
import type { CreateWorkItemTagRequest, WorkItemRole } from 'src/gen/model'
import type { GetKeys, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import { FormProvider, useForm, useFormState, useWatch } from 'react-hook-form'
import { ajvResolver } from '@hookform/resolvers/ajv'
import { fullFormats } from 'ajv-formats/dist/formats'
import { parseSchemaFields } from 'src/utils/jsonSchema'
import { colorSwatchComponentInputOption } from 'src/components/formGeneration/components'
import { CodeHighlight } from '@mantine/code-highlight'
import { useFormSlice } from 'src/slices/form'
import { entries } from 'src/utils/object'
import { JSONSchemaType } from 'ajv'
import { JSON_SCHEMA, OPERATION_AUTH } from 'src/config'
import { Tooltip as ReactTooltip } from 'react-tooltip'
import { AppTourProvider } from 'src/tours/AppTourProvider'
import { useTour } from '@reactour/tour'
import { useCreateWorkItemTag } from 'src/gen/work-item-tag/work-item-tag'
import { IconCheck, IconCross, IconX } from '@tabler/icons'
import { notifications } from '@mantine/notifications'
import { useUISlice } from 'src/slices/ui'

export default function CreateWorkItemTagForm() {
  const formSlice = useFormSlice()
  const uiSlice = useUISlice()

  const createWorkItemTagRequestSchema = JSON_SCHEMA.definitions.CreateWorkItemTagRequest
  const createWorkItemTagForm = useForm<CreateWorkItemTagRequest>({
    resolver: ajvResolver(createWorkItemTagRequestSchema as any, {
      strict: false,
      formats: fullFormats,
    }),
    mode: 'all',
    reValidateMode: 'onChange',
    defaultValues: {
      color: '#123123',
      description: 'a',
      name: 'c',
    },
  })
  const formName = 'createWorkItemTagForm'
  const createWorkItemTag = useCreateWorkItemTag()
  const authorization = OPERATION_AUTH.CreateWorkItemTag

  const tour = useTour()

  return (
    <>
      <Button
        className="tour-button"
        onClick={(e) => {
          tour.setIsOpen(true)
        }}
      >
        Open tour
      </Button>
      <Space p={10} />
      <Button className="tour-button-example">Click to continue tour</Button>
      <h3>Authorization:</h3>
      <CodeHighlight code={JSON.stringify(authorization, null, '  ')} language="json" />
      <h3>Form:</h3>
      <FormProvider {...createWorkItemTagForm}>
        <DynamicForm<CreateWorkItemTagRequest>
          onSubmit={(e) => {
            e.preventDefault()
            createWorkItemTagForm.handleSubmit(
              (data) => {
                console.log({ data })
                createWorkItemTag.mutate(
                  { data, projectName: uiSlice.project },
                  {
                    onSuccess(data, variables, context) {
                      formSlice.setCalloutErrors(formName, [])
                      console.log({ onSuccess: data })
                      notifications.show({
                        title: 'Success',
                        color: 'green',
                        icon: <IconCheck size="1.2rem" />,
                        autoClose: 5000,
                        message: 'Item created successfully',
                      })
                      createWorkItemTagForm.reset()
                    },
                    onError(error, variables, context) {
                      if (!error.response) return

                      // TODO: callouterror accepts httperror and will handle bare httperror and
                      // httperror with array of validationerror
                      console.log({ onError: error.response?.data })
                      formSlice.setCalloutErrors(formName, [error.response.data.detail])
                      notifications.show({
                        title: error.response.data.title,
                        color: 'red',
                        icon: <IconX size="1.2rem" />,
                        autoClose: 10000,
                        message: error.response.data.detail,
                      })
                    },
                  },
                )
              },
              (errors) => {
                console.log({ errors })
              },
            )(e)
          }}
          formName={formName}
          schemaFields={parseSchemaFields(createWorkItemTagRequestSchema)}
          options={{
            labels: {
              color: 'Color',
              description: 'Description',
              name: 'Name',
            },
            propsOverride: {
              description: {
                'data-tour': 'description-input',
              },
            },
            input: {
              description: {
                component: <Textarea resize="vertical" styles={{ root: { width: '100%' } }} />,
              },
              color: {
                component: colorSwatchComponentInputOption,
              },
            },
          }}
        ></DynamicForm>
      </FormProvider>
    </>
  )
}

import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import 'src/assets/css/pulsate.css'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
import { Button, Card, Popover, Space, Text, Textarea } from '@mantine/core'
import DynamicForm, { type SelectOptions, type DynamicFormOptions, InputOptions } from 'src/utils/formGeneration'
import type { CreateWorkItemTagRequest, DbWorkItemTag, User, WorkItemRole } from 'src/gen/model'
import type { GetKeys, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import { validateField } from 'src/utils/validation'
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

export default function Project() {
  const formSlice = useFormSlice()

  const createWorkItemTagRequestSchema = JSON_SCHEMA.definitions.CreateWorkItemTagRequest
  const createWorkItemTagForm = useForm<CreateWorkItemTagRequest>({
    resolver: ajvResolver(createWorkItemTagRequestSchema as any, {
      strict: false,
      formats: fullFormats,
    }),
    mode: 'all',
    reValidateMode: 'onChange',
  })
  const formName = 'createWorkItemTagForm'
  const { register, handleSubmit, control, formState } = createWorkItemTagForm
  const errors = formState.errors
  const createWIT = useCreateWorkItemTag()
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
                createWIT.mutate(
                  { data, projectName: 'demo' },
                  {
                    onSuccess(data, variables, context) {
                      console.log({ onSuccess: data })
                    },
                    onError(error, variables, context) {
                      if (!error) return

                      // FIXME: backend appends extraneous json to response
                      // {"detail":"commit unexpectedly resulted in rollback","error":"could not commit transaction (rollback error: tx is closed): commit unexpectedly resulted in rollback","loc":null,"status":500,"title":"Database error","type":"Private"}
                      console.log({ onError: error.response?.data })
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

import React, { useEffect, useState } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import 'src/assets/css/fonts.css'
import 'src/assets/css/overrides.css'
import 'src/assets/css/pulsate.css'
import FallbackLoading from 'src/components/Loading/FallbackLoading'
import { Textarea } from '@mantine/core'
import DynamicForm, {
  selectOptionsBuilder,
  type SelectOptions,
  type DynamicFormOptions,
  InputOptions,
} from 'src/utils/formGeneration'
import type { CreateWorkItemTagRequest, DbWorkItemTag, User, WorkItemRole } from 'src/gen/model'
import type { GetKeys, RecursiveKeyOfArray, PathType } from 'src/types/utils'
import { validateField } from 'src/utils/validation'
import { FormProvider, useForm, useFormState, useWatch } from 'react-hook-form'
import { ajvResolver } from '@hookform/resolvers/ajv'
import JSON_SCHEMA from 'src/client-validator/gen/dereferenced-schema.json'
import { fullFormats } from 'ajv-formats/dist/formats'
import { parseSchemaFields } from 'src/utils/jsonSchema'
import { colorSwatchComponentInputOption } from 'src/components/formGeneration/components'

export default function LandingPage() {
  const createWorkItemTagRequestSchema = JSON_SCHEMA.definitions.CreateWorkItemTagRequest
  const createWorkItemTagForm = useForm<CreateWorkItemTagRequest>({
    resolver: ajvResolver(createWorkItemTagRequestSchema as any, {
      strict: false,
      formats: fullFormats,
    }),
    mode: 'all',
    reValidateMode: 'onChange',
  })
  const { register, handleSubmit, control, formState } = createWorkItemTagForm
  const errors = formState.errors

  return (
    <FormProvider {...createWorkItemTagForm}>
      <DynamicForm<CreateWorkItemTagRequest>
        onSubmit={(e) => {
          e.preventDefault()
          createWorkItemTagForm.handleSubmit(
            (data) => {
              console.log({ data })
            },
            (errors) => {
              console.log({ errors })
            },
          )(e)
        }}
        formName="createWorkItemTagForm"
        schemaFields={parseSchemaFields(createWorkItemTagRequestSchema as any)}
        options={{
          labels: {
            color: 'Color',
            description: 'Description',
            name: 'Name',
          },

          input: {
            description: {
              // FIXME: Allow Enter when focusing on component.
              component: <Textarea styles={{ root: { width: '100%' } }} />,
            },
            color: {
              component: colorSwatchComponentInputOption,
            },
          },
        }}
      ></DynamicForm>
    </FormProvider>
  )
}

import {
  EuiBadge,
  EuiBasicTable,
  EuiButton,
  EuiCallOut,
  EuiCheckbox,
  EuiCodeBlock,
  EuiConfirmModal,
  EuiFieldText,
  EuiFlexGroup,
  EuiFlexItem,
  EuiForm,
  EuiFormRow,
  EuiHeader,
  EuiHealth,
  EuiIcon,
  EuiLink,
  EuiSelectable,
  EuiSelectableOption,
  EuiSpacer,
  EuiSuperSelect,
  EuiSuperSelectOption,
  EuiSuperSelectProps,
  EuiText,
  EuiTextColor,
  EuiTitle,
  formatDate,
  htmlIdGenerator,
} from '@elastic/eui'
import _, { capitalize, random } from 'lodash'
import React, { Fragment, useEffect, useReducer, useState } from 'react'
import type { ProjectConfig, Scope, Scopes, UpdateUserAuthRequest, UserResponse } from 'src/gen/model'
import { roleColor } from 'src/utils/colors'
import { joinWithAnd } from 'src/utils/format'
import scopes from '@scopes'
import roles from '@roles'
import type { Role } from 'src/client-validator/gen/models'
import PageTemplate from 'src/components/PageTemplate/PageTemplate'
import type { ValidationErrors } from 'src/client-validator/validate'
import { useUpdateUserAuthorization } from 'src/gen/user/user'
import { useForm, UseFormReturnType } from '@mantine/form'
import { useUISlice } from 'src/slices/ui'

const makeId = htmlIdGenerator('')
/**
 * alternative: once config comes in create a `Map` from array with key `path`
 * so that we dont have to filter (though performance doesnt matter, this page will barely be used)
 */
const boardConfig: ProjectConfig = {
  header: ['demoProject.ref', 'workItemType'],
  fields: [
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'workItemType',
      name: 'Type',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem',
      name: 'Demo project',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.metadata',
      name: 'Metadata',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.metadata.externalLink',
      name: 'External link',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.ref',
      name: 'Reference',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.line',
      name: 'Line number',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.KPIs',
      name: 'KPIs',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.KPIs.name',
      name: 'Name',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.KPIs.complexity',
      name: 'Complexity',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.KPIs.tags',
      name: 'Tags',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoProjectWorkItem.tags',
      name: 'Tags',
    },
  ],
}

export default function ProjectManagementPage() {
  const { addToast, dismissToast, theme } = useUISlice()

  const [calloutErrors, setCalloutError] = useState<ValidationErrors>(null)

  const form = useForm<any>({
    initialValues: boardConfig,
  })

  const getErrors = () =>
    calloutErrors ? calloutErrors?.errors?.map((v, i) => `${v.invalidParams.name}: ${v.invalidParams.reason}`) : null

  function onChange(e: any, key: string, path: string) {
    const idx = form.values['fields'].findIndex((f) => f.path === path)
    const newFieldsConfig = _.cloneDeep(form.values['fields'])
    const newField = newFieldsConfig[idx]
    if (typeof newField[key] === 'string') {
      newField[key] = e.target.value
    } else if (typeof newField[key] === 'boolean') {
      newField[key] = !newField[key]
    }
    newFieldsConfig[idx] = newField
    form.setFieldValue('fields', newFieldsConfig)
    console.log(newFieldsConfig)
  }

  const renderTextInput = (path: string) => {
    return (item, col) => {
      return (
        <EuiFormRow fullWidth>
          <EuiFieldText compressed defaultValue={item} onChange={(e) => onChange(e, path, col.path)} />
        </EuiFormRow>
      )
    }
  }

  const renderText = (item, col) => {
    return <p>{item}</p>
  }

  const renderCheckbox = (path: string) => {
    return (item, col) => {
      return (
        <EuiFormRow fullWidth>
          <EuiCheckbox
            compressed
            id={`checkbox-id-${col.path}-${makeId()}`}
            checked={form.values['fields'].find((f) => f.path === col.path)?.[path]}
            onChange={(e) => onChange(e, path, col.path)}
          />
        </EuiFormRow>
      )
    }
  }

  const columns = []

  ;[boardConfig['fields'][0]].forEach((field) => {
    Object.entries(field).forEach(([k, v]) => {
      const column = {}
      column['field'] = k
      column['name'] = k
      column['sortable'] = true
      column['truncateText'] = true
      if (k === 'id') {
      } else if (k === 'path') {
        column['render'] = renderText
        column['truncateText'] = false
        column['width'] = '40vw'
      } else if (typeof v === 'boolean') {
        column['render'] = renderCheckbox(k)
      } else if (typeof v === 'string') {
        column['render'] = renderTextInput(k)
      }
      column['render'] && columns.push(column)
    })
  })
  console.log('columns')
  console.log(columns)

  const getRowProps = (item) => {
    const { id } = item
    return {
      'data-test-subj': `row-${id}`,
      className: 'customRowClass',
      onClick: () => {
        null
      },
    }
  }

  const getCellProps = (item, column) => {
    const { id } = item
    const { field } = column
    return {
      className: 'customCellClass',
      'data-test-subj': `cell-${id}-${field}`,
      textOnly: true,
    }
  }

  const title = (
    <div>
      <EuiFlexGroup gutterSize="s" alignItems="center" responsive={false}>
        <EuiFlexItem grow={false}>
          <EuiIcon type="eraser" size="m" />
        </EuiFlexItem>

        <EuiFlexItem>
          <EuiTitle size="xs">
            <h3 style={{ color: 'dodgerblue' }}>Update project configuration</h3>
          </EuiTitle>
        </EuiFlexItem>
      </EuiFlexGroup>

      <EuiText size="s">
        <p>
          <EuiTextColor color="subdued">{_.unescape(`Update card visualization, etc.`)}</EuiTextColor>
        </p>
      </EuiText>
    </div>
  )

  const generateData = () => {
    return boardConfig.fields.map((f: any) => {
      f['id'] = f.path
      return f
    })
  }

  const element = (
    <>
      {getErrors()}
      {/* <EuiSpacer></EuiSpacer>
      <EuiTitle size="xs">
        <EuiText>Form</EuiText>
      </EuiTitle>
      <EuiCodeBlock language="json">{JSON.stringify(form, null, 4)}</EuiCodeBlock>
      <EuiSpacer></EuiSpacer> */}
      <EuiForm
        component="form"
        // onSubmit={form.onSubmit(onBoardConfigUpdateSubmit, handleError)}
        isInvalid={Boolean(form.errors.length)}
        error={getErrors()}
      >
        <EuiBasicTable
          tableCaption="Demo of EuiBasicTable"
          items={generateData()}
          rowHeader="firstName"
          columns={columns}
          rowProps={getRowProps}
          cellProps={getCellProps}
        />{' '}
      </EuiForm>
    </>
  )

  return (
    <PageTemplate header={{ children: title }} content={element} restrictWidth={'100vw'} buttons={[]} offset={100} />
  )
}

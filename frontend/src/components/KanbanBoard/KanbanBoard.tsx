import React, { useEffect, useState } from 'react'

import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import _, { random, uniqueId } from 'lodash'
import type { ProjectConfig } from 'src/gen/model'
import moment from 'moment'
// import { getGetProjectWorkitemsMock, getProjectMSW } from 'src/gen/project/project.msw'
import ProtectedComponent from 'src/components/Permissions/ProtectedComponent'
import useAuthenticatedUser from 'src/hooks/auth/useAuthenticatedUser'
import { generateColor } from 'src/utils/colors'
import { css } from '@emotion/css'
import type { ArrayElement } from 'src/types/utils'
import { isValidURL } from 'src/utils/urls'
import { removePrefix } from 'src/utils/strings'
import { Text } from '@mantine/core'

const exampleDemoWorkItem = {
  workItemType: 'type 1',
  demoWorkItem: {
    ref: 'ABCD-ABCD',
    line: 123,
    KPIs: [
      {
        complexity: 'kpi complexity 1',
        name: 'kpi name 1',
        tags: ['tag 1', 'tag 5'],
      },
      {
        complexity: 'kpi complexity 2',
        name: 'kpi name 2',
        tags: ['tag 1'],
      },
    ],
    tags: ['critical', 'external client'],
    metadata: {
      externalLink: 'https://externallink',
      count: 123456,
    },
  },
}

const boardConfig: ProjectConfig = {
  header: ['demoProject.ref', 'workItemType'],
  fields: [
    {
      isEditable: true,
      showCollapsed: false,
      isVisible: true,
      path: 'workItemType',
      name: 'Type',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoWorkItem',
      name: 'Demo project',
    },
    {
      isEditable: true,
      showCollapsed: false,
      isVisible: true,
      path: 'demoWorkItem.metadata',
      name: 'Metadata',
    },
    {
      isEditable: true,
      showCollapsed: false,
      isVisible: true,
      path: 'demoWorkItem.metadata.externalLink',
      name: 'External link',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoWorkItem.ref',
      name: 'Reference',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoWorkItem.line',
      name: 'Line number',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoWorkItem.KPIs',
      name: 'KPIs',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoWorkItem.KPIs.name',
      name: 'Name',
    },
    {
      isEditable: true,
      showCollapsed: false,
      isVisible: true,
      path: 'demoWorkItem.KPIs.complexity',
      name: 'Complexity',
    },
    {
      isEditable: true,
      showCollapsed: false,
      isVisible: true,
      path: 'demoWorkItem.KPIs.tags',
      name: 'Tags',
    },
    {
      isEditable: true,
      showCollapsed: true,
      isVisible: true,
      path: 'demoWorkItem.tags',
      name: 'Tags',
    },
  ],
}

export default function KanbanBoard() {
  return (
    <div>
      <Text>Home</Text>
    </div>
  )
}
function getNestedObjects(nestedObjects: ProjectConfig['fields'], field: ArrayElement<ProjectConfig['fields']>) {
  return nestedObjects.filter((f) => {
    const fieldCount = _.countBy(f.path)['.'] || 0
    const parentFieldCount = _.countBy(field.path)['.'] || 0
    return fieldCount > parentFieldCount
  })
}

import React, { useState } from 'react'
import {
  EuiDragDropContext,
  EuiDraggable,
  EuiDroppable,
  EuiButtonIcon,
  EuiPanel,
  euiDragDropMove,
  euiDragDropReorder,
  htmlIdGenerator,
  DragDropContextProps,
  EuiTitle,
  EuiText,
  EuiCard,
  EuiCode,
  EuiCodeBlock,
  EuiButton,
  EuiSpacer,
  EuiBadge,
  EuiCheckbox,
  EuiFlexGroup,
  EuiFlexItem,
  EuiLink,
  EuiHorizontalRule,
} from '@elastic/eui'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import _, { isArray, isObject, random, uniqueId } from 'lodash'
import type { DemoProjectWorkItemsResponse } from 'src/gen/model'
import moment from 'moment'
import { getGetProjectWorkitemsMock, getProjectMSW } from 'src/gen/project/project.msw'
import { StyledEuiCheckbox } from 'src/components/KanbanBoard/KanbanBoard.styles'
import ProtectedComponent from 'src/components/Permissions/ProtectedComponent'
import { useAuthenticatedUser } from 'src/hooks/auth/useAuthenticatedUser'
import { generateColor } from 'src/utils/colors'
import { css } from '@emotion/css'
import type { NestedPaths } from 'src/types/utils'
import { isValidURL } from 'src/utils/urls'

const makeId = htmlIdGenerator()

const exampleDemoProjectWorkItem = {
  workItemType: 'type 1',
  demoProjectWorkItem: {
    ref: 'ABCD-ABCD',
    line: 123,
    KPIs: [
      {
        complexity: 'kpi complexity 1',
        name: 'kpi name 1',
      },
      {
        complexity: 'kpi complexity 2',
        name: 'kpi name 2',
      },
    ],
    metadata: {
      externalLink: 'https://externallink',
      count: 123456,
    },
  },
}

const boardConfig = {
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
  ],
}

export default function KanbanBoard() {
  const [isModalVisible, setIsModalVisible] = useState(false)

  const showModal = () => setIsModalVisible(true)

  const hideModal = () => setIsModalVisible(false)

  const { user } = useAuthenticatedUser()
  const { addToast } = useUISlice()

  boardConfig.fields = _.sortBy(boardConfig.fields, 'path')

  const renderCard = (data: any) => {
    const skipFields = []

    return (
      <>
        {boardConfig.fields.map((field, i) => {
          if (skipFields.includes(field.path)) return
          if (_.get(data, field.path) === undefined) return

          const value = _.get(data, field.path)

          let element

          // TODO should accumulate elements in elements and panel elements here as well

          if (typeof value === 'object') {
            const nestedFields = boardConfig.fields.filter((f) => f.path.startsWith(field.path))
            const fieldNestedObjects = getNestedObjects(nestedFields, field)

            element = createCardPanel(nestedFields, fieldNestedObjects, skipFields, data, field)
          } else {
            element = createCardField(value, field)
          }

          return element
        })}
      </>
    )
  }

  const makeList = (number, start = 1) =>
    Array.from({ length: number }, (v, k) => k + start).map((el) => {
      return {
        content: (
          <EuiCard
            textAlign="left"
            title={
              <>
                <EuiFlexGroup direction="row" justifyContent="spaceBetween">
                  <EuiFlexItem>Card {el}</EuiFlexItem>
                  <EuiButtonIcon
                    iconType="documentEdit"
                    aria-label="Heart"
                    color="primary"
                    onClick={() => {
                      // TODO Navigate /workitem/:id
                      null
                    }}
                  />
                </EuiFlexGroup>
                <EuiHorizontalRule size="full" margin="xs"></EuiHorizontalRule>
              </>
            }
            // description={}
            titleElement="h3"
            hasBorder={false}
            paddingSize="none"
            display="plain"
            // footer={<></>}
          >
            {renderCard(exampleDemoProjectWorkItem)}
            <EuiSpacer />
            <EuiButton
              key={1}
              size="s"
              onClick={() =>
                addToast({
                  id: ToastId.AuthzError + uniqueId(),
                  title: 'clicked',
                  color: 'success',
                  iconType: 'alert',
                  toastLifeTimeMs: 15000,
                  text: 'clicked.',
                })
              }
            >
              New toast
            </EuiButton>
          </EuiCard>
        ),
        id: makeId(),
      }
    })

  const [list, setList] = useState([1, 2, 3, 4])
  const [list1, setList1] = useState(makeList(3))
  const [list2, setList2] = useState(makeList(3, 4))
  const [list3, setList3] = useState(makeList(1, 7))
  const [list4, setList4] = useState(makeList(1, 8))
  const lists = {
    COMPLEX_DROPPABLE_PARENT: list,
    COMPLEX_DROPPABLE_AREA_1: list1,
    COMPLEX_DROPPABLE_AREA_2: list2,
    COMPLEX_DROPPABLE_AREA_3: list3,
    COMPLEX_DROPPABLE_AREA_4: list4,
  }
  const actions = {
    COMPLEX_DROPPABLE_PARENT: setList,
    COMPLEX_DROPPABLE_AREA_1: setList1,
    COMPLEX_DROPPABLE_AREA_2: setList2,
    COMPLEX_DROPPABLE_AREA_3: setList3,
    COMPLEX_DROPPABLE_AREA_4: setList4,
  }
  const onDragEnd: DragDropContextProps['onDragEnd'] = ({ source, destination }) => {
    if (source && destination) {
      if (source.droppableId === destination.droppableId) {
        const items = euiDragDropReorder(lists[destination.droppableId], source.index, destination.index)

        actions[destination.droppableId](items)
      } else {
        const sourceId = source.droppableId
        const destinationId = destination.droppableId
        const result = euiDragDropMove(lists[sourceId], lists[destinationId], source, destination)

        actions[sourceId](result[sourceId])
        actions[destinationId](result[destinationId])
      }
    }
  }

  /**
   * TODO:
   *
   * - ignore card moved in own column
   * - order by target date desc
   * - render children dynamically:
   *    - date/string/number: text
   *    - bool: checkbox
   *    - array: EuiBadge or comma delim
   */

  return (
    <>
      <ProtectedComponent>
        <>
          {' '}
          <EuiButton>test</EuiButton>
        </>
      </ProtectedComponent>
      <ProtectedComponent requiredRole="superAdmin">
        <>
          {' '}
          <EuiButton>test superadmin</EuiButton>
        </>
      </ProtectedComponent>
      <EuiDragDropContext onDragEnd={onDragEnd}>
        <EuiDroppable
          droppableId="COMPLEX_DROPPABLE_PARENT"
          type="MACRO"
          direction="horizontal"
          withPanel
          spacing="l"
          className="eui-scrollBar"
          style={{ display: 'flex', maxWidth: '100vw', overflowX: 'auto' }}
        >
          {list.map((did, didx) => (
            <EuiDraggable
              key={did}
              index={didx}
              draggableId={`COMPLEX_DRAGGABLE_${did}`}
              spacing="l"
              style={{ minWidth: '25vw' }}
              disableInteractiveElementBlocking // Allows button to be drag handle
              hasInteractiveChildren
              isDragDisabled
              customDragHandle
            >
              {(provided) => (
                <>
                  <EuiPanel color="subdued" paddingSize="s">
                    <EuiTitle size="xs">
                      <EuiText textAlign="center">Column {didx}</EuiText>
                    </EuiTitle>
                    <EuiDroppable
                      droppableId={`COMPLEX_DROPPABLE_AREA_${did}`}
                      type="MICRO"
                      spacing="m"
                      style={{ flex: '1 0 50%' }}
                    >
                      {lists[`COMPLEX_DROPPABLE_AREA_${did}`].map(({ content, id }, idx) => (
                        <EuiDraggable key={id} index={idx} draggableId={id} spacing="m">
                          <EuiPanel>{content}</EuiPanel>
                        </EuiDraggable>
                      ))}
                    </EuiDroppable>
                  </EuiPanel>
                </>
              )}
            </EuiDraggable>
          ))}
        </EuiDroppable>
      </EuiDragDropContext>
    </>
  )

  function createCardPanel(
    fields: { isEditable: boolean; showCollapsed: boolean; isVisible: boolean; path: string; name: string }[],
    nestedObjects: { isEditable: boolean; showCollapsed: boolean; isVisible: boolean; path: string; name: string }[],
    skipFields: any[],
    data: any,
    currentField: { isEditable: boolean; showCollapsed: boolean; isVisible: boolean; path: string; name: string },
  ) {
    let element
    let elements = []
    const panelElements = []

    for (const field of fields) {
      if (skipFields.includes(field.path)) continue
      skipFields.push(field.path)
      if (nestedObjects.includes(field) && typeof _.get(data, field.path) === 'object') {
        const nestedFields = fields.filter((f) => f.path.startsWith(field.path))
        const fieldNestedObjects = getNestedObjects(nestedFields, field)
        // TODO render panels always last and add spacing inbetween
        const el = createCardPanel(nestedFields, fieldNestedObjects, skipFields, data, field)
        el && panelElements.push(el)
        continue
      }
      const el = createCardField(_.get(data, field.path), field)
      el && elements.push(el)
    }

    let title
    if (currentField.isVisible) {
      title = (
        <>
          <EuiTitle size="xxxs" css={{ color: 'dodgerblue' }}>
            <h4>{currentField.name}</h4>
          </EuiTitle>
          <EuiSpacer size="xs"></EuiSpacer>
        </>
      )
    }

    elements = elements.concat(<EuiSpacer size="s" />, ...panelElements)

    if (elements.length > 0) {
      element = (
        <>
          <EuiPanel paddingSize="s">
            {title}
            {elements}
          </EuiPanel>
          <EuiSpacer size="s"></EuiSpacer>
        </>
      )
    }
    return element
  }

  function createCardField(
    value: any,
    field: { isEditable: boolean; showCollapsed: boolean; isVisible: boolean; path: string; name: string },
  ) {
    let element

    if (value instanceof Date) {
      element = (
        <EuiText size="s" key={`${makeId('')}`}>
          <strong>{field.name}:</strong> {value.toString()}
        </EuiText>
      )
    } else if (typeof value === 'string' || typeof value === 'number') {
      if (typeof value === 'string') {
        if (isValidURL(value)) {
          value = <EuiLink href={value}>{value}</EuiLink>
        }
      }
      element = (
        <EuiText size="s" key={`${makeId('')}`}>
          <strong>{field.name}:</strong> {value}
        </EuiText>
      )
    } else if (Array.isArray(value)) {
      console.log('here')
      // TODO generate color from name.
      // workitem tags and types rendered separately from this, explicitly and have custom color
      const badges = value.map((item, idx) => (
        <EuiBadge key={`${makeId('')}-${idx}`} color={generateColor(item)}>
          {item}
        </EuiBadge>
      ))
      element = (
        <div key={`${makeId('')}`}>
          <strong>{field.name}:</strong> {badges}
        </div>
      )
    } else if (typeof value === 'boolean') {
      element = (
        <StyledEuiCheckbox
          key={`${makeId('')}`}
          readOnly
          style={{ alignContent: 'center' }}
          compressed
          id={`checkbox-${makeId('')}`}
          label={field.name}
          onChange={() => null}
          checked={value}
        ></StyledEuiCheckbox>
      )
    }
    return element
  }
}
function getNestedObjects(
  nestedObjects: { isEditable: boolean; showCollapsed: boolean; isVisible: boolean; path: string; name: string }[],
  field: { isEditable: boolean; showCollapsed: boolean; isVisible: boolean; path: string; name: string },
) {
  return nestedObjects.filter((f) => {
    const fieldCount = _.countBy(f.path)['.'] || 0
    const parentFieldCount = _.countBy(field.path)['.'] || 0
    return fieldCount > parentFieldCount
  })
}

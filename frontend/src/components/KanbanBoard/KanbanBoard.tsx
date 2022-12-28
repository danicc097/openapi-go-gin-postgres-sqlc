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
      isVisible: false,
      path: 'demoProjectWorkItem.metadata',
      name: 'Metadata',
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
      path: 'demoProjectWorkItem.metadata.externalLink',
      name: 'External link',
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
    const ignoreFields = []

    return (
      <>
        {boardConfig.fields.map((field, i) => {
          if (ignoreFields.includes(field.path)) return
          if (_.get(data, field.path) === undefined) return

          const value = _.get(data, field.path)
          // TODO group by type and then render
          // nested items inside EuiPanel

          const elements = []

          let element

          if (typeof value === 'object') {
            const nestedFields = boardConfig.fields.filter((f) => f.path.startsWith(field.path))
            nestedFields.forEach((field) => {
              ignoreFields.push(field.path)
              const el = createCardField(_.get(data, field.path), i, field)
              el && elements.push(el)
            })

            let title
            if (field.isVisible) {
              title = (
                <>
                  <EuiTitle size="xxxs">
                    <h4>{field.name}</h4>
                  </EuiTitle>
                  <EuiSpacer size="xs"></EuiSpacer>
                </>
              )
            }

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
          } else {
            element = createCardField(value, i, field)
          }

          console.log(elements)

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
            }
            description={
              <span>
                Just be sure not to add any <EuiCode>onClick</EuiCode> handler to the card if the children are also
                interactable.
              </span>
            }
            hasBorder={false}
            paddingSize="none"
            display="plain"
            // footer={'footer'}
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
              style={{ minWidth: '30vw' }}
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

  function createCardField(
    value: any,
    i: number,
    field: { isEditable: boolean; showCollapsed: boolean; isVisible: boolean; path: string; name: string },
  ) {
    let element

    if (value instanceof Date) {
      element = (
        <EuiText size="s" key={i}>
          <strong>{field.name}:</strong> {value.toString()}
        </EuiText>
      )
    } else if (typeof value === 'string' || typeof value === 'number') {
      element = (
        <EuiText size="s" key={i}>
          <strong>{field.name}:</strong> {value}
        </EuiText>
      )
    } else if (Array.isArray(value)) {
      // TODO generate color from name.
      // workitem tags and types rendered separately from this, explicitly and have custom color
      const badges = value.map((item, idx) => (
        <EuiBadge key={`${i}-${idx}`} color={generateColor(item)}>
          {item}
        </EuiBadge>
      ))
      element = (
        <div key={i}>
          <strong>{field.name}:</strong> {badges}
        </div>
      )
    } else if (typeof value === 'boolean') {
      element = (
        <StyledEuiCheckbox
          key={i}
          readOnly
          style={{ alignContent: 'center' }}
          compressed
          id={`checkbox-${i}`}
          label={field.name}
          onChange={() => null}
          checked={value}
        ></StyledEuiCheckbox>
      )
    }
    return element
  }
}

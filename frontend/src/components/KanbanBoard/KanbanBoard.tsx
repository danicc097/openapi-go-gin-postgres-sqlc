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
} from '@elastic/eui'
import { ToastId } from 'src/utils/toasts'
import { useUISlice } from 'src/slices/ui'
import { isArray, isObject, random, uniqueId } from 'lodash'
import type { DemoProjectWorkItemsResponse } from 'src/gen/model'
import moment from 'moment'
import { getGetProjectWorkitemsMock, getProjectMSW } from 'src/gen/project/project.msw'

const makeId = htmlIdGenerator()

type SampleCardData = {
  someDate: Date
  someBoolean: boolean
  someOtherBoolean: boolean
  sometext: string
  someOptText?: string
  someList: string[]
}

// UI title mapping for fields
const CardDataNames: Record<keyof SampleCardData, string> = {
  someDate: 'Some date',
  someBoolean: 'Some boolean',
  someOtherBoolean: 'Some other boolean',
  sometext: 'Some text',
  someOptText: 'Some optional text',
  someList: 'Some list',
}

// config panel allows naming as per full path, showing e.g.
// baseWorkItem.workItemTypeID : [ Name   ] [✓] Visible
const ProjectWorkitemsNames: Record<string, string> = {
  'baseWorkItem.workItemTypeID': 'Work item type ID',
}

export default function KanbanBoard() {
  const { addToast } = useUISlice()

  const sampleCard: SampleCardData = {
    someBoolean: true, // checkbox
    someOtherBoolean: false, // checkbox
    someDate: new Date(),
    someList: ['item 1', 'item 2'], // EuiBadge
    sometext: 'content for sometext.\n More content.',
  }

  const demoProjectWI = getGetProjectWorkitemsMock()

  // to ignore nested objects fields configuration panel needs
  // to use fully qualified name based on path
  // then _.get by `path` from field name mapping
  const renderComponents = (parentPath: string[], [key, val]): void => {
    const path = [...parentPath, String(key)]
    console.log(path.join('.'))
    if (val === null || val === undefined) {
      return
    }
    // won't catch custom classes
    if (val?.constructor === Object) {
      console.log(`${key} is an object`)
      Object.entries(val).forEach(([key, val]) => renderComponents(path, [key, val]))
      console.log(`end of nested object ${key}`)
    } else if (isArray(val)) {
      console.log(`${key} is a list`)
    } else if (val instanceof Date) {
      console.log(`${key} is a date`)
    } else if (typeof val === 'boolean') {
      console.log(`${key} is a boolean`)
    } else if (typeof val === 'string') {
      console.log(`${key} is a string`)
    }
  }

  console.log('%c demoProjectWI', 'color: #c92a2a')
  console.log(demoProjectWI)
  Object.entries(demoProjectWI).forEach(([key, val]) => renderComponents([], [key, val]))

  const makeList = (number, start = 1) =>
    Array.from({ length: number }, (v, k) => k + start).map((el) => {
      return {
        content: (
          <EuiCard
            textAlign="left"
            title={`Card ${el}`}
            description={
              <span>
                Just be sure not to add any <EuiCode>onClick</EuiCode> handler to the card if the children are also
                interactable.
              </span>
            }
            hasBorder={false}
            paddingSize="none"
            display="plain"
          >
            <EuiCodeBlock language="html" paddingSize="s">
              {'<yoda>Hello, young Skywalker</yoda>'}
            </EuiCodeBlock>
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
  )
}

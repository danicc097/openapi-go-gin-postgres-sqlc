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
import { isArray, random, uniqueId } from 'lodash'

const makeId = htmlIdGenerator()

type SampleCardData = {
  someDate: Date
  someBoolean: boolean
  someOtherBoolean: boolean
  sometext: string
  someOptText?: string
  someList: string[]
}

const CardDataNames: Record<keyof SampleCardData, string> = {
  someDate: 'Some date',
  someBoolean: 'Some boolean',
  someOtherBoolean: 'Some other boolean',
  sometext: 'Some text',
  someOptText: 'Some optional text',
  someList: 'Some list',
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

  const [list, setList] = useState([1, 2])
  const [list1, setList1] = useState(makeList(3))
  const [list2, setList2] = useState(makeList(3, 4))
  const lists = {
    COMPLEX_DROPPABLE_PARENT: list,
    COMPLEX_DROPPABLE_AREA_1: list1,
    COMPLEX_DROPPABLE_AREA_2: list2,
  }
  const actions = {
    COMPLEX_DROPPABLE_PARENT: setList,
    COMPLEX_DROPPABLE_AREA_1: setList1,
    COMPLEX_DROPPABLE_AREA_2: setList2,
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
        style={{ display: 'flex' }}
      >
        {list.map((did, didx) => (
          <EuiDraggable
            key={did}
            index={didx}
            draggableId={`COMPLEX_DRAGGABLE_${did}`}
            spacing="l"
            style={{ flex: '1 0 50%' }}
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

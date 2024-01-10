import { Popover } from '@mantine/core'
import { Tooltip as ReactTooltip } from 'react-tooltip'

export interface CustomPopoverProps {
  content: JSX.Element
  selector: string
}

// For tours, use https://docs.react.tours/tour/quickstart instead
export default function CustomPopover({ content, selector }: CustomPopoverProps) {
  return (
    <ReactTooltip
      id="my-tooltip"
      place="right-end"
      isOpen
      anchorSelect={selector}
      imperativeModeOnly // gets crazy repositioning
      style={{ zIndex: 10000 }}
      content={
        (
          <Popover id="my-tooltip--card" withinPortal>
            {content}
          </Popover>
        ) as any
      }
      clickable
    />
  )
}

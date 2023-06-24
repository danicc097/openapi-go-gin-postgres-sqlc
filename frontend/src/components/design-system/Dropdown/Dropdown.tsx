import React from 'react'
import { Menu, type MenuProps } from '@mantine/core'

import useStyles from './Dropdown.styles'
import { shadows } from '../config'

interface IDropdownProps
  extends Pick<
    MenuProps,
    | 'withArrow'
    | 'opened'
    | 'offset'
    | 'position'
    | 'onOpen'
    | 'onClose'
    | 'withinPortal'
    | 'middlewares'
    | 'disabled'
    | 'width'
    | 'styles'
  > {
  control: React.ReactNode
  children: React.ReactNode
}

export function Dropdown({ control, withArrow = true, offset = 10, children, ...props }: IDropdownProps) {
  const { classes, theme } = useStyles()

  return (
    <Menu
      withArrow={withArrow}
      transitionProps={{ duration: 0 }}
      radius={7}
      offset={offset}
      shadow={theme.colorScheme === 'dark' ? shadows.dark : shadows.light}
      classNames={classes}
      clickOutsideEvents={['click', 'mousedown', 'touchstart']}
      middlewares={{ flip: false, shift: false }}
      {...props}
    >
      <Menu.Target>{control}</Menu.Target>
      <Menu.Dropdown>{children}</Menu.Dropdown>
    </Menu>
  )
}

Dropdown.Item = Menu.Item
Dropdown.Label = Menu.Label
Dropdown.Divider = Menu.Divider

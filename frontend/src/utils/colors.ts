import _ from 'lodash'
import type { Role, WorkItemRole } from 'src/gen/model'
import { DefaultMantineColor } from '@mantine/core'

const LIGHT_BLUE = '#4EC5F1'
const LIGHT_GREEN = '#0DF2C8'
const LIGHT_ORANGE = '#C072DA'
const LIGHT_GREY = '#BAB5B5'
const LIGHT_RED = '#9C1C1C'
const LIGHT_PURPLE = '#5E2F74'

type RoleColors = {
  [key in Role]: string
}

const ROLE_COLORS: RoleColors = {
  guest: LIGHT_GREY,
  user: LIGHT_BLUE,
  advancedUser: LIGHT_ORANGE,
  manager: LIGHT_GREEN,
  admin: LIGHT_PURPLE,
  superAdmin: LIGHT_RED,
}

export const roleColor = (role: Role) => {
  return ROLE_COLORS[role]
}

type WorkItemRoleColors = {
  [key in WorkItemRole]: string
}

const WORK_ITEM_ROLE_COLORS: WorkItemRoleColors = {
  preparer: LIGHT_GREY,
  reviewer: LIGHT_ORANGE,
}

export const workItemRoleColor = (role: WorkItemRole) => {
  return WORK_ITEM_ROLE_COLORS[role]
}

export const COLORS = [
  '#00BFB3',
  '#FF6D6D',
  '#0b8f77',
  '#FFB200',
  '#7F00FF',
  '#FF8C00',
  '#00BFFF',
  '#FF00FF',
  '#DCDCDC',
  '#cf2620',
  '#008080',
  '#FFD700',
  '#FFA500',
  '#FF4500',
  '#800000',
  '#800080',
  '#808000',
  '#00FF00',
  '#00FFFF',
  '#000080',
  '#0000FF',
  '#4B0082',
  '#EE82EE',
  '#00BFB3',
  '#FF6D6D',
  '#0b8f77',
  '#FFB200',
  '#7F00FF',
  '#FF8C00',
]

export const COLOR_BLIND_PALETTE = ['#999999', '#E69F00', '#56B4E9', '#009E73', '#0072B2', '#D55E00', '#CC79A7']

export function getContrastYIQ(hexColor) {
  if (!hexColor) return `#aaaaaa`

  const hex = hexColor.replace('#', '')
  const [r, g, b] = hex.match(/.{2}/g).map((val) => parseInt(val, 16))

  const yiq = (r * 299 + g * 587 + b * 114) / 1000

  return yiq >= 128 ? 'white' : 'black'
}
export function generateColor(str: string): string {
  if (str === '' || !str) {
    return '#aaa'
  }
  let num = 0
  for (const ch of str) {
    num += ch.charCodeAt(0)
  }

  let hex = Math.floor(num * 4554323).toString(16)
  if (hex.length < 6) {
    hex = _.padStart(hex, 6, '0')
  } else if (hex.length > 6) {
    hex = hex.slice(0, 6)
  }

  return `#${hex}`
}

export const colorBlindPalette = [
  // Red-Green Colorblind Friendly
  '#004c6d', // Dark Blue
  '#177e89', // Teal
  '#56a3a6', // Light Teal
  '#f76f8e', // Salmon
  '#8e063b', // Dark Red
  '#c91f37', // Red
  '#f15a29', // Orangee
  '#f6a02e', // Yellow
  '#f5d44e', // Light Yellow
  '#bedbbb', // Pale Green
  '#6aa84f', // Green
  '#87a878', // Light Green

  // Blue-Yellow Colorblind Friendly
  '#153956', // Dark Blue
  '#1f699e', // Blue
  '#3d85c6', // Sky Blue
  '#98b4d4', // Light Blue
  '#f3be56', // Yellow
  '#f18d32', // Orange
  '#ea5837', // Salmon
  '#c93756', // Red
  '#5a343b', // Dark Red
  '#6c7f5a', // Green
  '#a6b985', // Light Green
  '#c0d29c', // Pale Green

  // Purple-Blue Colorblind Friendly
  '#212946', // Dark Blue
  '#324e7b', // Blue
  '#4968a6', // Royal Blue
  '#6c8dc6', // Sky Blue
  '#b0b8db', // Light Blue
  '#ce5a57', // Red
  '#e37973', // Light Red
  '#ed9c96', // Salmon
  '#f2b6b3', // Light Salmon
  '#f8cbb5', // Peach
  '#ffb56c', // Orange
  '#ffc778', // Light Orange

  // Gray Colorblind Friendly
  '#333333', // Dark Gray
  '#666666', // Gray
  '#999999', // Light Gray
  '#cccccc', // Pale Gray
  '#f7f7f7', // White
  '#000000', // Black

  // Rainbow Colorblind Friendly
  '#1f77b4', // Blue
  '#ff7f0e', // Orange
  '#2ca02c', // Green
  '#d62728', // Red
  '#9467bd', // Purple
  '#8c564b', // Brown
  '#e377c2', // Pink
  '#7f7f7f', // Gray
  '#bcbd22', // Olive
  '#17becf', // Teal
]

export function scopeColor(scopePermission?: string): string {
  switch (scopePermission) {
    case 'read':
      return '#008000'
    case 'write':
    case 'edit':
      return '#f8ce45'
    case 'delete':
      return '#990000'
    default:
      return '#1971c2'
  }
}

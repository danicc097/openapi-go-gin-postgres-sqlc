import SCOPES from 'src/scopes'
import _ from 'lodash'
import type { Role } from 'src/gen/model'

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

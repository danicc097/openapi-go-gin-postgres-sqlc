import { makeHighContrastColor } from '@elastic/eui'
import type roles from '@roles'
import scopes from '@scopes'

const LIGHT_BLUE = 'rgb(78, 197, 241)'
const LIGHT_GREEN = 'rgb(13, 242, 200)'
const LIGHT_ORANGE = 'rgb(192, 114, 218)'
const LIGHT_GREY = 'rgb(186, 181, 181)'
const LIGHT_RED = 'rgb(156, 28, 28)'
const LIGHT_PURPLE = 'rgb(94, 47, 116)'

type RoleColors = {
  [key in keyof typeof roles]: string
}

const ROLE_COLORS: RoleColors = {
  guest: LIGHT_GREY,
  user: LIGHT_BLUE,
  advancedUser: LIGHT_ORANGE,
  manager: LIGHT_GREEN,
  admin: LIGHT_PURPLE,
  superAdmin: LIGHT_RED,
}

export const roleColor = (role: keyof typeof roles) => {
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

export const getContrastYIQ = (hc) => {
  const [r, g, b] = [0, 2, 4].map((p) => parseInt(hc.substr(p, 2), 16))
  return (r * 299 + g * 587 + b * 114) / 1000 >= 128 ? 'black' : 'white'
}

import { capitalize } from 'lodash'
import type { UserAvatarProps } from 'src/components/UserAvatar/UserAvatar'

export const getAvatarName = ({ user }: UserAvatarProps) => capitalize(user?.full_name ?? user?.username ?? 'Anonymous')

const currencyFormatter = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
})

export const formatPrice = (price: number) => (price ? currencyFormatter.format(price) : price)

export const truncate = (str: string, n = 200, useWordBoundary = false) => {
  if (!str || str?.length <= n) return str
  const subString = str.substr(0, n - 1)
  return `${useWordBoundary ? subString.substr(0, subString.lastIndexOf(' ')) : subString}&hellip;`
}

export const joinWithAnd = (arr: string[]) => {
  if (arr.length === 1) return arr[0]
  return `${arr.slice(0, -1).join(', ')} and ${arr.slice(-1)[0]}`
}

export const getMaxStringPixelLength = (arr: string[]) => {
  return Math.max(...arr.map((str) => getStringPixelLength(str))) + 20
}

export const getStringPixelLength = (str: string) => {
  const span = document.createElement('span')
  span.style.visibility = 'hidden'
  span.style.fontFamily = 'sans-serif'
  span.style.position = 'absolute'
  span.style.whiteSpace = 'nowrap'
  span.style.fontSize = '12px'
  span.textContent = str
  document.body.appendChild(span)
  const width = span.clientWidth
  document.body.removeChild(span)
  return width
}

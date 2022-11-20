import { capitalize } from 'lodash'
import { COLOR_BLIND_PALETTE, getContrastYIQ } from 'src/utils/colors'

export const createTextFileWithCreds = ({ email, password }: { email: string; password: string }) => {
  const element = document.createElement('a')
  const file = new Blob([`Email: ${email}\nPassword: ${password}`], { type: 'text/plain' })
  element.href = URL.createObjectURL(file)
  element.download = `${email}.txt`
  document.body.appendChild(element)
  element.click()
  element.remove()
}

export const createAvatarImageDataUrl = (str: string) => {
  const normStr = capitalize(str)
  const canvas = document.createElement('canvas')
  canvas.width = 150
  canvas.height = 150
  const ctx = canvas.getContext('2d')
  const color = COLOR_BLIND_PALETTE[normStr.charCodeAt(0) % COLOR_BLIND_PALETTE.length]
  ctx.fillStyle = color
  ctx.beginPath()
  ctx.arc(75, 75, 75, 0, 2 * Math.PI)
  ctx.fill()
  ctx.font = 'bold 70px "Exo 2"'
  ctx.fillStyle = getContrastYIQ(color) === 'black' ? 'white' : 'black'
  ctx.textAlign = 'center'
  ctx.textBaseline = 'middle'
  ctx.fillText(normStr.substring(0, 2), 75, 75)
  const dataUrl = canvas.toDataURL('image/png')
  console.log('dataUrl', dataUrl)
  return dataUrl
}

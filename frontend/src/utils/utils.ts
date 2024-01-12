export function getCurrentFileName() {
  const error = new Error()
  const stack = error.stack || ''
  console.log({ stack })
  const caller = stack.split('\n')[2]
  const fileNameMatch = caller?.match(/\/([^\/]+\.jsx|\.js)/)

  if (fileNameMatch) {
    return fileNameMatch[1]
  }

  return null
}

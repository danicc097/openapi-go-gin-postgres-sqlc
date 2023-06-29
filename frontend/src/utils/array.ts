export function removeElementByIndex(arr: any[], index: number) {
  if (!Array.isArray(arr) || arr.length === 0) {
    return arr
  }

  if (index >= arr.length) {
    return arr
  }

  if (index === arr.length - 1 || index === -1) {
    arr.pop()
  } else {
    arr.splice(index, 1)
  }

  return arr
}

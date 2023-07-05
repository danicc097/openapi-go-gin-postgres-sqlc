/**
 *  Removes element mutating original array.
 */
export function removeElementByIndex(arr: any[], index: number) {
  if (!Array.isArray(arr) || arr.length === 0) {
    return
  }

  if (index >= arr.length) {
    return
  }

  if (index === arr.length - 1 || index === -1) {
    arr.pop()
  } else {
    arr.splice(index, 1)
  }
}

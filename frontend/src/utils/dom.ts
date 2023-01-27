interface Coord {
  xStart: number
  yStart: number
  xEnd: number
  yEnd: number
}

export function getAbsolutePosition(element: HTMLElement): Coord {
  if (!element) return

  const rect = element.getBoundingClientRect()
  const scrollLeft = window.pageXOffset || document.documentElement.scrollLeft
  const scrollTop = window.pageYOffset || document.documentElement.scrollTop
  return {
    xStart: rect.left + scrollLeft,
    yStart: rect.top + scrollTop,
    xEnd: rect.right, // fixme
    yEnd: rect.bottom, // fixme
  }
}

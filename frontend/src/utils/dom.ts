interface Coord {
  xStart: number
  yStart: number
  xEnd: number
  yEnd: number
}

export function getAbsolutePosition(element: HTMLElement): Coord {
  if (!element) return { xStart: 0, yStart: 0, xEnd: 0, yEnd: 0 }

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

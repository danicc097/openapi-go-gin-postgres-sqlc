import { useEffect, useRef, useState } from 'react'

export function useCarousel(items: any[], interval: number): Record<string, any> {
  const timeoutRef: any = useRef()

  const [shouldAnimate, setShouldAnimate] = useState<boolean>(true)
  const [current, setCurrent] = useState<number>(0)

  useEffect(() => {
    const next = (current + 1) % items.length
    if (shouldAnimate) {
      timeoutRef.current = setTimeout(() => setCurrent(next), interval)
    }

    return () => clearTimeout(timeoutRef.current)
  }, [current, items.length, interval, shouldAnimate])

  return { current, setShouldAnimate, timeoutRef }
}

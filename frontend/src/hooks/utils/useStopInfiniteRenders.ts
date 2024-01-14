import { useEffect, useRef } from 'react'

const useRenderCount = () => {
  // src: https://medium.com/better-programming/how-to-properly-use-the-react-useref-hook-in-concurrent-mode-38c54543857b
  const renderCount = useRef(0)
  let renderCountLocal = renderCount.current
  useEffect(() => {
    renderCount.current = renderCountLocal
  })
  renderCountLocal++
  return renderCount.current
}

export default function useStopInfiniteRenders(maxRenders: number) {
  const renderCount = useRenderCount()

  if (renderCount > maxRenders && import.meta.env.DEV) throw new Error(`Infinite renders limit reached (${maxRenders})`)
}

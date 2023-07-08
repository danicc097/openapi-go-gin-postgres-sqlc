import { useRef } from 'react'

export default function useRenders() {
  const rendersRef = useRef(0)
  rendersRef.current++
  return rendersRef.current
}

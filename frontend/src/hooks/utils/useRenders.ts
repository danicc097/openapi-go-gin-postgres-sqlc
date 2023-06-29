import { useRef } from 'react'

export default function useRenders() {
  let rendersRef = useRef(0)
  rendersRef.current++
  return rendersRef.current
}

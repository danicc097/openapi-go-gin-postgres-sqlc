import { EuiLoadingSpinner, EuiLoadingSpinnerProps } from '@elastic/eui'
import { motion } from 'framer-motion'
import React from 'react'

export default function InfiniteSpinner(props: EuiLoadingSpinnerProps) {
  return (
    <motion.div
      animate={{ rotate: 360 }}
      transition={{ duration: 1, repeat: Infinity }}
      style={{ width: '100%', height: '100%', display: 'flex', justifyContent: 'center', alignItems: 'center' }}
    >
      <EuiLoadingSpinner {...props} />
    </motion.div>
  )
}

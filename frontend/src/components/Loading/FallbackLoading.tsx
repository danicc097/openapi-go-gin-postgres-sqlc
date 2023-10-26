import { motion } from 'framer-motion'
import cloudsDark from 'src/assets/logo/two-white-clouds.svg'
import cloudsLight from 'src/assets/logo/two-black-clouds.svg'
import React from 'react'
import { RouteLoading } from './FallbackLoading.styles'
import { useUISlice } from 'src/slices/ui'
import { useMantineColorScheme, useMantineTheme } from '@mantine/core'

export default function FallbackLoading() {
  const { colorScheme } = useMantineColorScheme()

  return (
    <RouteLoading>
      <motion.div
        className="logo"
        animate={{
          y: [0, -20, 0],
          rotate: [0, 0, 0],
          transition: {
            duration: 2,
            loop: Infinity,
            ease: 'easeInOut',
          },
        }}
      >
        <img src={colorScheme === 'dark' ? cloudsDark : cloudsLight} width="80" />
      </motion.div>
      {/* // animate a boxshadow below the svg that grows and shrinks width*/}
      <motion.div
        animate={{
          // zoom in in the x direction
          zoom: [1, 1.2, 1],
          transition: {
            duration: 2,
            loop: Infinity,
            ease: 'easeInOut',
          },
        }}
      >
        <div
          style={{
            width: '110px',
            height: '10px',
            float: 'right',
            left: '50%',
            bottom: '50%',
            borderRadius: '50%',
            boxShadow: '0 50px 14px rgba(0, 0, 0, 0.64)',
          }}
        ></div>
      </motion.div>
    </RouteLoading>
  )
}

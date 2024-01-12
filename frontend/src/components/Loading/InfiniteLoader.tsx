import { motion } from 'framer-motion'
import React from 'react'
import classes from './InfiniteLoader.module.css'
import { Flex } from '@mantine/core'

export default function InfiniteLoader() {
  return (
    <Flex direction={'row'}>
      <div className={classes.dotLoader}></div>
      <div className={classes.dotLoader}></div>
      <div className={classes.dotLoader}></div>
    </Flex>
  )
}

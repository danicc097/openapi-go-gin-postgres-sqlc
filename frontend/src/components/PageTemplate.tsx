import React, { type ReactElement } from 'react'
import { Container, Paper } from '@mantine/core'

type PageTemplateProps = {
  children: ReactElement
}

const PageTemplate = ({ children }: PageTemplateProps) => {
  return (
    <Container size="sm" style={{ paddingTop: '2rem', paddingBottom: '2rem' }}>
      <Paper p="md" shadow="sm">
        {children}
      </Paper>
    </Container>
  )
}

export default PageTemplate

import React, { type ReactElement } from 'react'
import { Container, Paper, useMantineTheme } from '@mantine/core'
import { css } from '@emotion/react'

type PageTemplateProps = {
  children: ReactElement
  minWidth?: string | number
}

const PageTemplate = ({ children, minWidth }: PageTemplateProps) => {
  const theme = useMantineTheme()

  return (
    <Container size="sm" style={{ paddingTop: '2rem', paddingBottom: '2rem', minWidth }}>
      <Paper p="md" shadow="lg" c={theme.primaryColor}>
        {children}
      </Paper>
    </Container>
  )
}

export default PageTemplate

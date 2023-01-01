import { useNavigate } from 'react-router-dom'
import { EuiEmptyPrompt, EuiButton } from '@elastic/eui'
import React from 'react'

type NotFoundPageProps = {
  notFoundItem?: string
  notFoundError?: string
}
export default function NotFoundPage({
  notFoundItem = 'Page',
  notFoundError = "Looks like there's nothing there. We must have misplaced it!",
}: NotFoundPageProps) {
  const navigate = useNavigate()

  return (
    <EuiEmptyPrompt
      iconType="editorStrike"
      title={<h2>{notFoundItem} Not Found</h2>}
      body={<p>{notFoundError}</p>}
      actions={
        <EuiButton color="primary" fill onClick={() => navigate(-1)}>
          Go Back
        </EuiButton>
      }
    />
  )
}

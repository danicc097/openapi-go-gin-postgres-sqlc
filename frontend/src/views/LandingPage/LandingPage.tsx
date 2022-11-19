import { EuiPage, EuiPageBody, EuiFlexGroup, EuiFlexItem, EuiTitle, EuiText, EuiButton } from '@elastic/eui'
import { useEffect } from 'react'
import { useUISlice } from 'src/slices/ui'
import { ToastId } from 'src/utils/toasts'
import * as S from './LandingPage.styles'

export default function LandingPage() {
  const { switchTheme, addToast, removeToast } = useUISlice()

  useEffect(() => {
    addToast({
      id: ToastId.AuthRedirect,
      title: 'redirecting',
      color: 'warning',
      iconType: 'alert',
      toastLifeTimeMs: 15000,
      text: 'A message about redirection.',
    })
  }, [addToast])

  return (
    <EuiPage>
      <EuiPageBody component="section" align="center">
        <EuiFlexGroup direction="column" alignItems="center">
          <EuiFlexItem>
            <EuiTitle>
              <EuiText>My App</EuiText>
            </EuiTitle>
          </EuiFlexItem>
          <EuiFlexItem>
            <EuiButton
              onClick={() =>
                addToast({
                  id: ToastId.AuthzError,
                  title: 'clicked',
                  color: 'success',
                  iconType: 'alert',
                  toastLifeTimeMs: 15000,
                  text: 'clicked.',
                })
              }
            >
              Submit
            </EuiButton>
            <EuiButton
              onClick={() =>
                removeToast({
                  id: ToastId.AuthzError,
                })
              }
            >
              Submit
            </EuiButton>
          </EuiFlexItem>
        </EuiFlexGroup>
      </EuiPageBody>
    </EuiPage>
  )
}

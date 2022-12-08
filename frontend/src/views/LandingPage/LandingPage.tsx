import { EuiPage, EuiPageBody, EuiFlexGroup, EuiFlexItem, EuiTitle, EuiText, EuiButton } from '@elastic/eui'
import { useEffect } from 'react'
import KanbanBoard from 'src/components/KanbanBoard/KanbanBoard'
import PageTemplate from 'src/components/PageTemplate/PageTemplate'
import { useUISlice } from 'src/slices/ui'
import { ToastId } from 'src/utils/toasts'

export default function LandingPage() {
  const { setTheme, addToast, dismissToast } = useUISlice()

  useEffect(() => {
    null
  }, [])

  return (
    <PageTemplate
      content={<KanbanBoard></KanbanBoard>}
      header={{ description: 'My header' }}
      buttons={[
        <EuiButton
          key={1}
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
          New toast
        </EuiButton>,
      ]}
    ></PageTemplate>
  )
}

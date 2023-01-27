import { useState } from 'react'
import { EuiConfirmModal, EuiConfirmModalProps } from '@elastic/eui'

type ModalConfig = {
  title: string
  message: string
  buttonColor: EuiConfirmModalProps['buttonColor']
  onCancel: () => void
  onConfirm: () => void
}

export default function useConfirmModal(config: ModalConfig) {
  const [isModalVisible, setIsModalVisible] = useState(false)

  const showModal = () => setIsModalVisible(true)

  const hideModal = () => setIsModalVisible(false)

  const renderModal = () => (
    <EuiConfirmModal
      title={config.title}
      onCancel={() => {
        hideModal()
        config.onCancel()
      }}
      onConfirm={() => {
        hideModal()
        config.onConfirm()
      }}
      cancelButtonText="Cancel"
      confirmButtonText="Confirm"
      defaultFocusedButton="cancel"
      buttonColor={config.buttonColor}
    >
      {config.message}
    </EuiConfirmModal>
  )

  return { isModalVisible, showModal, hideModal, renderModal }
}

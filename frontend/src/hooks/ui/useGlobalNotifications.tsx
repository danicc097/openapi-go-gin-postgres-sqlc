import moment from 'moment'
import { useEffect, useCallback } from 'react'
import { shallowEqual } from 'react-redux'
import { useAppDispatch, useAppSelector } from 'src/redux/hooks'
import { GlobalNotificationsActionCreators } from 'src/redux/modules/feed/globalNotifications'
import { extractErrorMessages } from 'src/utils/errors'

export function useGlobalNotifications() {
  const dispatch = useAppDispatch()

  const isLoading = useAppSelector((state) => state.feed.globalNotifications.isLoading)
  const error = useAppSelector((state) => state.feed.globalNotifications.error, shallowEqual)
  const errorList = extractErrorMessages(error)

  const createNotification = ({ notification }) => {
    return dispatch(GlobalNotificationsActionCreators.createNotification({ notification }))
  }

  const deleteNotification = ({ id }) => {
    return dispatch(GlobalNotificationsActionCreators.deleteNotification({ id }))
  }

  return {
    isLoading,
    errorList,
    createNotification,
    deleteNotification,
  }
}

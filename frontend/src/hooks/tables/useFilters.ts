import { pluralize } from 'inflection'
import { useState } from 'react'

const DELETED_ENTITY_FILTER_STATES = [null, true, false]

export const useDeletedEntityFilter = (entity: string) => {
  const [deletedEntityFilter, setDeletedEntityFilter] = useState(0)

  const toggleDeletedUsersFilter = () => {
    const newStatus = (deletedEntityFilter + 1) % DELETED_ENTITY_FILTER_STATES.length
    setDeletedEntityFilter(newStatus)
  }

  const getLabelText = () => {
    switch (DELETED_ENTITY_FILTER_STATES[deletedEntityFilter]) {
      case null:
        return `Showing all ${pluralize(entity)}`
      case true:
        return `Showing deleted ${pluralize(entity)} only`
      case false:
        return `Hiding deleted ${pluralize(entity)}`
      default:
        return ''
    }
  }

  return {
    getLabelText,
    toggleDeletedUsersFilter,
    deletedEntityFilterState: DELETED_ENTITY_FILTER_STATES[deletedEntityFilter],
  }
}

import { AnyAction } from '@reduxjs/toolkit'

export function errorState(state, action: AnyAction) {
  return {
    ...state,
    isLoading: false,
    error: action.error,
  }
}

export function loadingState(state) {
  return {
    ...state,
    isLoading: true,
  }
}

export function successState(state) {
  return {
    ...state,
    isLoading: false,
    error: null,
  }
}

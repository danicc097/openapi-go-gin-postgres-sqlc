import { CaseReducer, createSlice, PayloadAction } from '@reduxjs/toolkit'
import { RootState } from '../store'

type Theme = 'dark' | 'light'

type SliceState = {
  theme: Theme
  toastList: unknown[]
}

const addToast: CaseReducer<SliceState, PayloadAction<unknown>> = (state, action) => {
  state.toastList.push(action.payload)
}

const switchTheme: CaseReducer<SliceState, PayloadAction<null>> = (state) => {
  state.theme = state.theme === 'dark' ? 'light' : 'dark'
  // TODO save to LS
}

const uiSlice = createSlice({
  name: 'ui',
  initialState: {
    theme: 'dark', // TODO load from LS
    toastList: [],
  } as SliceState,
  reducers: {
    switchTheme,
    addToast,
  },
})

export default uiSlice

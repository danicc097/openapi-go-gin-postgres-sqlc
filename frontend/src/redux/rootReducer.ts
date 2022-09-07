import { combineReducers } from 'redux'
import { internalApi } from './slices/gen/internalApi'
import uiSlice from './slices/ui'

const rootReducer = combineReducers({
  internalApi: internalApi.reducer,
  ui: uiSlice.reducer,
})
export default rootReducer

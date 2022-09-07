import { combineReducers } from 'redux'
import { internalApi } from './slices/gen/internalApi'

const rootReducer = combineReducers({
  internalApi: internalApi.reducer,
})
export default rootReducer

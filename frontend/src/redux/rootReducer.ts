import { combineReducers } from 'redux'
import { internalApi } from '../store/internalApi'

const rootReducer = combineReducers({
  internalApi: internalApi.reducer,
})
export default rootReducer

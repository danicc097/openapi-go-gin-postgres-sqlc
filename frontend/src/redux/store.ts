import { configureStore, ThunkAction, Action, Store } from '@reduxjs/toolkit'
import { internalApi } from 'src/redux/slices/gen/internalApi'
import rootReducer from './rootReducer'

// thunk mw already in rtk's configureStore
const store = configureStore({
  reducer: rootReducer,
  ...(import.meta.env.NODE_ENV === 'production' && { devTools: false }),
  middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(internalApi.middleware),
})

export default function configureReduxStore() {
  // store.dispatch(AuthActionCreators.fetchUserFromToken());

  if (import.meta.env.NODE_ENV !== 'production' && import.meta.hot) {
    import.meta.hot.accept('./rootReducer', () => store.replaceReducer(rootReducer))
  }

  return store
}

export type AppDispatch = typeof store.dispatch

export type RootState = ReturnType<typeof store.getState>

export type AppThunk<ReturnType = void> = ThunkAction<ReturnType, RootState, unknown, Action<string>>

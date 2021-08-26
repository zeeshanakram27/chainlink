import { Reducer } from 'redux'
import { Actions, ResourceActionType } from './actions'

export type State = Record<string, object>

const INITIAL_STATE: State = {}

const reducer: Reducer<State, Actions> = (state = INITIAL_STATE, action) => {
  console.log('stx:reducer:state', state, 'stx:reducer:action', action)
  switch (action.type) {
    case ResourceActionType.UPSERT_STACKS_ACCOUNT_BALANCE:
      console.log('action.stx: ', action.data.stx)
      return { ...state, ...action.data.stx }
    default:
      return state
  }
}

export default reducer

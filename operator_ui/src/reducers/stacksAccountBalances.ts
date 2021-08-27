import { Reducer } from 'redux'
import { Actions, ResourceActionType } from './actions'

export type State = Record<string, object>

const INITIAL_STATE: State = {}

const reducer: Reducer<State, Actions> = (state = INITIAL_STATE, action) => {
  switch (action.type) {
    case ResourceActionType.UPSERT_STACKS_ACCOUNT_BALANCE:
      return { ...state, ...action.data.sTxKeys }
    default:
      return state
  }
}

export default reducer

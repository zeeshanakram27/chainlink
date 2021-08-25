import { Reducer } from 'redux'
import { Actions, ResourceActionType } from './actions'

export type State = Record<string, object>

const INITIAL_STATE: State = {}

const reducer: Reducer<State, Actions> = (state = INITIAL_STATE, action) => {
  console.log("eth:reducer:state", state,"eth:reducer:action", action)
  switch (action.type) {
    case ResourceActionType.UPSERT_ACCOUNT_BALANCE:
      return { ...state, ...action.data.eThKeys }
    default:
      return state
  }
}

export default reducer

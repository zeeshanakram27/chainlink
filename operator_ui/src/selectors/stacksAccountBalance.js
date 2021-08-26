import build from 'redux-object'

export default (state) => {
  const address = Object.keys(state.stacksAccountBalances)
  console.log('state', state, 'address', address)
  if (address) {
    return build(state, 'stacksAccountBalances', address)
  }
}

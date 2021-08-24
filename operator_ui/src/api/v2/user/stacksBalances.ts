//import * as jsonapi from 'utils/json-api-client'
import { boundMethod } from 'autobind-decorator'
//import * as presenters from 'core/store/presenters'

/**
 * AccountBalances returns the account balances of STACKS & STX-LINK.
 *
 * @example "<application>/keys/eth"
 */
export const ACCOUNT_BALANCES_ENDPOINT = '/v2/keys/eth'

export class StacksBalances {
  constructor() {}

  /**
   * Get account balances in STACKS and STX-LINK
   */
  @boundMethod
  public async getAccountBalances(): Promise<any> {
    const response = await fetch(
      'https://stacks-node-api.mainnet.stacks.co/extended/v1/address/SP4VX529NN4XTFCR7WJ2VXB4GNGR2R7KHSN29RSK/balances',
      {
        method: 'GET',
      },
    )
    console.log('here we are stacks')
    console.log(response.json())
    return response.json()
  }

  //   private accountBalances = this.api.fetchResource<
  //     {},
  //     presenters.AccountBalance[],
  //     {}
  //   >(ACCOUNT_BALANCES_ENDPOINT)
}

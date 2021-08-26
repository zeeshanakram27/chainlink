import * as jsonapi from 'utils/json-api-client'
import { boundMethod } from 'autobind-decorator'
//import * as presenters from 'core/store/presenters'

/**
 * AccountBalances returns the account balances of STACKS & STX-LINK.
 *
 * @example "<application>/keys/eth"
 */
export const ACCOUNT_BALANCES_ENDPOINT =
  'https://stacks-node-api.mainnet.stacks.co/extended/v1/address/SP4VX529NN4XTFCR7WJ2VXB4GNGR2R7KHSN29RSK/balances'

export class StacksBalances {
  constructor(private api: jsonapi.Api) {}

  /**
   * Get account balances in STACKS and STX-LINK
   */
  @boundMethod
  public async getAccountBalances(): Promise<jsonapi.ApiResponse<any>> {
    console.log('stacks: ', await this.accountBalances())
    return this.accountBalances()
  }

  // privateresponse =  fetch(
  //     'https://stacks-node-api.mainnet.stacks.co/extended/v1/address/SP4VX529NN4XTFCR7WJ2VXB4GNGR2R7KHSN29RSK/balances',
  //     {
  //       method: 'GET',
  //     },
  //   )
  private accountBalances = this.api.fetchResource<{}, any, {}>(
    ACCOUNT_BALANCES_ENDPOINT,
    undefined,
    true,
  )
}

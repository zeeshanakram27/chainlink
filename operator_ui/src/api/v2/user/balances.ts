import * as jsonapi from 'utils/json-api-client'
import { boundMethod } from 'autobind-decorator'
import * as presenters from 'core/store/presenters'

/**
 * AccountBalances returns the account balances of ETH & LINK.
 *
 * @example "<application>/keys/eth"
 */
export const ACCOUNT_BALANCES_ENDPOINT = '/v2/keys/eth'

export class Balances {
  constructor(private api: jsonapi.Api) {}

  /**
   * Get account balances in ETH and LINK
   */
  @boundMethod
  public async getAccountBalances(): Promise<
    jsonapi.ApiResponse<presenters.AccountBalance[]>
  > {
    // const response = await fetch(
    //   'https://stacks-node-api.mainnet.stacks.co/extended/v1/address/SP4VX529NN4XTFCR7WJ2VXB4GNGR2R7KHSN29RSK/balances',
    //   {
    //     method: 'GET',
    //   },
    // )
    // console.log(response.json())
    console.log('here it is')
    const p = this.accountBalances()
    console.log('here goes p ', p)
    return this.accountBalances()
  }

  private accountBalances = this.api.fetchResource<
    {},
    presenters.AccountBalance[],
    {}
  >(ACCOUNT_BALANCES_ENDPOINT)
}

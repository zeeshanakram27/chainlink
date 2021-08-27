import * as jsonapi from 'utils/json-api-client'
import { boundMethod } from 'autobind-decorator'
import * as presenters from 'core/store/presenters'

/**
 * AccountBalances returns the account balances of STX & STX-LINK.
 *
 * @example "<application>/keys/stx"
 */
export const STACKS_ACCOUNT_BALANCES_ENDPOINT = '/v2/keys/stx'

export class StacksBalances {
  constructor(private api: jsonapi.Api) {}

  /**
   * Get account balances in STX and STX-LINK
   */
  @boundMethod
  public getAccountBalances(): Promise<
    jsonapi.ApiResponse<presenters.StacksAccountBalance[]>
  > {
    return this.accountBalances()
  }

  private accountBalances = this.api.fetchResource<
    {},
    presenters.StacksAccountBalance[],
    {}
  >(STACKS_ACCOUNT_BALANCES_ENDPOINT)
}

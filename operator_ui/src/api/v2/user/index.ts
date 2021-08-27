import { Api } from 'utils/json-api-client'
import { Balances } from './balances'
import { StacksBalances } from './stacksBalances'

export class User {
  constructor(private api: Api) {}

  public balances = new Balances(this.api)
  public stacksBalances = new StacksBalances(this.api)
}

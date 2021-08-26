import { Tooltip } from 'components/Tooltip'
import { PaddedCard } from 'components/PaddedCard'
import Typography from '@material-ui/core/Typography'
import { BigNumber } from 'bignumber.js'
import numeral from 'numeral'
import React, { FC } from 'react'

const WEI_PER_TOKEN = new BigNumber(10 ** 18)

const formatBalance = (val: string) => {
  const b = new BigNumber(val)
  const tokenBalance = b.dividedBy(WEI_PER_TOKEN).toNumber()
  return {
    formatted: numeral(tokenBalance).format('0.200000a'),
    unformatted: tokenBalance,
  }
}

const valAndTooltip = ({ value, error }: OwnProps) => {
  let val: string
  let tooltip: string

  if (error) {
    val = error
    tooltip = 'Error'
  } else if (value == null) {
    val = '...'
    tooltip = 'Loading...'
  } else {
    console.log('value before formatting', value)
    const balance = formatBalance(value)
    console.log('value after formatting', value)
    val = balance.formatted
    tooltip = balance.unformatted.toString()
  }

  return { val, tooltip }
}

// CHECKME
interface OwnProps {
  title: string
  value?: string
  stxValue?: string
  error?: string
}

const TokenBalance: FC<OwnProps> = (props) => {
  const { val, tooltip } = valAndTooltip(props)
  console.log(`this is val ${val} & tooltip ${tooltip}`)
  const { stxValue } = props
  return (
    <PaddedCard>
      <Typography variant="h5" color="secondary">
        {props.title}
      </Typography>
      <Typography variant="body1" color="textSecondary">
        <Tooltip title={tooltip}>
          <span>{stxValue ? stxValue : val}</span>
        </Tooltip>
      </Typography>
    </PaddedCard>
  )
}

export default TokenBalance

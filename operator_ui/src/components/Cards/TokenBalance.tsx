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

const formatStacksBalance = (val: string) => {
  // const b = new BigNumber(val)
  // const tokenBalance = b.dividedBy(1000000).toNumber()
  //  const tokenBalance = b.toNumber()
  return {
    formatted: val,
    unformatted: val,
  }
}

const valAndTooltip = ({ value, stxValue, error }: OwnProps) => {
  let val: string
  let stxVal: string
  let tooltip: string
  let stxTooltip: string

  if (error) {
    val = error
    stxVal = error
    tooltip = 'Error'
    stxTooltip = 'Error'
  } else if (value == null && stxValue == null) {
    val = '...'
    tooltip = 'Loading...'
    stxVal = '...'
    stxTooltip = 'Loading...'
  } else {
    const balance = formatBalance(value ? value : '0')
    const stxBalance = formatStacksBalance(stxValue ? stxValue : '0')
    val = balance.formatted
    stxVal = stxBalance.formatted
    tooltip = balance.unformatted.toString()
    stxTooltip = stxBalance.unformatted.toString()
  }

  return { val, stxVal, tooltip, stxTooltip }
}

// CHECKME
interface OwnProps {
  title: string
  value?: string
  stxValue?: string
  error?: string
}

const TokenBalance: FC<OwnProps> = (props) => {
  const { val, stxVal, tooltip, stxTooltip } = valAndTooltip(props)
  
  return (
    <PaddedCard>
      <Typography variant="h5" color="secondary">
        {props.title}
      </Typography>
      <Typography variant="body1" color="textSecondary">
        <Tooltip title={stxVal ? stxTooltip : tooltip}>
          <span>{stxVal ? stxVal : val}</span>
        </Tooltip>
      </Typography>
    </PaddedCard>
  )
}

export default TokenBalance

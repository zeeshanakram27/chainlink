package bulletprooftxmanager

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func newDynamicFeeAttempt(ks KeyStore, chainID *big.Int, etx models.EthTx, gasTipCap, gasFeeCap *big.Int, gasLimit uint64) (models.EthTxAttempt, error) {
	d := newDynamicFeeTransaction(
		uint64(*etx.Nonce),
		etx.ToAddress,
		etx.Value.ToInt(),
		gasLimit,
		chainID,
		gasTipCap,
		gasFeeCap,
		etx.EncodedPayload,
		types.AccessList{},
	)
	tx := types.NewTx(&d)
	attempt, err := newSignedAttempt(ks, etx, chainID, tx)
	if err != nil {
		return attempt, err
	}
	attempt.GasTipCap = utils.NewBig(gasTipCap)
	attempt.GasFeeCap = utils.NewBig(gasFeeCap)
	return attempt, nil
}

func newLegacyAttempt(ks KeyStore, chainID *big.Int, etx models.EthTx, gasPrice *big.Int, gasLimit uint64) (models.EthTxAttempt, error) {
	l := newLegacyTransaction(
		uint64(*etx.Nonce),
		etx.ToAddress,
		etx.Value.ToInt(),
		gasLimit,
		gasPrice,
		etx.EncodedPayload,
	)

	tx := types.NewTx(&l)
	attempt, err := newSignedAttempt(ks, etx, chainID, tx)
	if err != nil {
		return attempt, err
	}
	attempt.GasPrice = utils.NewBig(gasPrice)
	return attempt, nil
}

func newSignedAttempt(ks KeyStore, etx models.EthTx, chainID *big.Int, tx *types.Transaction) (attempt models.EthTxAttempt, err error) {
	hash, signedTxBytes, err := signTx(ks, etx.FromAddress, tx, chainID)
	if err != nil {
		return attempt, errors.Wrapf(err, "error using account %s to sign transaction %v", etx.FromAddress.String(), etx.ID)
	}

	attempt.State = models.EthTxAttemptInProgress
	attempt.SignedRawTx = signedTxBytes
	attempt.EthTxID = etx.ID
	attempt.Hash = hash

	return attempt, nil
}

func newLegacyTransaction(nonce uint64, to common.Address, value *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) types.LegacyTx {
	return types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	}
}

func newDynamicFeeTransaction(nonce uint64, to common.Address, value *big.Int, gasLimit uint64, chainID, gasTipCap, gasFeeCap *big.Int, data []byte, accessList types.AccessList) types.DynamicFeeTx {
	return types.DynamicFeeTx{
		ChainID:    chainID,
		Nonce:      nonce,
		GasTipCap:  gasTipCap,
		GasFeeCap:  gasFeeCap,
		Gas:        gasLimit,
		To:         &to,
		Value:      value,
		Data:       data,
		AccessList: accessList,
	}
}

func signTx(keyStore KeyStore, address common.Address, tx *types.Transaction, chainID *big.Int) (common.Hash, []byte, error) {
	signedTx, err := keyStore.SignTx(address, tx, chainID)
	if err != nil {
		return common.Hash{}, nil, errors.Wrap(err, "signTx failed")
	}
	rlp := new(bytes.Buffer)
	if err := signedTx.EncodeRLP(rlp); err != nil {
		return common.Hash{}, nil, errors.Wrap(err, "signTx failed")
	}
	return signedTx.Hash(), rlp.Bytes(), nil

}

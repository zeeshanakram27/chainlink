package adapters

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	strpkg "github.com/smartcontractkit/chainlink/core/store"
	"github.com/smartcontractkit/chainlink/core/store/models"
)

const (
	// DataFormatBytes instructs the EthTx Adapter to treat the input value as a
	// bytes string, rather than a hexadecimal encoded bytes32
	StacksDataFormatBytes = "bytes"
)

// EthTx holds the Address to send the result to and the FunctionSelector
// to execute.
type StxTx struct {
	ToAddress common.Address `json:"address"`
	// NOTE: FromAddress is deprecated and kept for backwards compatibility, new job specs should use fromAddresses
	FromAddress      common.Address          `json:"fromAddress,omitempty"`
	FromAddresses    []common.Address        `json:"fromAddresses,omitempty"`
	FunctionSelector models.FunctionSelector `json:"functionSelector"`
	// DataPrefix is typically a standard first argument
	// to chainlink callback calls - usually the requestID
	DataPrefix hexutil.Bytes `json:"dataPrefix"`
	DataFormat string        `json:"format"`
	GasLimit   uint64        `json:"gasLimit,omitempty"`

	// Optional list of desired encodings for ResultCollectKey arguments.
	// i.e. ["uint256", "bytes32"]
	ABIEncoding []string `json:"abiEncoding"`

	// MinRequiredOutgoingConfirmations only works with bulletprooftxmanager
	MinRequiredOutgoingConfirmations uint64 `json:"minRequiredOutgoingConfirmations,omitempty"`
}

// TODO: Would need to implement TaskType() and Perform() to complete the minimum requirement of core adapters.
// TaskType returns the type of Adapter.
func (e *StxTx) TaskType() models.TaskType {
	return TaskTypeEthTx
}

func (e *StxTx) Perform(input models.RunInput, store *strpkg.Store, keyStore *keystore.Master) models.RunOutput {
	data, err := models.JSON{}.Add("result", input.Data().String())
	if err != nil {
		return models.NewRunOutputError(err)
	}

	jp := JSONParse{}
	input = input.CloneWithData(data)
	return jp.Perform(input, store, keyStore)
}

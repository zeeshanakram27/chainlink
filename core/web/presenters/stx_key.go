package presenters

import (
	"time"

	"github.com/smartcontractkit/chainlink/core/assets"
)

// ETHKeyResource represents a ETH key JSONAPI resource. It holds the hex
// representation of the address plus its ETH & LINK balances
type STXKeyResource struct {
	JAID
	Address        string       `json:"address"`
	StxBalance     *assets.Eth  `json:"stxBalance"`
	StxLinkBalance *assets.Link `json:"stxLinkBalance"`
	NextNonce      int64        `json:"nextNonce"`
	IsFunding      bool         `json:"isFunding"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
	DeletedAt      *time.Time   `json:"deletedAt"`
}

// GetName implements the api2go EntityNamer interface
//
// This is named as such for backwards compatibility with the operator ui
// TODO - Standardise this to stxKeys
func (r STXKeyResource) GetName() string {
	return "sTXKeys"
}

// NewETHKeyOption defines a functional option which allows customisation of the
// EthKeyResource
type NewSTXKeyOption func(*STXKeyResource) error

// NewETHKeyResource constructs a new ETHKeyResource from a Key.
//
// Use the functional options to inject the ETH and LINK balances
func NewSTXKeyResource(k string, opts ...NewSTXKeyOption) (*STXKeyResource, error) {
	r := &STXKeyResource{
		JAID:           NewJAID(k),
		Address:        k,
		StxBalance:     nil,
		StxLinkBalance: nil,
	}

	for _, opt := range opts {
		err := opt(r)

		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

func SetSTXKeyStxBalance(ethBalance *assets.Eth) NewSTXKeyOption {
	return func(r *STXKeyResource) error {
		r.StxBalance = ethBalance

		return nil
	}
}

func SetSTXKeyStxLinkBalance(linkBalance *assets.Link) NewSTXKeyOption {
	return func(r *STXKeyResource) error {
		r.StxLinkBalance = linkBalance

		return nil
	}
}

package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/web/presenters"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type stxBalanceResponse struct {
	Stx struct {
		Balance                   string `json:"balance"`
		TotalSent                 string `json:"total_sent"`
		TotalReceived             string `json:"total_received"`
		TotalFeesSent             string `json:"total_fees_sent"`
		TotalMinerRewardsReceived string `json:"total_miner_rewards_received"`
		LockTxID                  string `json:"lock_tx_id"`
		Locked                    string `json:"locked"`
		LockHeight                int    `json:"lock_height"`
		BurnchainLockHeight       int    `json:"burnchain_lock_height"`
		BurnchainUnlockHeight     int    `json:"burnchain_unlock_height"`
	} `json:"stx"`
	FungibleTokens interface {
	} `json:"fungible_tokens"`
	NonFungibleTokens interface {
	} `json:"non_fungible_tokens"`
}

// ETHKeysController manages account keys
type STXKeysController struct {
	App chainlink.Application
}

// Index returns the node's Ethereum keys and the account balances of ETH & LINK.
// Example:
//  "<application>/keys/stx"
func (skc *STXKeysController) Index(c *gin.Context) {
	stacksAddr := os.Getenv("STACKS_ACCOUNT_ADDRESS")
	stacksNodeURL := os.Getenv("STACKS_NODE_URL")
	resource, err := presenters.NewSTXKeyResource(stacksAddr,
		skc.setStxBalance(stacksAddr, stacksNodeURL),
	)
	if err != nil {
		jsonAPIError(c, http.StatusInternalServerError, err)
		return
	}

	jsonAPIResponse(c, resource, "keys")
}

func (ekc *STXKeysController) setStxBalance(stacksAddr string, stacksNodeURL string) presenters.NewSTXKeyOption {
	url := fmt.Sprintf("https://%s/extended/v1/address/%s/balances", stacksNodeURL, stacksAddr)
	method := "GET"

	return func(r *presenters.STXKeyResource) error {
		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return errors.Errorf("error creating STX balance request: %v", err)
		}

		res, err := client.Do(req)
		if err != nil || res.StatusCode != http.StatusOK {
			return errors.Errorf("error getting STX balance response: %v", err)
		}
		defer res.Body.Close()

		Body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.Errorf("error reading STX balance response: %v", err)
		}

		var sbr stxBalanceResponse
		err = json.Unmarshal(Body, &sbr)
		if err != nil {
			return errors.Errorf("error getting stacks balance from Stacks node: %v", err)
		}
		bal := new(big.Int)
		bal, ok := bal.SetString(sbr.Stx.Balance, 10)
		if !ok {
			return errors.Errorf("error parsing stacks balance to big.Int: %v", err)

		}
		r.StxBalance = (*assets.Eth)(bal)
		return nil
	}
}

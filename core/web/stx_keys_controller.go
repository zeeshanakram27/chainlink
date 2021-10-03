package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/web/presenters"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const STX_LINK_IDENTIFIER = "::stxlink-token"

var (
	STACKS_ACCOUNT_ADDRESS = os.Getenv("STACKS_ACCOUNT_ADDRESS")
	STACKS_NODE_URL        = os.Getenv("STACKS_NODE_URL")
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
	FungibleTokens map[string]struct {
		Balance       string `json:"balance"`
		TotalSent     string `json:"total_sent"`
		TotalReceived string `json:"total_received"`
	} `json:"fungible_tokens"`
	NonFungibleTokens interface {
	} `json:"non_fungible_tokens"`
}

// STXKeysController manages account keys
type STXKeysController struct {
	App chainlink.Application
}

// Index returns the node's Stacks keys and the account balances of STX & STX-LINK.
// Example:
//  "<application>/keys/stx"
func (skc *STXKeysController) Index(c *gin.Context) {

	resource, err := presenters.NewSTXKeyResource(STACKS_ACCOUNT_ADDRESS,
		skc.setBalances(STACKS_ACCOUNT_ADDRESS, STACKS_NODE_URL),
	)
	if err != nil {
		jsonAPIError(c, http.StatusInternalServerError, err)
		return
	}

	jsonAPIResponse(c, resource, "keys")
}

func (ekc *STXKeysController) setBalances(stacksAddr string, stacksNodeURL string) presenters.NewSTXKeyOption {
	url := fmt.Sprintf("%s/extended/v1/address/%s/balances", stacksNodeURL, stacksAddr)
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
		stxBal := new(big.Int)
		stxBal, ok := stxBal.SetString(sbr.Stx.Balance, 10)
		if !ok {
			return errors.Errorf("error parsing stx balance to big.Int: %v", err)

		}
		r.StxBalance = (*assets.Eth)(stxBal)
		fmt.Println("stxBal: ", stxBal, r.StxBalance)
		stxLinkBal := new(big.Int)
		stxLinkBal, ok = stxLinkBal.SetString(sbr.getFungibleTokenBalance(STX_LINK_IDENTIFIER), 10)
		if !ok {
			return errors.Errorf("error parsing stxLink balance to big.Int: %v", err)

		}
		r.StxLinkBalance = (*assets.Link)(stxLinkBal)
		fmt.Println("stxLinkBal: ", stxLinkBal, r.StxLinkBalance)
		return nil
	}
}

func (sbr stxBalanceResponse) getFungibleTokenBalance(name string) string {
	for identifier, info := range sbr.FungibleTokens {
		if strings.HasSuffix(identifier, name) {
			return info.Balance
		}

	}
	return "0"
}

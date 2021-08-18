package web

import (
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/web/presenters"

	"github.com/gin-gonic/gin"
)

type ChainsController struct {
	App chainlink.Application
}

func (cc *ChainsController) Index(c *gin.Context, size, page, offset int) {
	var count int
	var err error

	var resources []presenters.ChainResource
	for _, chain := range chains {
		resources = append(resources, *presenters.NewChainResource(chain))
	}

	paginatedResponse(c, "chain", size, page, chain, count, err)
}

func (cc *ChainsController) Create(c *gin.Context) {

}

func (cc *ChainsController) Delete(c *gin.Context) {

}
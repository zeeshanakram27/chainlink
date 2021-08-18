package web

import (
	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/web/presenters"

	"github.com/gin-gonic/gin"
)

type NodesController struct {
	App chainlink.Application
}

func (nc *NodesController) Index(c *gin.Context, size, page, offset int) {
	id := c.Param("ID")

	var count int
	var err error

	if id == "" {
		// fetch nodes for chain ID
	} else {
		// fetch all nodes
	}

	var resources []presenters.ChainResource
	for _, node := range nodes {
		resources = append(resources, *presenters.NewChainResource(node))
	}

	paginatedResponse(c, "node", size, page, node, count, err)
}

func (nc *NodesController) Create(c *gin.Context) {

}

func (nc *NodesController) Delete(c *gin.Context) {

}
package cmd

import (
	"errors"

	"github.com/smartcontractkit/chainlink/core/web/presenters"
	"github.com/urfave/cli"
	"go.uber.org/multierr"
)

type NodePresenter struct {
	presenters.NodeResource
}

func (p *NodePresenter) ToRow() []string {
	row := []string{
		p.GetID(),
		p.Name,
		string(p.ChainID),
		p.WSURL.String(),
		p.HTTPURL.String(),
		p.CreatedAt.String(),
		p.UpdatedAt.String(),
	}
	return row
}

type NodePresenters []NodePresenter

// RenderTable implements TableRenderer
func (ps NodePresenters) RenderTable(rt RendererTable) error {
	headers := []string{"ID", "Name", "Chain ID", "Websocket URL", "Created", "Updated"}
	rows := [][]string{}

	for _, p := range ps {
		rows = append(rows, p.ToRow())
	}

	renderList(headers, rows, rt.Writer)

	return nil
}

// IndexNodes returns all nodes.
func (cli *Client) IndexNodes(c *cli.Context) (err error) {
	return cli.getPage("/v2/nodes", c.Int("page"), &NodePresenters{})
}

// ShowNode returns the info for the given Node name.
func (cli *Client) ShowNode(c *cli.Context) (err error) {
	if !c.Args().Present() {
		return cli.errorOut(errors.New("must pass the id of the node to be shown"))
	}
	nodeID := c.Args().First()
	resp, err := cli.HTTP.Get("/v2/nodes/" + nodeID)
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	return cli.renderAPIResponse(resp, &NodePresenter{})
}

// CreateNode adds a new node to the nodelink node
func (cli *Client) CreateNode(c *cli.Context) (err error) {
	if !c.Args().Present() {
		return cli.errorOut(errors.New("must pass in the node's parameters [JSON blob | JSON filepath]"))
	}

	buf, err := getBufferFromJSON(c.Args().First())
	if err != nil {
		return cli.errorOut(err)
	}

	resp, err := cli.HTTP.Post("/v2/nodes", buf)
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	return cli.renderAPIResponse(resp, &NodePresenter{})
}

// RemoveNode removes a specific Node by name.
func (cli *Client) RemoveNode(c *cli.Context) (err error) {
	if !c.Args().Present() {
		return cli.errorOut(errors.New("must pass the id of the node to be removed"))
	}
	nodeID := c.Args().First()
	resp, err := cli.HTTP.Delete("/v2/nodes/" + nodeID)
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	return cli.renderAPIResponse(resp, &NodePresenter{})
}

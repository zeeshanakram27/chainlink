package cmd

import (
	"errors"

	"github.com/smartcontractkit/chainlink/core/web/presenters"
	"github.com/urfave/cli"
	"go.uber.org/multierr"
)

type ChainPresenter struct {
	presenters.ChainResource
}

func (p *ChainPresenter) ToRow() []string {
	row := []string{
		p.GetID(),
		// p.Config,
		p.CreatedAt.String(),
		p.UpdatedAt.String(),
	}
	return row
}

type ChainPresenters []ChainPresenter

// RenderTable implements TableRenderer
func (ps ChainPresenters) RenderTable(rt RendererTable) error {
	headers := []string{"ID", "Config", "Created", "Updated"}
	rows := [][]string{}

	for _, p := range ps {
		rows = append(rows, p.ToRow())
	}

	renderList(headers, rows, rt.Writer)

	return nil
}

// IndexChains returns all chains.
func (cli *Client) IndexChains(c *cli.Context) (err error) {
	return cli.getPage("/v2/chains/evm", c.Int("page"), &ChainPresenters{})
}

// ShowChain returns the info for the given Chain name.
// func (cli *Client) ShowChain(c *cli.Context) (err error) {
// 	if !c.Args().Present() {
// 		return cli.errorOut(errors.New("must pass the id of the chain to be shown"))
// 	}
// 	chainID := c.Args().First()
// 	resp, err := cli.HTTP.Get("/v2/chains/evm/" + chainID)
// 	if err != nil {
// 		return cli.errorOut(err)
// 	}
// 	defer func() {
// 		if cerr := resp.Body.Close(); cerr != nil {
// 			err = multierr.Append(err, cerr)
// 		}
// 	}()

// 	return cli.renderAPIResponse(resp, &ChainPresenter{})
// }

// CreateChain adds a new chain to the chainlink node
func (cli *Client) CreateChain(c *cli.Context) (err error) {
	if !c.Args().Present() {
		return cli.errorOut(errors.New("must pass in the chain's parameters [JSON blob | JSON filepath]"))
	}

	buf, err := getBufferFromJSON(c.Args().First())
	if err != nil {
		return cli.errorOut(err)
	}

	resp, err := cli.HTTP.Post("/v2/chains/evm", buf)
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	return cli.renderAPIResponse(resp, &ChainPresenter{})
}

// RemoveChain removes a specific Chain by name.
func (cli *Client) RemoveChain(c *cli.Context) (err error) {
	if !c.Args().Present() {
		return cli.errorOut(errors.New("must pass the id of the chain to be removed"))
	}
	chainID := c.Args().First()
	resp, err := cli.HTTP.Delete("/v2/chains/evm/" + chainID)
	if err != nil {
		return cli.errorOut(err)
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			err = multierr.Append(err, cerr)
		}
	}()

	return cli.renderAPIResponse(resp, &ChainPresenter{})
}

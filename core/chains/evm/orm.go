package evm

import (
	"github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/smartcontractkit/sqlx"
	null "gopkg.in/guregu/null.v4"
)

type ORM interface {
	CreateChain(id utils.Big, config types.ChainCfg) (types.Chain, error)
	DeleteChain(id utils.Big) error
	Chains(offset, limit int) ([]types.Chain, int, error)
	CreateNode(data NewNode) (types.Node, error)
	DeleteNode(id int64) error
	Nodes(offset, limit int) ([]types.Node, int, error)
}

type orm struct {
	db *sqlx.DB
}

var _ ORM = (*orm)(nil)

func NewORM(db *sqlx.DB) ORM {
	return &orm{db}
}
func (o *orm) CreateChain(id utils.Big, config types.ChainCfg) (chain types.Chain, err error) {
	sql := `INSERT INTO evm_chains (id, cfg, created_at, updated_at) VALUES ($1, $2, now(), now()) RETURNING *`
	err = o.db.Get(&chain, sql, id, config)
	return chain, err
}
func (o *orm) DeleteChain(id utils.Big) error {
	sql := `DELETE FROM evm_chains WHERE id = $1`
	_, err := o.db.Exec(sql, id)
	// TODO: check result.RowsAffected?
	return err
}
func (o *orm) Chains(offset, limit int) ([]types.Chain, int, error) {
	return nil, 0, nil
}

type NewNode struct {
	Name       string      `json:"name"`
	EVMChainID utils.Big   `json:"evm_chain_id"`
	WSURL      null.String `json:"ws_url" db:"ws_url"`
	HTTPURL    string      `json:"http_url" db:"http_url"`
	SendOnly   bool        `json:"send_only"`
}

func (o *orm) CreateNode(data NewNode) (node types.Node, err error) {
	sql := `INSERT INTO nodes (name, evm_chain_id, ws_url, http_url, send_only, created_at, updated_at)
	VALUES (:name, :evm_chain_id, :ws_url, :http_url, :send_only, now(), now())
	RETURNING *;`
	stmt, err := o.db.PrepareNamed(sql)
	if err != nil {
		return node, err
	}
	err = stmt.Get(&node, data)
	return node, err
}

func (o *orm) DeleteNode(id int64) error {
	sql := `DELETE FROM nodes WHERE id = $1`
	_, err := o.db.Exec(sql, id)
	// TODO: check result.RowsAffected?
	return err
}
func (o *orm) Nodes(offset, limit int) ([]types.Node, int, error) {
	return nil, 0, nil
}

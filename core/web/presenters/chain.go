package presenters

import (
	"net/url"
	"time"
)

type ChainResource struct {
	JAID
	Config    map[string]interface{} `json:"config"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

// GetName implements the api2go EntityNamer interface
func (r ChainResource) GetName() string {
	return "chain"
}

type NodeResource struct {
	JAID
	Name      string    `json:"name"`
	ChainID   uint      `json:"chainID"`
	WSURL     *url.URL  `json:"wsURL"`
	HTTPURL   *url.URL  `json:"httpURL"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetName implements the api2go EntityNamer interface
func (r NodeResource) GetName() string {
	return "node"
}

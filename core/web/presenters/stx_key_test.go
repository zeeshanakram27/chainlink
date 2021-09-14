package presenters

import (
	"fmt"
	"testing"

	"github.com/manyminds/api2go/jsonapi"
	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSTXKeyResource(t *testing.T) {
	key := "ST1PQHQKV0RJXZFY1DGX8MNSNYVE3VGZJSRTPGZGM"
	r, err := NewSTXKeyResource(key,
		SetSTXKeyStxBalance(assets.NewEth(1)),
		SetSTXKeyStxLinkBalance(assets.NewLink(1)),
	)
	require.NoError(t, err)

	assert.Equal(t, assets.NewEth(1), r.StxBalance)
	assert.Equal(t, assets.NewLink(1), r.StxLinkBalance)

	b, err := jsonapi.Marshal(r)
	require.NoError(t, err)

	expected := fmt.Sprintf(`
	{
		"data":{
		   "type":"sTXKeys",
		   "id":"%s",
		   "attributes":{
			  "address":"%s",
			  "stxBalance":"1",
			  "stxLinkBalance":"1",
			  "nextNonce":0,
			  "isFunding":false,
			  "createdAt":"0001-01-01T00:00:00Z",
			  "updatedAt":"0001-01-01T00:00:00Z",
			  "deletedAt":null
		   }
		}
	 }
	`, key, key)

	assert.JSONEq(t, expected, string(b))

}

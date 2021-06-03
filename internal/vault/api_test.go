package vault

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	vaultAddress = "https://example.com"
	vaultToken   = "s.SecretToken"
)

func TestNewAPI(t *testing.T) {
	_ = os.Setenv("VAULT_ADDR", vaultAddress)
	_ = os.Setenv("VAULT_TOKEN", vaultToken)

	api, err := NewAPI()

	assert.NoError(t, err)

	assert.Equal(t, api.Config.Address, vaultAddress)
	assert.Equal(t, api.Client.Token(), vaultToken)
}

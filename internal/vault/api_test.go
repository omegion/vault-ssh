package vault

import (
	"os"
	"testing"

	hashivault "github.com/hashicorp/vault/vault"

	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/builtin/logical/ssh"
	"github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	vaultAddress = "https://example.com"
	vaultToken   = "s.SecretToken"
)

type VaultAPITestSuite struct {
	suite.Suite
	Cluster  *hashivault.TestCluster
	VaultAPI API
}

func (suite *VaultAPITestSuite) SetupSuite() {
	suite.T().Helper()

	coreConfig := &hashivault.CoreConfig{
		LogicalBackends: map[string]logical.Factory{
			"ssh": ssh.Factory,
		},
	}

	suite.Cluster = hashivault.NewTestCluster(suite.T(), coreConfig, &hashivault.TestClusterOptions{
		HandlerFunc: http.Handler,
	})

	suite.Cluster.Start()

	suite.VaultAPI = API{
		Client: suite.Cluster.Cores[0].Client,
	}
}

func (suite *VaultAPITestSuite) TearDownSuite() {
	suite.Cluster.Cleanup()
}

func (suite *VaultAPITestSuite) TestNewAPI() {
	_ = os.Setenv("VAULT_ADDR", vaultAddress)
	_ = os.Setenv("VAULT_TOKEN", vaultToken)

	vAPI, err := NewAPI()

	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), vAPI.Config.Address, vaultAddress)
	assert.Equal(suite.T(), vAPI.Client.Token(), vaultToken)
}

func TestVaultAPITestSuite(t *testing.T) {
	suite.Run(t, new(VaultAPITestSuite))
}

func (suite *VaultAPITestSuite) TestMount() {
	options := api.MountInput{Type: "ssh"}

	err := suite.VaultAPI.Mount("ssh", &options)

	assert.NoError(suite.T(), err)
}

func (suite *VaultAPITestSuite) Test_1_Write() {
	data := map[string]interface{}{
		"test_data": true,
	}

	_, err := suite.VaultAPI.Write("cubbyhole/test", data)

	assert.NoError(suite.T(), err)
}

func (suite *VaultAPITestSuite) Test_2_Read() {
	secret, err := suite.VaultAPI.Read("cubbyhole/test")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), true, secret.Data["test_data"])
}

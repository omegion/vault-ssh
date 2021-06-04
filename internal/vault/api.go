package vault

import (
	"os"
	"time"

	"github.com/hashicorp/vault/api"
)

const (
	tlsSkipVerify = true
	maxRetries    = 3
	timeout       = 15 * time.Second
	// SSHEngineDefaultName Default engine name.
	SSHEngineDefaultName = "ssh-test"
	// SSHEngineDefaultRoleName Default role name.
	SSHEngineDefaultRoleName = "general"
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/api_mock.go -package=mocks github.com/omegion/vault-ssh/internal/vault APIInterface
// APIInterface is an interface for Vault API.
type APIInterface interface {
	Mount(path string, input *api.MountInput) error
	Write(path string, options map[string]interface{}) (*api.Secret, error)
	Read(path string) (*api.Secret, error)
}

// API is main struct of Vault.
type API struct {
	Config *api.Config
	Client *api.Client
}

// NewAPI creates AI struct for Vault.
func NewAPI() (API, error) {
	config := api.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")
	config.MaxRetries = maxRetries
	config.Timeout = timeout

	err := config.ConfigureTLS(&api.TLSConfig{Insecure: tlsSkipVerify})
	if err != nil {
		return API{}, err
	}

	client, err := api.NewClient(config)
	if err != nil {
		return API{}, err
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	return API{
		Config: config,
		Client: client,
	}, nil
}

// Mount enables SSH engine with a path.
func (a API) Mount(path string, mountInfo *api.MountInput) error {
	return a.Client.Sys().Mount(path, mountInfo)
}

// Write creates CA certificate in ssh engine.
func (a API) Write(path string, data map[string]interface{}) (*api.Secret, error) {
	return a.Client.Logical().Write(path, data)
}

// Read gets CA certificate from ssh engine.
func (a API) Read(path string) (*api.Secret, error) {
	return a.Client.Logical().Read(path)
}

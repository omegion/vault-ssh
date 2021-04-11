package vault

import (
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/vault/api"
)

const (
	tlsSkipVerify            = true
	maxRetries               = 3
	timeout                  = 15 * time.Second
	SSHEngineDefaultName     = "ssh-test"
	SSHEngineDefaultRoleName = "general"
	SignedKeyDefaultPath     = "signed-key.pub"
)

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

// SealStatus returns status of seal.
func (a API) SealStatus() (api.SealStatusResponse, error) {
	status, err := a.Client.Sys().SealStatus()
	if err != nil {
		return api.SealStatusResponse{}, err
	}

	return *status, nil
}

// Unseal starts to unseal with given shard.
func (a API) Unseal(shard string) (api.SealStatusResponse, error) {
	status, err := a.Client.Sys().Unseal(shard)
	if err != nil {
		return api.SealStatusResponse{}, err
	}

	fmt.Printf("Unsealed with shard: %s\n", shard)

	return *status, nil
}

// EnableSSHEngine enables SSH engine with a path.
func (a API) EnableSSHEngine(engineName string) error {
	options := api.MountInput{Type: "ssh"}

	path := fmt.Sprintf("%s", engineName)

	err := a.Client.Sys().Mount(path, &options)
	if err != nil {
		return err
	}
	return nil
}

// CreateCACertificate creates CA certificate in ssh engine.
func (a API) CreateCACertificate(engineName string) error {
	options := map[string]interface{}{
		"generate_signing_key": true,
	}

	path := fmt.Sprintf("%s/config/ca", engineName)

	_, err := a.Client.Logical().Write(path, options)
	if err != nil {
		return err
	}

	return nil
}

// GetCACertificate gets CA certificate from ssh engine.
func (a API) GetCACertificate(engineName string) (string, error) {
	path := fmt.Sprintf("%s/config/ca", engineName)

	secret, err := a.Client.Logical().Read(path)
	if err != nil {
		return "", err
	}

	if publicKey, ok := secret.Data["public_key"]; ok {
		return fmt.Sprintf("%v", publicKey), nil
	}

	return "", nil
}

// CreateRole creates a role in ssh engine.
func (a API) CreateRole(engineName, roleName string) error {
	options := map[string]interface{}{
		"allow_user_certificates": true,
		"allowed_users":           "*",
		"allowed_extensions":      "permit-pty,permit-port-forwarding",
		"default_extensions": map[string]interface{}{
			"permit-pty": "",
		},
		"key_type":     "ca",
		"default_user": "root",
		"ttl":          "1m0s",
	}

	path := fmt.Sprintf("%s/roles/%s", engineName, roleName)

	_, err := a.Client.Logical().Write(path, options)
	if err != nil {
		return err
	}

	return nil
}

// Sign signs given public key with SSH engine and role.
func (a API) Sign(engineName, roleName, publicKey string) (string, error) {
	options := map[string]interface{}{
		"public_key": publicKey,
	}

	path := fmt.Sprintf("%s/sign/%s", engineName, roleName)

	secret, err := a.Client.Logical().Write(path, options)
	if err != nil {
		return "", err
	}

	if signedKey, ok := secret.Data["signed_key"]; ok {
		return fmt.Sprintf("%v", signedKey), nil
	}

	return "", nil
}

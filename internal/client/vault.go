package client

import (
	"github.com/omegion/vault-ssh/internal/controller"
	"github.com/omegion/vault-ssh/internal/vault"
)

// VaultInterface is an interface for Client.
type VaultInterface interface {
	EnableSSHEngine(path string) error
	CreateCACertificate(engineName string) error
	GetCACertificate(engineName string) (string, error)
	CreateRole(engineName, roleName string) error
	Sign(engineName, roleName string, publicKey []byte) (string, error)
}

// EnableSSHEngine enables SSH engine with a path.
func (c Client) EnableSSHEngine(path string) error {
	api, err := vault.NewAPI()
	if err != nil {
		return err
	}

	return controller.NewController(api).EnableSSHEngine(path)
}

// CreateCACertificate creates CA certificate in ssh engine.
func (c Client) CreateCACertificate(engineName string) error {
	api, err := vault.NewAPI()
	if err != nil {
		return err
	}

	return controller.NewController(api).CreateCACertificate(engineName)
}

// GetCACertificate gets CA certificate from ssh engine.
func (c Client) GetCACertificate(engineName string) (string, error) {
	api, err := vault.NewAPI()
	if err != nil {
		return "", err
	}

	return controller.NewController(api).GetCACertificate(engineName)
}

// CreateRole creates a role in ssh engine.
func (c Client) CreateRole(engineName, roleName string) error {
	api, err := vault.NewAPI()
	if err != nil {
		return err
	}

	return controller.NewController(api).CreateRole(engineName, roleName)
}

// Sign signs given public key with SSH engine and role.
func (c Client) Sign(engineName, roleName string, publicKey []byte) (string, error) {
	api, err := vault.NewAPI()
	if err != nil {
		return "", err
	}

	return controller.NewController(api).Sign(engineName, roleName, publicKey)
}

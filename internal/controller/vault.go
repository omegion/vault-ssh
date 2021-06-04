package controller

import (
	"fmt"

	"github.com/hashicorp/vault/api"
	"github.com/omegion/vault-ssh/internal/vault"
)

// Controller is main struct of Vault.
type Controller struct {
	API vault.APIInterface
}

// NewController is a factory to create a Controller.
func NewController(api vault.APIInterface) *Controller {
	return &Controller{API: api}
}

// EnableSSHEngine enables SSH engine with a path.
func (c Controller) EnableSSHEngine(path string) error {
	options := api.MountInput{Type: "ssh"}

	return c.API.Mount(path, &options)
}

// CreateCACertificate creates CA certificate in ssh engine.
func (c Controller) CreateCACertificate(engineName string) error {
	data := map[string]interface{}{
		"generate_signing_key": true,
	}

	_, err := c.API.Write(fmt.Sprintf("%s/config/ca", engineName), data)
	if err != nil {
		return err
	}

	return nil
}

// GetCACertificate gets CA certificate from ssh engine.
func (c Controller) GetCACertificate(engineName string) (string, error) {
	path := fmt.Sprintf("%s/config/ca", engineName)

	secret, err := c.API.Read(path)
	if err != nil {
		return "", err
	}

	if publicKey, ok := secret.Data["public_key"]; ok {
		return fmt.Sprintf("%v", publicKey), nil
	}

	return "", vault.SecretDataNotFound{
		Key:  "public_key",
		Path: path,
	}
}

// CreateRole creates a role in ssh engine.
func (c Controller) CreateRole(engineName, roleName string) error {
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

	_, err := c.API.Write(path, options)
	if err != nil {
		return err
	}

	return nil
}

// Sign signs given public key with SSH engine and role.
func (c Controller) Sign(engineName, roleName string, publicKey []byte) (string, error) {
	options := map[string]interface{}{
		"public_key": string(publicKey),
	}

	path := fmt.Sprintf("%s/sign/%s", engineName, roleName)

	secret, err := c.API.Write(path, options)
	if err != nil {
		return "", err
	}

	if signedKey, ok := secret.Data["signed_key"]; ok {
		return fmt.Sprintf("%v", signedKey), nil
	}

	return "", vault.SecretDataNotFound{
		Key:  "signed_key",
		Path: path,
	}
}

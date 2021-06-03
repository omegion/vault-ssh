package vault

import (
	"errors"
	"fmt"
	"testing"

	vaultApi "github.com/hashicorp/vault/api"

	"github.com/omegion/vault-ssh/internal/vault/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	vaultConfigPath = "test/config/ca"
	vaultSignPath   = "test/sign/test-role"
)

func TestNewController(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	controller := NewController(api)

	assert.Equal(t, api, controller.API)
}

func TestEnableSSHEngine(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	expectedPath := "test/path/engine"

	api.EXPECT().Mount(
		expectedPath,
		&vaultApi.MountInput{Type: "ssh"},
	).Return(nil)

	controller := NewController(api)

	err := controller.EnableSSHEngine(expectedPath)

	assert.NoError(t, err)
}

func TestCreateCACertificate(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	options := map[string]interface{}{
		"generate_signing_key": true,
	}

	secret := vaultApi.Secret{}

	api.EXPECT().Write(vaultConfigPath, options).Return(&secret, nil)

	controller := NewController(api)

	err := controller.CreateCACertificate("test")

	assert.NoError(t, err)
}

func TestCreateCACertificate_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	options := map[string]interface{}{
		"generate_signing_key": true,
	}

	secret := vaultApi.Secret{}

	api.EXPECT().Write(vaultConfigPath, options).Return(&secret, errors.New("custom-error"))

	controller := NewController(api)

	err := controller.CreateCACertificate("test")

	assert.EqualError(t, err, "custom-error")
}

func TestGetCACertificate(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	expectedPublicKey := "X"

	secret := vaultApi.Secret{Data: map[string]interface{}{"public_key": expectedPublicKey}}

	api.EXPECT().Read(vaultConfigPath).Return(&secret, nil)

	controller := NewController(api)

	publicKey, err := controller.GetCACertificate("test")

	assert.NoError(t, err)
	assert.Equal(t, expectedPublicKey, publicKey)
}

func TestGetCACertificate_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	api.EXPECT().Read(vaultConfigPath).Return(&vaultApi.Secret{}, errors.New("custom-error"))

	controller := NewController(api)

	_, err := controller.GetCACertificate("test")

	assert.EqualError(t, err, "custom-error")
}

func TestGetCACertificate_PublicKeyNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	api.EXPECT().Read(vaultConfigPath).Return(&vaultApi.Secret{}, nil)

	controller := NewController(api)

	_, err := controller.GetCACertificate("test")

	assert.EqualError(t, err, fmt.Sprintf("Vault secret not found with path %s, key public_key", vaultConfigPath))
}

func TestCreateRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	path := "test/roles/test-role"

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

	api.EXPECT().Write(path, options).Return(&vaultApi.Secret{}, nil)

	controller := NewController(api)

	err := controller.CreateRole("test", "test-role")

	assert.NoError(t, err)
}

func TestCreateRole_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	path := "test/roles/test-role"

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

	api.EXPECT().Write(path, options).Return(nil, errors.New("custom-error"))

	controller := NewController(api)

	err := controller.CreateRole("test", "test-role")

	assert.EqualError(t, err, "custom-error")
}

func TestSign(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	expectedPublicKey := "X"
	expectedSignedKey := "Y"

	options := map[string]interface{}{
		"public_key": expectedPublicKey,
	}

	secret := vaultApi.Secret{Data: map[string]interface{}{"signed_key": expectedSignedKey}}

	api.EXPECT().Write(vaultSignPath, options).Return(&secret, nil)

	controller := NewController(api)

	publicKey, err := controller.Sign("test", "test-role", expectedPublicKey)

	assert.NoError(t, err)
	assert.Equal(t, publicKey, expectedSignedKey)
}

func TestSign_Write_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	expectedPublicKey := "X"

	options := map[string]interface{}{
		"public_key": expectedPublicKey,
	}

	api.EXPECT().Write(vaultSignPath, options).Return(&vaultApi.Secret{}, errors.New("custom-error"))

	controller := NewController(api)

	_, err := controller.Sign("test", "test-role", expectedPublicKey)

	assert.EqualError(t, err, "custom-error")
}

func TestSign_SecretData_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	expectedPublicKey := "X"

	options := map[string]interface{}{
		"public_key": expectedPublicKey,
	}

	secret := vaultApi.Secret{Data: map[string]interface{}{}}

	api.EXPECT().Write(vaultSignPath, options).Return(&secret, nil)

	controller := NewController(api)

	_, err := controller.Sign("test", "test-role", expectedPublicKey)

	assert.EqualError(t, err, fmt.Sprintf("Vault secret not found with path %s, key signed_key", vaultSignPath))
}

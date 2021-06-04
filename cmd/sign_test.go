package cmd

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/omegion/vault-ssh/internal/client/mocks"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

const (
	expectedEngineName = "ssh"
	expectedRoleName   = "test"
)

func TestSignCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mocks.NewMockInterface(ctrl)

	expectedPublicKey := "X"
	expectedPublicKeyPath := "fixtures/public-key.pem"

	cmd := &cobra.Command{}
	cmd.Flags().String("engine", expectedEngineName, "")
	cmd.Flags().String("role", expectedRoleName, "")
	cmd.Flags().String("public-key", expectedPublicKeyPath, "")

	buffer, err := ioutil.ReadFile(expectedPublicKeyPath)
	assert.NoError(t, err)

	c.EXPECT().Sign(expectedEngineName, expectedRoleName, buffer).Return(expectedPublicKey, nil).Times(1)

	err = signE(c, cmd, []string{})

	assert.NoError(t, err)
}

func TestSignCommand_File_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mocks.NewMockInterface(ctrl)

	cmd := &cobra.Command{}
	cmd.Flags().String("engine", expectedEngineName, "")
	cmd.Flags().String("role", expectedRoleName, "")
	cmd.Flags().String("public-key", "", "")

	err := signE(c, cmd, []string{})

	assert.EqualError(t, err, "open : no such file or directory")
}

func TestSignCommand_Controller_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mocks.NewMockInterface(ctrl)

	expectedPublicKey := "X"
	expectedPublicKeyPath := "fixtures/public-key.pem"

	cmd := &cobra.Command{}
	cmd.Flags().String("engine", expectedEngineName, "")
	cmd.Flags().String("role", expectedRoleName, "")
	cmd.Flags().String("public-key", expectedPublicKeyPath, "")

	buffer, err := ioutil.ReadFile(expectedPublicKeyPath)
	assert.NoError(t, err)

	c.EXPECT().
		Sign(expectedEngineName, expectedRoleName, buffer).
		Return(expectedPublicKey, errors.New("custom-error")).
		Times(1)

	err = signE(c, cmd, []string{})

	assert.EqualError(t, err, "custom-error")
}

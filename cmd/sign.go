package cmd

import (
	"fmt"
	"github.com/omegion/vault-ssh/pkg/vault"
	"github.com/spf13/cobra"
	"io/ioutil"
)

// Sign signs given public key with SSH engine and role.
func Sign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign",
		Short: "Signs given public key with SSH engine and role.",
		RunE: func(cmd *cobra.Command, args []string) error {
			engineName, _ := cmd.Flags().GetString("engine")
			roleName, _ := cmd.Flags().GetString("role")
			publicKeyPath, _ := cmd.Flags().GetString("public-key")

			publicKey, err := readFile(publicKeyPath)
			if err != nil {
				return err
			}

			api, err := vault.NewAPI()
			if err != nil {
				return err
			}

			key, err := api.Sign(engineName, roleName, publicKey)
			if err != nil {
				return err
			}

			fmt.Println(key)

			return nil
		},
	}

	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")
	cmd.Flags().String("role", vault.SSHEngineDefaultRoleName, "Role name for SSH engine")
	cmd.Flags().String("public-key", vault.SSHEngineDefaultRoleName, "Public key to sign")

	return cmd
}

func readFile(path string) (string, error) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

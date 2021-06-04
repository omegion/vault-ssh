package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/omegion/vault-ssh/internal/controller"
	"github.com/omegion/vault-ssh/internal/vault"
	"github.com/spf13/cobra"
)

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")

	if err := cmd.MarkFlagRequired("engine"); err != nil {
		cobra.CheckErr(err)
	}

	cmd.Flags().String("role", vault.SSHEngineDefaultRoleName, "Role name for SSH engine")

	if err := cmd.MarkFlagRequired("role"); err != nil {
		cobra.CheckErr(err)
	}

	cmd.Flags().String("public-key", vault.SSHEngineDefaultRoleName, "Public key to sign")

	if err := cmd.MarkFlagRequired("public-key"); err != nil {
		cobra.CheckErr(err)
	}
}

// Sign signs given public key with SSH engine and role.
func Sign() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign",
		Short: "Signs given public key with SSH engine and role.",
		RunE: func(cmd *cobra.Command, args []string) error {
			engineName, _ := cmd.Flags().GetString("engine")
			roleName, _ := cmd.Flags().GetString("role")
			publicKeyPath, _ := cmd.Flags().GetString("public-key")

			buffer, err := ioutil.ReadFile(publicKeyPath)
			if err != nil {
				return err
			}

			api, err := vault.NewAPI()
			if err != nil {
				return err
			}

			publicKey, err := controller.NewController(api).Sign(engineName, roleName, buffer)
			if err != nil {
				return err
			}

			//nolint:forbidigo // make an exception.
			fmt.Println(publicKey)

			return nil
		},
	}

	setupGetCommand(cmd)

	return cmd
}

package cmd

import (
	"fmt"
	"github.com/omegion/vault-ssh/pkg/vault"
	"github.com/spf13/cobra"
)

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().String("path", vault.SSHEngineDefaultName, "SSH engine path")
}

// Enable enables SSH engine.
func Enable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enables SSH Engine.",
		RunE: func(cmd *cobra.Command, args []string) error {
			path, _ := cmd.Flags().GetString("path")

			api, err := vault.NewAPI()
			if err != nil {
				return err
			}

			err = api.EnableSSHEngine(path)
			if err != nil {
				return err
			}

			fmt.Printf("\"%s\" SSH Engine enabled.\n", path)

			return nil
		},
	}

	setupGetCommand(cmd)

	return cmd
}

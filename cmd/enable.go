package cmd

import (
	"fmt"

	"github.com/omegion/vault-ssh/internal/vault"

	"github.com/spf13/cobra"
)

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

			controller := vault.NewController(api)

			err = controller.EnableSSHEngine(path)
			if err != nil {
				return err
			}

			fmt.Printf("\"%s\" SSH Engine enabled.\n", path)

			return nil
		},
	}

	cmd.Flags().String("path", vault.SSHEngineDefaultName, "SSH engine path")

	return cmd
}

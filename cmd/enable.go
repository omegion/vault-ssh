package cmd

import (
	"github.com/omegion/vault-ssh/internal/controller"
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

			return controller.NewController(api).EnableSSHEngine(path)
		},
	}

	cmd.Flags().String("path", vault.SSHEngineDefaultName, "SSH engine path")

	return cmd
}

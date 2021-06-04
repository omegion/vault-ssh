package certificate

import (
	"github.com/omegion/vault-ssh/internal/controller"
	"github.com/omegion/vault-ssh/internal/vault"
	"github.com/spf13/cobra"
)

// Create creates CA certificate for SSH engine.
func Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates CA certificate for SSH engine.",
		RunE: func(cmd *cobra.Command, args []string) error {
			engineName, _ := cmd.Flags().GetString("engine")

			api, err := vault.NewAPI()
			if err != nil {
				return err
			}

			return controller.NewController(api).CreateCACertificate(engineName)
		},
	}

	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")

	return cmd
}

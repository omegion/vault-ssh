package certificate

import (
	"fmt"

	"github.com/omegion/vault-ssh/pkg/vault"

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

			err = api.CreateCACertificate(engineName)
			if err != nil {
				return err
			}

			fmt.Printf("Certificate created for SSH Engine \"%s\".\n", engineName)

			return nil
		},
	}

	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")

	return cmd
}

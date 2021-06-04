package certificate

import (
	"fmt"

	"github.com/omegion/vault-ssh/pkg/vault"

	"github.com/spf13/cobra"
)

// Get get CA certificate for SSH engine.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets CA certificate for SSH engine.",
		RunE: func(cmd *cobra.Command, args []string) error {
			engineName, _ := cmd.Flags().GetString("engine")

			api, err := vault.NewAPI()
			if err != nil {
				return err
			}

			publicKey, err := api.GetCACertificate(engineName)
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", publicKey)

			return nil
		},
	}

	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")

	return cmd
}

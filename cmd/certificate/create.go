package certificate

import (
	"github.com/omegion/vault-ssh/internal/client"
	"github.com/omegion/vault-ssh/internal/vault"

	"github.com/spf13/cobra"
)

// Create creates CA certificate for SSH engine.
func Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates CA certificate for SSH engine.",
		RunE:  client.With(enableSSHEngineE),
	}

	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")

	return cmd
}

func enableSSHEngineE(c client.Interface, cmd *cobra.Command, args []string) error {
	engineName, _ := cmd.Flags().GetString("engine")

	return c.CreateCACertificate(engineName)
}

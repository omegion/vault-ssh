package cmd

import (
	"fmt"

	"github.com/omegion/vault-ssh/internal/client"

	"github.com/omegion/vault-ssh/internal/vault"

	"github.com/spf13/cobra"
)

// Enable enables SSH engine.
func Enable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enables SSH Engine.",
		RunE:  client.With(enableSSHEngineE),
	}

	cmd.Flags().String("path", vault.SSHEngineDefaultName, "SSH engine path")

	return cmd
}

func enableSSHEngineE(c client.Interface, cmd *cobra.Command, args []string) error {
	path, _ := cmd.Flags().GetString("path")

	err := c.EnableSSHEngine(path)
	if err != nil {
		return err
	}

	fmt.Printf("\"%s\" SSH Engine enabled.\n", path)

	return nil
}

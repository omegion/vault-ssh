package cmd

import (
	"fmt"

	"github.com/omegion/vault-ssh/internal/client"
	"github.com/omegion/vault-ssh/internal/vault"
	log "github.com/sirupsen/logrus"
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

	if err := c.EnableSSHEngine(path); err != nil {
		return err
	}

	log.Infoln(fmt.Sprintf("\"%s\" SSH Engine enabled.", path))

	return nil
}

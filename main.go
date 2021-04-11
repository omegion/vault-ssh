package main

import (
	"github.com/omegion/vault-ssh/cmd/certificate"
	"github.com/omegion/vault-ssh/cmd/role"
	"os"

	"github.com/omegion/vault-ssh/cmd"

	"github.com/spf13/cobra"
)

// RootCommand is the main entry point of this application.
func RootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:          "vault-ssh",
		Short:        "Vault SSH Manager",
		Long:         "CLI command to manage SSH connections with Vault",
		SilenceUsage: true,
	}

	root.AddCommand(cmd.Version())
	root.AddCommand(cmd.Enable())
	root.AddCommand(cmd.Sign())
	root.AddCommand(certificate.Certificate())
	root.AddCommand(role.Role())

	return root
}

func main() {
	if err := RootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

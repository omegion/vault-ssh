package main

import (
	"os"

	commander "github.com/omegion/cobra-commander"
	"github.com/omegion/vault-ssh/cmd"
	"github.com/omegion/vault-ssh/cmd/certificate"
	"github.com/omegion/vault-ssh/cmd/role"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:          "vault-ssh",
		Short:        "Vault SSH Manager",
		Long:         "CLI command to manage SSH connections with Vault.",
		SilenceUsage: true,
	}

	c := commander.NewCommander(root).
		SetCommand(
			cmd.Version(),
			cmd.Enable(),
			cmd.Sign(),
			certificate.Certificate(),
			role.Role(),
		).
		Init()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}

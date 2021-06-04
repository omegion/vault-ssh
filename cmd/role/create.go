package role

import (
	"github.com/omegion/vault-ssh/internal/client"

	"github.com/omegion/vault-ssh/internal/vault"

	"github.com/spf13/cobra"
)

// Create creates a role for SSH engine.
func Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a role for SSH engine.",
		RunE:  client.With(createRoleE),
	}

	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")
	cmd.Flags().String("name", vault.SSHEngineDefaultRoleName, "Role name for SSH engine")

	return cmd
}

func createRoleE(c client.Interface, cmd *cobra.Command, args []string) error {
	engineName, _ := cmd.Flags().GetString("engine")
	roleName, _ := cmd.Flags().GetString("name")

	return c.CreateRole(engineName, roleName)
}

package role

import (
	"fmt"
	"github.com/omegion/vault-ssh/pkg/vault"
	"github.com/spf13/cobra"
)

// Create creates a role for SSH engine.
func Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a role for SSH engine.",
		RunE: func(cmd *cobra.Command, args []string) error {
			engineName, _ := cmd.Flags().GetString("engine")
			roleName, _ := cmd.Flags().GetString("name")

			api, err := vault.NewAPI()
			if err != nil {
				return err
			}

			err = api.CreateRole(engineName, roleName)
			if err != nil {
				return err
			}

			fmt.Printf("\"%s\" created for SSH Engine \"%s\" enabled.\n", roleName, engineName)

			return nil
		},
	}

	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")
	cmd.Flags().String("name", vault.SSHEngineDefaultRoleName, "Role name for SSH engine")

	return cmd
}

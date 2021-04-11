package role

import (
	"github.com/spf13/cobra"
)

// Role is main command for role commands.
func Role() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "role",
		Short: "Manages roles for SSH engine.",
	}

	cmd.AddCommand(Create())

	return cmd
}

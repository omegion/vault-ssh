package certificate

import (
	"github.com/spf13/cobra"
)

// Certificate is main command for certificate commands.
func Certificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "certificate",
		Short: "Manages certificates for SSH engine.",
	}

	cmd.AddCommand(Create())
	cmd.AddCommand(Get())

	return cmd
}

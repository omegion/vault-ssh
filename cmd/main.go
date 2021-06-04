package cmd

import (
	"github.com/omegion/vault-ssh/cmd/certificate"
	"github.com/omegion/vault-ssh/cmd/role"
	"github.com/omegion/vault-ssh/internal/client"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Commander is a struct for command system.
type Commander struct {
	Root     *cobra.Command
	LogLevel string
}

// NewCommander is a factory for Commander.
func NewCommander() *Commander {
	return &Commander{}
}

// SetRootCommand sets Root command.
func (c *Commander) SetRootCommand() {
	c.Root = &cobra.Command{
		Use:          "vault-ssh",
		Short:        "Vault SSH Manager",
		Long:         "CLI command to manage SSH connections with Vault.",
		SilenceUsage: true,
	}
}

func (c *Commander) setPersistentFlags() {
	c.Root.PersistentFlags().String("logLevel", "info", "Set the logging level. One of: debug|info|warn|error")
}

func (c *Commander) setLogger() {
	c.LogLevel, _ = c.Root.Flags().GetString("logLevel")

	level, err := log.ParseLevel(c.LogLevel)
	if err != nil {
		cobra.CheckErr(err)
	}

	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	})
}

// Setup is entrypoint for the commands.
func (c *Commander) Setup() {
	cobra.OnInitialize(func() {
		c.setLogger()
	})

	c.SetRootCommand()
	c.setPersistentFlags()

	c.Root.AddCommand(Version())
	c.Root.AddCommand(Enable())
	c.Root.AddCommand(Sign())
	c.Root.AddCommand(certificate.Certificate())
	c.Root.AddCommand(role.Role())
}

// WithClient is a wrapper for testing.
func (c *Commander) WithClient(
	fn func(c client.Interface, cmd *cobra.Command, args []string) error,
) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()

		return fn(c, cmd, args)
	}
}

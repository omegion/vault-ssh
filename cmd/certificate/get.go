package certificate

import (
	"github.com/omegion/vault-ssh/internal/client"
	"github.com/omegion/vault-ssh/internal/vault"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// Get get CA certificate for SSH engine.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Gets CA certificate for SSH engine.",
		RunE:  client.With(getCACertificateE),
	}

	cmd.Flags().String("engine", vault.SSHEngineDefaultName, "SSH engine path")

	return cmd
}

func getCACertificateE(c client.Interface, cmd *cobra.Command, args []string) error {
	engineName, _ := cmd.Flags().GetString("engine")

	publicKey, err := c.GetCACertificate(engineName)
	if err != nil {
		return err
	}

	log.Infoln(publicKey)

	return nil
}

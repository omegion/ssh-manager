package cmd

import (
	"fmt"

	"github.com/omegion/ssh-manager/internal/io"
	"github.com/omegion/ssh-manager/internal/provider"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// setupGetCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "Name")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("provider", "", "Provider")

	if err := cmd.MarkFlagRequired("provider"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().Bool("read-only", false, "Do not write fetched SSH keys")
}

// Get acquires SSH key from given provider.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get SSH key from given provider.",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			providerName, _ := cmd.Flags().GetString("provider")
			readOnly, _ := cmd.Flags().GetBool("read-only")

			commander := provider.NewCommander()

			prv, err := decideProvider(&providerName, &commander)
			if err != nil {
				return err
			}

			item, err := prv.Get(name)
			if err != nil {
				return err
			}

			log.Infoln(fmt.Sprintf("SSH Keys are fetched for %s.", name))

			for _, field := range item.Values {
				fileName := item.Name

				if field.Name == "public_key" {
					fileName = fmt.Sprintf("%s.pub", item.Name)
				}

				if readOnly {
					log.Infoln(fmt.Sprintf("%s\n%s", field.Name, field.Value))

					continue
				}

				err := io.WriteSSHKey(fileName, []byte(field.Value))
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	setupGetCommand(cmd)

	return cmd
}

func decideProvider(name *string, commander *provider.Commander) (provider.APIInterface, error) {
	switch *name {
	case provider.BitwardenCommand:
		return provider.Bitwarden{Commander: *commander}, nil
	case provider.OnePasswordCommand:
		return provider.OnePassword{Commander: *commander}, nil
	default:
		return provider.Bitwarden{}, provider.NotFound{Name: name}
	}
}

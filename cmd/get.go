package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/ssh-manager/internal/controller"
	"github.com/omegion/ssh-manager/internal/io"
	"github.com/omegion/ssh-manager/internal/provider"
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

	cmd.Flags().Bool("read-only", false, "Do not write fetched Manager keys")
	cmd.Flags().String("bucket", "", "S3 Bucket Name")
}

// Get acquires Manager key from given provider.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get Manager key from given provider.",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			providerName, _ := cmd.Flags().GetString("provider")
			readOnly, _ := cmd.Flags().GetBool("read-only")
			bucket, _ := cmd.Flags().GetString("bucket")

			options := provider.GetOptions{Name: name}

			if bucket != "" {
				options.Bucket = &bucket
			}

			item, err := controller.NewManager(&providerName).Get(options)
			if err != nil {
				return err
			}

			log.Infoln(fmt.Sprintf("Manager Keys are fetched for %s.", name))

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

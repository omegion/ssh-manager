package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/ssh-manager/internal/controller"
)

// setupListCommand sets default flags.
func setupListCommand(cmd *cobra.Command) {
	cmd.Flags().String("provider", "", "Provider")

	if err := cmd.MarkFlagRequired("provider"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

// List acquires Manager keys from given provider.
func List() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Manager keys from given provider.",
		RunE: func(cmd *cobra.Command, args []string) error {
			providerName, _ := cmd.Flags().GetString("provider")

			items, err := controller.NewManager(&providerName).List()
			if err != nil {
				return err
			}

			log.Infoln("Manager Keys are fetched.")

			for _, item := range items {
				log.Infoln(item.Name)
			}

			return nil
		},
	}

	setupListCommand(cmd)

	return cmd
}

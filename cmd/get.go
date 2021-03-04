package cmd

import (
	"github.com/omegion/lpass-ssh/pkg/lpass"
	"github.com/spf13/cobra"
	"log"
)

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "Name")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

// Get acquires SSH key from LastPass.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Add SSH key to LastPass.",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")

			sshKey := lpass.SSHKey{
				Name: name,
			}

			lastPass := lpass.NewLastPass()

			err := lastPass.Get(sshKey)
			if err != nil {
				return err
			}

			return nil
		},
	}

	setupGetCommand(cmd)

	return cmd
}

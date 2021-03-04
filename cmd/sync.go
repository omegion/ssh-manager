package cmd

import (
	"fmt"
	"github.com/omegion/lpass-ssh/pkg/lpass"

	"github.com/spf13/cobra"
)

// Sync installs all SSH keys from LastPass to local machine.
func Sync() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync SSH keys from LastPass.",
		RunE: func(cmd *cobra.Command, args []string) error {
			lastPass := lpass.NewLastPass()
			err := lastPass.Sync()
			if err != nil {
				return err
			}

			fmt.Printf("SSH Keys are synced.\n")

			return nil
		},
	}

	return cmd
}

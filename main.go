package main

import (
	"github.com/spf13/cobra"
	"os"

	commander "github.com/omegion/cobra-commander"

	"github.com/omegion/ssh-manager/cmd"
)

func main() {
	root := &cobra.Command{
		Use:          "vault-unseal",
		Short:        "Vault Auto Unseal",
		Long:         "CLI command to automatically unseal Vault",
		SilenceUsage: true,
	}

	c := commander.NewCommander(root).
		SetCommand(
			cmd.Version(),
			cmd.Get(),
			cmd.Add(),
			cmd.List(),
		).
		Init()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}

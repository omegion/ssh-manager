package main

import (
	"os"

	commander "github.com/omegion/cobra-commander"
	"github.com/spf13/cobra"

	"github.com/omegion/ssh-manager/cmd"
)

func main() {
	root := &cobra.Command{
		Use:          "vault-unseal",
		Short:        "Vault Auto Unseal",
		Long:         "CLI command to automatically unseal Vault",
		SilenceUsage: true,
	}

	comm := commander.NewCommander(root).
		SetCommand(
			cmd.Version(),
			cmd.Get(),
			cmd.Add(),
			cmd.List(),
		).
		Init()

	if err := comm.Execute(); err != nil {
		os.Exit(1)
	}
}

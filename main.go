package main

import (
	"os"

	commander "github.com/omegion/cobra-commander"
	"github.com/spf13/cobra"

	"github.com/omegion/ssh-manager/cmd"
)

func main() {
	root := &cobra.Command{
		Use:          "ssh-manager",
		Short:        "SSH Manager for your keys",
		Long:         "SSH Manager for your keys on 1Password and Bitwarden",
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

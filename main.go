package main

import (
	"github.com/omegion/bw-ssh/cmd"

	"github.com/spf13/cobra"
)

// RootCommand is the main entry point of this application.
func RootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:          "bw-ssh",
		Short:        "Bitwarden SSH Manager",
		Long:         "CLI command to manage SSH keys stored on Bitwarden",
		SilenceUsage: true,
	}

	root.AddCommand(cmd.Version())
	root.AddCommand(cmd.Add())
	root.AddCommand(cmd.Get())

	return root
}

func main() {
	root := RootCommand()
	_ = root.Execute()
}

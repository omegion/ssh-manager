package main

import (
	"github.com/omegion/lpass-ssh/cmd"

	"github.com/spf13/cobra"
)

// RootCommand is the main entry point of this application.
func RootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "lpass-ssh",
		Short: "LastPass SSH Manager",
		Long:  "CLI command to manage SSH keys stored on LastPass",
	}

	root.AddCommand(cmd.Version())
	root.AddCommand(cmd.Sync())
	root.AddCommand(cmd.Add())
	root.AddCommand(cmd.Get())

	return root
}

func main() {
	root := RootCommand()
	_ = root.Execute()
}

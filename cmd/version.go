package cmd

import (
	"fmt"

	"github.com/omegion/lpass-ssh/pkg/info"

	"github.com/spf13/cobra"
)

// Version prints version/build.
func Version() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version/build number",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%s %s\n", info.AppName, info.Version)

			return nil
		},
	}

	return cmd
}

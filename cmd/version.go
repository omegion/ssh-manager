package cmd

import (
	"github.com/go-asset/build"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// AppName is the name of the Application.
var AppName = "ssh-manager" //nolint:gochecknoglobals // versioning

// Version prints version/build.
func Version() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version/build number",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := build.ReadVersion(AppName)
			if err != nil {
				return err
			}

			log.Infoln(version)

			return nil
		},
	}

	return cmd
}

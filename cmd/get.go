package cmd

import (
	"fmt"
	"github.com/omegion/bw-ssh/pkg/exec"
	"log"

	"github.com/omegion/bw-ssh/pkg/bw"
	"github.com/omegion/bw-ssh/pkg/io"

	"github.com/spf13/cobra"
)

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "Name")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

// Get acquires SSH key from Bitwarden.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get SSH key from Bitwarden.",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")

			bitwarden := bw.Bitwarden{
				Commander: exec.Commander{},
			}

			item, err := bitwarden.Get(name)
			if err != nil {
				return err
			}

			if item.IsExists() {
				for _, field := range item.Notes {
					fileName := item.Name

					if field.Name == "public_key" {
						fileName = fmt.Sprintf("%s.pub", item.Name)
					}

					err := io.WriteSSHKey(fileName, []byte(field.Value))
					if err != nil {
						return err
					}
				}

				fmt.Printf("SSH Key %s added.\n", name)

				return nil
			}

			fmt.Printf("Item %s not found\n", name)

			return nil
		},
	}

	setupGetCommand(cmd)

	return cmd
}

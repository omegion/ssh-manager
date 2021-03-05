package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/omegion/bw-ssh/pkg/bw"
	"github.com/omegion/bw-ssh/pkg/exec"

	"github.com/spf13/cobra"
)

// setupAddCommand sets default flags.
func setupAddCommand(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "Name")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("public-key", "", "Public Key file")

	if err := cmd.MarkFlagRequired("public-key"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("private-key", "", "Private Key file")

	if err := cmd.MarkFlagRequired("private-key"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}
}

// Add creates SSH key into Bitwarden.
func Add() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add SSH key to Bitwarden.",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			publicKeyFileName, _ := cmd.Flags().GetString("public-key")
			privateKeyFileName, _ := cmd.Flags().GetString("private-key")

			publicKey, err := readFile(publicKeyFileName)
			if err != nil {
				return err
			}

			privateKey, err := readFile(privateKeyFileName)
			if err != nil {
				return err
			}

			item := bw.Item{
				Type: 1,
				Name: name,
				Notes: []bw.Field{
					{
						Name:  "public_key",
						Value: publicKey,
					},
					{
						Name:  "private_key",
						Value: privateKey,
					},
				},
			}

			bitwarden := bw.Bitwarden{
				Commander: exec.Commander{},
			}
			err = bitwarden.Add(item)
			if err != nil {
				return err
			}

			fmt.Print("Key saved.\n", name)

			return nil
		},
	}

	setupAddCommand(cmd)

	return cmd
}

func readFile(path string) (string, error) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

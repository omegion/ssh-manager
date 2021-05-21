package cmd

import (
	"fmt"
	"github.com/omegion/bw-ssh/internal/provider"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
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

	cmd.Flags().String("provider", "", "Provider")

	if err := cmd.MarkFlagRequired("provider"); err != nil {
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
			providerName, _ := cmd.Flags().GetString("provider")

			publicKey, err := readFile(publicKeyFileName)
			if err != nil {
				return err
			}

			privateKey, err := readFile(privateKeyFileName)
			if err != nil {
				return err
			}

			item := provider.Item{
				Name: name,
				Values: []provider.Field{
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

			commander := provider.NewCommander()

			prv, err := decideProvider(&providerName, &commander)
			if err != nil {
				return err
			}

			err = prv.Add(&item)
			if err != nil {
				return err
			}

			log.Infoln(fmt.Sprintf("SSH Keys saved for %s.", name))

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

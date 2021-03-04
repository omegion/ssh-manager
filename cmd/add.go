package cmd

import (
	"github.com/omegion/lpass-ssh/pkg/lpass"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
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

// Add creates SSH key into LastPass.
func Add() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add SSH key to LastPass.",
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

			sshKey := lpass.SSHKey{
				Name:       name,
				PublicKey:  publicKey,
				PrivateKey: privateKey,
			}

			lastPass := lpass.NewLastPass()

			err = lastPass.Add(sshKey)
			if err != nil {
				return err
			}

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

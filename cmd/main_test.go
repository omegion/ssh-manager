package cmd

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCommander_NewCommander(t *testing.T) {
	commander := NewCommander()

	assert.Equal(t, (*cobra.Command)(nil), commander.Root)
	assert.Equal(t, "", commander.LogLevel)
}

func TestCommander_SetRootCommand(t *testing.T) {
	commander := NewCommander()
	commander.SetRootCommand()

	assert.Equal(t, "vault-unseal", commander.Root.Use)
	assert.Equal(t, "Vault Auto Unseal", commander.Root.Short)
	assert.Equal(t, "CLI command to automatically unseal Vault", commander.Root.Long)
	assert.Equal(t, true, commander.Root.SilenceUsage)
}

func TestCommander_Setup(t *testing.T) {
	commander := NewCommander()
	commander.Setup()

	commander.Root.SetArgs([]string{"version"})

	_, err := commander.Root.ExecuteC()

	expectedLogLevelFlag, _ := commander.Root.Flags().GetString("logLevel")

	assert.Equal(t, nil, err)
	assert.Equal(t, "info", commander.LogLevel)
	assert.Equal(t, "info", log.GetLevel().String())
	assert.Equal(t, "info", expectedLogLevelFlag)
}

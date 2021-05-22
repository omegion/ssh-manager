package ssh

import (
	"bytes"
	"context"

	"github.com/omegion/ssh-manager/internal/provider"

	log "github.com/sirupsen/logrus"
)

// Add adds SSH key to the local agent.
func Add(path string) error {
	commander := provider.NewCommander()

	command := commander.Executor.CommandContext(
		context.Background(),
		"ssh-add",
		path,
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	_, err := command.Output()
	if err != nil {
		return provider.ExecutionFailedError{Command: "ssh-add", Message: stderr.String()}
	}

	log.Debugln("SSH key loaded to the agent.")

	return nil
}

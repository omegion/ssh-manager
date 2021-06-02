package ssh

import (
	"bytes"
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/omegion/ssh-manager/internal"
	"github.com/omegion/ssh-manager/internal/provider"
)

// Add adds SSH key to the local agent.
func Add(path string, commander internal.Commander) error {
	command := commander.Executor.CommandContext(
		context.Background(),
		"ssh-add",
		path,
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	if _, err := command.Output(); err != nil {
		return provider.ExecutionFailedError{Command: "ssh-add", Message: stderr.String()}
	}

	log.Debugln("SSH key loaded to the agent.")

	return nil
}

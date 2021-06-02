package test

import (
	"testing"

	"github.com/omegion/ssh-manager/internal"

	"github.com/stretchr/testify/assert"

	"github.com/omegion/ssh-manager/internal"
)

func TestNewExecutor(t *testing.T) {
	commands := make([]FakeCommand, 0)

	expectedOutput := []byte("test-output")
	command := FakeCommand{
		Command: "cat",
		StdOut:  expectedOutput,
	}

	commands = append(commands, command)

	commander := internal.Commander{Executor: NewExecutor(commands)}
	cmd := commander.Executor.Command("cat")

	output, err := cmd.Output()

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
	assert.Equal(t, []string{command.Command}, command.ToCmd().Argv)
}

func TestNewExecutor_Failure(t *testing.T) {
	commands := make([]FakeCommand, 0)
	command := FakeCommand{
		Command: "cat",
		StdErr:  []byte("custom-error"),
	}

	commands = append(commands, command)

	commander := internal.Commander{Executor: NewExecutor(commands)}
	cmd := commander.Executor.Command("cat")

	_, err := cmd.Output()

	assert.EqualError(t, err, "custom-error")
	assert.Equal(t, []string{command.Command}, command.ToCmd().Argv)
}

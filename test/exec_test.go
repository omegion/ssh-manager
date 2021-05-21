package test_test

import (
	"context"
	"testing"

	"github.com/acquia/kaas-cluster/test"
	"github.com/stretchr/testify/assert"
)

func TestExecutor_singleCall(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               "my fancy command",
			StdOut:                []byte("standard output"),
			StdErr:                []byte("standard error"),
			ExpectedNumberOfCalls: 1,
		},
	})

	command := e.CommandContext(context.Background(), "my", "fancy", "command")
	output, err := command.Output()

	assert.NoError(t, err)
	assert.NoError(t, e.Validate())

	assert.Equal(t, "standard output", string(output))
}

func TestExecutor_multipleCalls_failed_missingCall(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               "my fancy command",
			StdOut:                []byte("standard output"),
			StdErr:                []byte("standard error"),
			ExpectedNumberOfCalls: 2,
		},
	})

	command := e.CommandContext(context.Background(), "my", "fancy", "command")
	output, err := command.Output()

	assert.NoError(t, err)
	assert.Equal(t, "standard output", string(output))

	assert.Error(t, e.Validate())
}

func TestExecutor_multipleCommands(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               "my fancy command",
			StdOut:                []byte("standard output"),
			StdErr:                []byte("standard error"),
			ExpectedNumberOfCalls: 1,
		},
		{
			Command:               "another fancy call",
			StdOut:                []byte("another standard output"),
			StdErr:                []byte("another standard error"),
			ExpectedNumberOfCalls: 1,
		},
	})

	command := e.CommandContext(context.Background(), "my", "fancy", "command")
	output, err := command.Output()

	assert.NoError(t, err)
	assert.Equal(t, "standard output", string(output))

	command = e.CommandContext(context.Background(), "another", "fancy", "call")
	output, err = command.CombinedOutput()

	assert.NoError(t, err)
	assert.Equal(t, "another standard output\nanother standard error", string(output))

	assert.NoError(t, e.Validate())
}

func TestExecutor_multipleCommands_missingCall_failed(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               "my fancy command",
			StdOut:                []byte("standard output"),
			StdErr:                []byte("standard error"),
			ExpectedNumberOfCalls: 1,
		},
		{
			Command:               "another fancy call",
			StdOut:                []byte("another standard output"),
			StdErr:                []byte("another standard error"),
			ExpectedNumberOfCalls: 1,
		},
	})

	command := e.CommandContext(context.Background(), "my", "fancy", "command")
	output, err := command.Output()

	assert.NoError(t, err)
	assert.Equal(t, "standard output", string(output))

	assert.Error(t, e.Validate())
}

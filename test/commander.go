package test

import (
	"bytes"
	"errors"
	"strings"

	"k8s.io/utils/exec"
	testingexec "k8s.io/utils/exec/testing"
)

// FakeCommand is a command wrapper for testing.
type FakeCommand struct {
	Command string
	StdOut  []byte
	StdErr  []byte
}

// ToCmd converts FakeCommand to FakeCmd.
func (c FakeCommand) ToCmd() testingexec.FakeCmd {
	return testingexec.FakeCmd{
		Argv: strings.Split(c.Command, " "),
		OutputScript: []testingexec.FakeAction{
			func() ([]byte, []byte, error) {
				if bytes.Equal(c.StdErr, []byte("")) {
					return c.StdOut, nil, nil
				}
				//nolint:goerr113 // allow static errors.
				return c.StdOut, nil, errors.New(string(c.StdErr))
			},
		},
	}
}

// NewExecutor is a factory for Commander testing.
func NewExecutor(commands []FakeCommand) *testingexec.FakeExec {
	cmdActions := make([]testingexec.FakeCommandAction, 0)

	for i := range commands {
		fakeCmd := commands[i].ToCmd()

		cmdActions = append(cmdActions, func(c string, args ...string) exec.Cmd {
			return &fakeCmd
		})
	}

	return &testingexec.FakeExec{
		ExactOrder:    true,
		CommandScript: cmdActions,
	}
}

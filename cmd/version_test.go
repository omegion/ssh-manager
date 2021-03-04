package cmd

import (
	"testing"
)

func TestVersionCommand(t *testing.T) {
	_, err := executeCommand(Sync())
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}

func TestAddCommand(t *testing.T) {
	_, err := executeCommand(Add(),
		"--name=ssh-key-test",
		"--private-key=/Users/hakan/.ssh/hetzner",
		"--public-key=/Users/hakan/.ssh/hetzner.pub",
	)
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}

func TestGetCommand(t *testing.T) {
	_, err := executeCommand(Get(),
		"--name=test",
	)
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}

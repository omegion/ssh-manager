package cmd

import (
	"testing"
)

func TestGetCommand(t *testing.T) {
	_, err := executeCommand(Get(),
		"--name=ssh-test-1",
	)
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}

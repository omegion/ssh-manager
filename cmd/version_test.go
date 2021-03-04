package cmd

import (
	"testing"
)

func TestGetCommand(t *testing.T) {
	_, err := executeCommand(Version())
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}

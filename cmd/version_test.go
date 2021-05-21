package cmd

import (
	"testing"
)

func Test_GetCommand(t *testing.T) {
	_, err := executeCommand(Get(), "--name=test", "--provider=bw")
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}

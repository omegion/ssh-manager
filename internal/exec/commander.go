package exec

import "os/exec"

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/commander_mock.go -package=mocks github.com/omegion/bw-ssh/pkg/exec CommanderInterface
// CommanderInterface is an interface for Commander.
type CommanderInterface interface {
	Output(string, ...string) ([]byte, error)
}

// Commander is custom wrapper for exec.Command.
type Commander struct{}

// Output is for executing for exec.Command.
func (c Commander) Output(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}

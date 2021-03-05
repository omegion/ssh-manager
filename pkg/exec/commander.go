package exec

import "os/exec"

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/commander_mock.go -package=mocks github.com/omegion/bw-ssh/pkg/exec CommanderInterface
type CommanderInterface interface {
	Output(string, ...string) ([]byte, error)
}

type Commander struct{}

func (c Commander) Output(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}

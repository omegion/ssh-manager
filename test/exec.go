package test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"k8s.io/utils/exec"
)

type CommandWithOutput struct {
	Command               string
	StdOut                []byte
	StdErr                []byte
	ExpectedNumberOfCalls int
	NumberCalls           int
}

type Executor struct {
	stack []CommandWithOutput
}

func NewExecutor(commands []CommandWithOutput) *Executor {
	return &Executor{stack: commands}
}

func (e *Executor) Validate() error {
	for _, cmd := range e.stack {
		if cmd.ExpectedNumberOfCalls != cmd.NumberCalls {
			return MissingCallError{Command: cmd.Command, Expected: cmd.ExpectedNumberOfCalls, Actual: cmd.NumberCalls}
		}
	}

	return nil
}

func (e *Executor) Reset() {
	for idx := range e.stack {
		e.stack[idx].NumberCalls = 0
	}
}

func (e *Executor) Command(cmd string, args ...string) exec.Cmd {
	return e.CommandContext(context.Background(), cmd, args...)
}

func (e *Executor) CommandContext(ctx context.Context, cmd string, args ...string) exec.Cmd {
	command := strings.Join(append([]string{cmd}, args...), " ")

	for idx, c := range e.stack {
		if c.ExpectedNumberOfCalls == c.NumberCalls {
			continue
		}

		if c.Command == command {
			e.stack[idx].NumberCalls++

			return &mockCmd{
				Command: c.Command,
				StdOut:  c.StdOut,
				StdErr:  c.StdErr,
				Context: ctx,
				Error:   nil,
			}
		}

		break
	}

	return &mockCmd{
		Command: command,
		StdOut:  []byte{},
		StdErr:  []byte{},
		Context: ctx,
		Error:   UnexpectedCommandCallError{Command: command},
	}
}

func (e *Executor) LookPath(file string) (string, error) {
	return file, nil
}

type mockCmd struct {
	Command string
	StdOut  []byte
	StdErr  []byte
	Context context.Context
	Error   error
}

func (m *mockCmd) Run() error {
	return m.Error
}

func (m *mockCmd) CombinedOutput() ([]byte, error) {
	out := []byte{}

	out = append(out, m.StdOut...)
	out = append(out, '\n')
	out = append(out, m.StdErr...)

	return out, m.Error
}

func (m *mockCmd) Output() ([]byte, error) {
	return m.StdOut, m.Error
}

func (m *mockCmd) StdoutPipe() (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewReader(m.StdOut)), m.Error
}

func (m *mockCmd) StderrPipe() (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewReader(m.StdErr)), m.Error
}

func (m *mockCmd) Start() error {
	return m.Error
}

func (m *mockCmd) Wait() error {
	return m.Error
}

func (m *mockCmd) SetDir(dir string)       {}
func (m *mockCmd) SetStdin(in io.Reader)   {}
func (m *mockCmd) SetStdout(out io.Writer) {}
func (m *mockCmd) SetStderr(out io.Writer) {}
func (m *mockCmd) SetEnv(env []string)     {}
func (m *mockCmd) Stop()                   {}

type UnexpectedCommandCallError struct {
	Command string
}

func (e UnexpectedCommandCallError) Error() string {
	return fmt.Sprintf("unexpected command: %s", e.Command)
}

type MissingCallError struct {
	Command  string
	Expected int
	Actual   int
}

func (e MissingCallError) Error() string {
	return fmt.Sprintf("missing command call: %s; expected = %d; got = %d", e.Command, e.Expected, e.Actual)
}

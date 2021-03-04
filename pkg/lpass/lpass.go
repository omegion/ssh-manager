package lpass

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
)

const (
	lastPassCommand = "lpass"
)

type LastPass struct {
	Host    string
	Options []string
}

func NewLastPass() LastPass {
	return LastPass{}
}

func (l *LastPass) Sync() error {
	return nil
}

func (l *LastPass) Add(key SSHKey) error {
	printCmd := exec.Command("printf", key.Serialize())
	cmd := exec.Command(lastPassCommand, l.addOptions(key)...)
	cmd.Stdin, _ = printCmd.StdoutPipe()
	cmd.Stdout = os.Stdout
	_ = cmd.Start()
	_ = printCmd.Run()

	err := cmd.Wait()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError
		}
	}

	return nil
}

func (l *LastPass) Get(key SSHKey) error {
	var outBuffer bytes.Buffer

	cmd := exec.Command(lastPassCommand, l.getOptions(key)...)
	cmd.Stdout = &outBuffer
	//cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError
		}
	}

	var note []Note

	err = json.Unmarshal(outBuffer.Bytes(), &note)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(note[0].Note), &key)
	if err != nil {
		panic(err)
	}

	//fmt.Printf(outBuffer.String())
	return nil
}

func (l LastPass) addOptions(key SSHKey) []string {
	options := []string{
		"add",
		"--non-interactive",
		"--sync=now",
		"--note-type=ssh-key",
	}

	options = append(options, l.Options...)
	options = append(options, fmt.Sprintf(`%s`, key.GetPath()))

	return options
}

func (l *LastPass) getOptions(key SSHKey) []string {
	options := []string{
		"show",
		"--json",
	}

	options = append(options, l.Options...)
	options = append(options, fmt.Sprintf(`%s`, key.GetPath()))

	return options
}

func (l *LastPass) Run(args []string) error {
	cmd := exec.Command(lastPassCommand, args...)
	//cmd.Stdout = os.Stdout
	//cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

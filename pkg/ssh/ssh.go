package ssh

import "os/exec"

// Add adds SSH key to the local agent.
func Add(path string) error {
	options := []string{
		path,
	}

	_, err := exec.Command("ssh-add", options...).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError
		}

		return err
	}

	return nil
}

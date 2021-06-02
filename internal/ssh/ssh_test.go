package ssh

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/omegion/ssh-manager/internal"
	"github.com/omegion/ssh-manager/test"
)

func TestAdd(t *testing.T) {
	path := "/var/test"

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf("ssh-add %s", path),
		},
	}

	commander := internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := Add(path, commander)

	assert.NoError(t, err)
}

func TestAdd_Failure(t *testing.T) {
	path := "/var/test"

	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf("ssh-add %s", path),
			StdErr:  []byte("custom-error"),
		},
	}

	commander := internal.Commander{Executor: test.NewExecutor(expectedCommands)}

	err := Add(path, commander)

	assert.EqualError(t, err, "'ssh-add': Execution failed: ")
}

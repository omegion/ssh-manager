package provider_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/omegion/ssh-manager/internal"
	"github.com/omegion/ssh-manager/internal/provider"
	"github.com/omegion/ssh-manager/test"
)

const (
	encodedValues = "W3sibmFtZSI6InByaXZhdGVfa2V5IiwidmFsdWUiOiJYIn0seyJuYW1lIjoicHVibGljX2tleSIsInZhbHVlIjoiWSJ9XQ=="
)

func TestOnePassword_Add(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf("op get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdErr:  test.Must(test.LoadFixture("op_get.txt")),
		},
		{
			Command: fmt.Sprintf(
				"op create item login notesPlain=%s --title %s%s --tags %s",
				encodedValues,
				provider.BitwardenDefaultPrefix,
				"test",
				strings.Replace(provider.OnePasswordDefaultPrefix, "__", "", 1),
			),
			StdOut: test.Must(test.LoadFixture("op_add.txt")),
		},
	}

	onep := provider.OnePassword{
		Commander: internal.Commander{Executor: test.NewExecutor(expectedCommands)},
	}

	item := provider.Item{
		Name: "test",
		Values: []provider.Field{
			{
				Name:  "private_key",
				Value: "X",
			},
			{
				Name:  "public_key",
				Value: "Y",
			},
		},
	}

	err := onep.Add(&item)

	assert.NoError(t, err)
}

func TestOnePassword_Get(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf("op get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdOut:  test.Must(test.LoadFixture("op_get.txt")),
		},
	}

	op := provider.OnePassword{
		Commander: internal.Commander{Executor: test.NewExecutor(expectedCommands)},
	}

	item, err := op.Get("test")

	assert.NoError(t, err)
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "X", item.ID)
}

func TestOnePassword_GetNotFound(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf("op get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdErr:  test.Must(test.LoadFixture("op_get_not_found.txt")),
		},
	}

	op := provider.OnePassword{
		Commander: internal.Commander{Executor: test.NewExecutor(expectedCommands)},
	}

	_, err := op.Get("test")

	assert.Error(t, err)
}

func TestOnePassword_List(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: fmt.Sprintf(
				"op list items --categories login --tags %s",
				strings.Replace(provider.OnePasswordDefaultPrefix, "__", "", 1),
			),
			StdOut: test.Must(test.LoadFixture("op_list.txt")),
		},
	}

	op := provider.OnePassword{
		Commander: internal.Commander{Executor: test.NewExecutor(expectedCommands)},
	}

	items, err := op.List()

	expectedItems := map[string]string{
		"X": "test1",
		"Y": "test2",
	}

	assert.NoError(t, err)

	for _, item := range items {
		assert.Equal(t, expectedItems[item.ID], item.Name)
	}
}

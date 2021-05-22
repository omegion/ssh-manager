package provider_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/omegion/ssh-manager/internal/provider"
	"github.com/omegion/ssh-manager/test"

	"github.com/stretchr/testify/assert"
)

const (
	encodedValues = "W3sibmFtZSI6InByaXZhdGVfa2V5IiwidmFsdWUiOiJYIn0seyJuYW1lIjoicHVibGljX2tleSIsInZhbHVlIjoiWSJ9XQ=="
)

func TestOnePassword_Add(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command: fmt.Sprintf(
				"op create item login notesPlain=%s --title %s%s --tags %s",
				encodedValues,
				provider.BitwardenDefaultPrefix,
				"test",
				strings.Replace(provider.OnePasswordDefaultPrefix, "__", "", 1),
			),
			StdOut:                test.Must(test.LoadFixture("op_add.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	op := provider.OnePassword{
		Commander: commander,
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

	err := op.Add(&item)

	assert.NoError(t, err)
	assert.NoError(t, e.Validate())
}

func TestOnePassword_Get(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               fmt.Sprintf("op get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdOut:                test.Must(test.LoadFixture("op_get.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	op := provider.OnePassword{
		Commander: commander,
	}

	item, err := op.Get("test")

	assert.NoError(t, err)
	assert.NoError(t, e.Validate())
	assert.Equal(t, "test", item.Name)
	assert.Equal(t, "X", item.ID)
}

func TestOnePassword_GetNotFound(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               fmt.Sprintf("op get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdOut:                []byte{},
			StdErr:                test.Must(test.LoadFixture("op_get_not_found.txt")),
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	op := provider.OnePassword{
		Commander: commander,
	}

	_, err := op.Get("test")

	assert.Error(t, err)
	assert.NoError(t, e.Validate())
}

func TestOnePassword_List(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command: fmt.Sprintf(
				"op list items --categories login --tags %s",
				strings.Replace(provider.OnePasswordDefaultPrefix, "__", "", 1),
			),
			StdOut:                test.Must(test.LoadFixture("op_list.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	op := provider.OnePassword{
		Commander: commander,
	}

	items, err := op.List()

	expectedItems := map[string]string{
		"X": "test1",
		"Y": "test2",
	}

	assert.NoError(t, err)
	assert.NoError(t, e.Validate())

	for _, item := range items {
		assert.Equal(t, expectedItems[item.ID], item.Name)
	}
}

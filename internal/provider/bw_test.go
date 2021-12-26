package provider_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/omegion/ssh-manager/internal"
	"github.com/omegion/ssh-manager/internal/provider"
	"github.com/omegion/ssh-manager/test"
)

func TestBitwarden_Add(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: "bw sync",
		},
		{
			Command: fmt.Sprintf("bw get item %s%s", provider.BitwardenDefaultPrefix, "test"),
		},
		{
			//nolint:lll // allow long lines.
			Command: "bw create item eyJpZCI6bnVsbCwidHlwZSI6MSwibmFtZSI6IlNTSEtleXNfX3Rlc3QiLCJub3RlcyI6Ilczc2libUZ0WlNJNkluQnlhWFpoZEdWZmEyVjVJaXdpZG1Gc2RXVWlPaUpZSW4wc2V5SnVZVzFsSWpvaWNIVmliR2xqWDJ0bGVTSXNJblpoYkhWbElqb2lXU0o5WFE9PSIsImxvZ2luIjoidGVzdCJ9",
			StdOut:  test.Must(test.LoadFixture("bw_add.txt")),
		},
	}

	bitw := provider.Bitwarden{
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

	err := bitw.Add(&item)

	assert.NoError(t, err)
}

func TestBitwarden_Get(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: "bw sync",
		},
		{
			Command: fmt.Sprintf("bw get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdOut:  test.Must(test.LoadFixture("bw_get.txt")),
		},
	}

	bw := provider.Bitwarden{
		Commander: internal.Commander{Executor: test.NewExecutor(expectedCommands)},
	}

	item, err := bw.Get("test")

	assert.NoError(t, err)
	assert.Equal(t, "test", item.Name)
}

func TestBitwarden_GetNotFound(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: "bw sync",
		},
		{
			Command: fmt.Sprintf("bw get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdErr:  test.Must(test.LoadFixture("bw_get_not_found.txt")),
		},
	}

	bw := provider.Bitwarden{
		Commander: internal.Commander{Executor: test.NewExecutor(expectedCommands)},
	}

	_, err := bw.Get("test")

	assert.EqualError(t, err, "'bw get': Execution failed: ")
}

func TestBitwarden_Sync(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: "bw sync",
		},
	}

	bw := provider.Bitwarden{
		Commander: internal.Commander{Executor: test.NewExecutor(expectedCommands)},
	}

	err := bw.Sync()

	assert.NoError(t, err)
}

func TestBitwarden_List(t *testing.T) {
	expectedCommands := []test.FakeCommand{
		{
			Command: "bw sync",
		},
		{
			Command: fmt.Sprintf("bw list items --search %s", provider.BitwardenDefaultPrefix),
			StdOut:  test.Must(test.LoadFixture("bw_list.txt")),
		},
	}

	bw := provider.Bitwarden{
		Commander: internal.Commander{Executor: test.NewExecutor(expectedCommands)},
	}

	items, err := bw.List()

	expectedItems := []string{
		"test1",
		"test2",
	}

	assert.NoError(t, err)

	for idx, item := range items {
		assert.Equal(t, expectedItems[idx], item.Name)
	}
}

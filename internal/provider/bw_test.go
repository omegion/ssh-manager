package provider_test

import (
	"fmt"
	"testing"

	"github.com/omegion/ssh-manager/internal/provider"
	"github.com/omegion/ssh-manager/test"

	"github.com/stretchr/testify/assert"
)

func TestBitwarden_Add(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			//nolint:lll // get this one from fixture.
			Command:               "bw create item eyJpZCI6bnVsbCwidHlwZSI6MSwibmFtZSI6IlNTSEtleXNfX3Rlc3QiLCJub3RlcyI6Ilczc2libUZ0WlNJNkluQnlhWFpoZEdWZmEyVjVJaXdpZG1Gc2RXVWlPaUpZSW4wc2V5SnVZVzFsSWpvaWNIVmliR2xqWDJ0bGVTSXNJblpoYkhWbElqb2lXU0o5WFE9PSIsImxvZ2luIjoidGVzdCJ9",
			StdOut:                test.Must(test.LoadFixture("bw_add.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	bw := provider.Bitwarden{
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

	err := bw.Add(&item)

	assert.NoError(t, err)
	assert.NoError(t, e.Validate())
}

func TestBitwarden_Get(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               "bw sync",
			StdOut:                test.Must(test.LoadFixture("bw_sync.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
		{
			Command:               fmt.Sprintf("bw get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdOut:                test.Must(test.LoadFixture("bw_get.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	bw := provider.Bitwarden{
		Commander: commander,
	}

	item, err := bw.Get("test")

	assert.NoError(t, err)
	assert.NoError(t, e.Validate())
	assert.Equal(t, "test", item.Name)
}

func TestBitwarden_GetNotFound(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               "bw sync",
			StdOut:                test.Must(test.LoadFixture("bw_sync.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
		{
			Command:               fmt.Sprintf("bw get item %s%s", provider.BitwardenDefaultPrefix, "test"),
			StdOut:                []byte{},
			StdErr:                test.Must(test.LoadFixture("bw_get_not_found.txt")),
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	bw := provider.Bitwarden{
		Commander: commander,
	}

	_, err := bw.Get("test")

	assert.Error(t, err)
	assert.NoError(t, e.Validate())
}

func TestBitwarden_Sync(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               "bw sync",
			StdOut:                test.Must(test.LoadFixture("bw_sync.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	bw := provider.Bitwarden{
		Commander: commander,
	}

	err := bw.Sync()

	assert.NoError(t, err)
	assert.NoError(t, e.Validate())
}

func TestBitwarden_List(t *testing.T) {
	e := test.NewExecutor([]test.CommandWithOutput{
		{
			Command:               "bw sync",
			StdOut:                test.Must(test.LoadFixture("bw_sync.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
		{
			Command:               fmt.Sprintf("bw list items --search %s", provider.BitwardenDefaultPrefix),
			StdOut:                test.Must(test.LoadFixture("bw_list.txt")),
			StdErr:                []byte{},
			ExpectedNumberOfCalls: 1,
		},
	})

	commander := provider.NewCommander()
	commander.Executor = e

	bw := provider.Bitwarden{
		Commander: commander,
	}

	items, err := bw.List()

	expectedItems := []string{
		"test1",
		"test2",
	}

	assert.NoError(t, err)
	assert.NoError(t, e.Validate())

	for idx, item := range items {
		assert.Equal(t, expectedItems[idx], item.Name)
	}

}

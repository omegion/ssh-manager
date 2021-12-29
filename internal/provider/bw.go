package provider

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/omegion/ssh-manager/internal"
)

const (
	// BitwardenDefaultPrefix default prefix for BitwardenItem.
	BitwardenDefaultPrefix = "SSHKeys__"
	// BitwardenCommand base command for Bitwarden.
	BitwardenCommand = "bw"
)

// Bitwarden for connection.
type Bitwarden struct {
	Commander internal.Commander
}

// BitwardenItem is item adapter for provider Item.
type BitwardenItem struct {
	ID    *string `json:"id"`
	Type  int     `json:"type"`
	Name  string  `json:"name"`
	Notes string  `json:"notes"`
	Login string  `json:"login"`
}

// GetName returns name of the provider.
func (b Bitwarden) GetName() string {
	return BitwardenCommand
}

// Add adds given item to Bitwarden.
func (b Bitwarden) Add(item *Item) error {
	_, err := b.Get(GetOptions{
		Name: item.Name,
	})
	if err == nil {
		return ItemAlreadyExistsError{Name: item.Name}
	}

	encodedValues, err := item.EncodeValues()
	if err != nil {
		return err
	}

	bwItem := BitwardenItem{
		ID:    nil,
		Type:  1,
		Name:  fmt.Sprintf("%s%s", BitwardenDefaultPrefix, item.Name),
		Notes: encodedValues,
		Login: item.Name,
	}

	var bwItemByte []byte

	bwItemByte, err = json.Marshal(bwItem)
	if err != nil {
		return err
	}

	command := b.Commander.Executor.CommandContext(
		context.Background(),
		BitwardenCommand,
		"create",
		"item",
		base64.StdEncoding.EncodeToString(bwItemByte),
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	_, err = command.Output()
	if err != nil {
		return ExecutionFailedError{Command: "bw get", Message: stderr.String()}
	}

	return nil
}

// Get gets Item from Bitwarden with given item name.
func (b Bitwarden) Get(options GetOptions) (*Item, error) {
	err := b.Sync()
	if err != nil {
		return &Item{}, err
	}

	command := b.Commander.Executor.CommandContext(
		context.Background(),
		BitwardenCommand,
		"get",
		"item",
		fmt.Sprintf("%s%s", BitwardenDefaultPrefix, options.Name),
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	log.Debugln(fmt.Sprintf("Getting item %s in Bitwarden.", options.Name))

	output, err := command.Output()
	if err != nil {
		return &Item{}, ExecutionFailedError{Command: "bw get", Message: stderr.String()}
	}

	var tmpItem struct {
		ID    *string `json:"id"`
		Name  string  `json:"name"`
		Notes string  `json:"notes"`
	}

	err = json.Unmarshal(output, &tmpItem)
	if err != nil {
		return &Item{}, err
	}

	item := Item{
		ID:   *tmpItem.ID,
		Name: strings.Replace(tmpItem.Name, BitwardenDefaultPrefix, "", 1),
	}

	log.Debugln(fmt.Sprintf("Decoding keys for item %s.", options.Name))

	decodedRawNotes, err := base64.StdEncoding.DecodeString(tmpItem.Notes)
	if err != nil {
		return &Item{}, err
	}

	err = json.Unmarshal(decodedRawNotes, &item.Values)
	if err != nil {
		return &Item{}, err
	}

	return &item, nil
}

// List lists all added SSH keys.
func (b Bitwarden) List(options ListOptions) ([]*Item, error) {
	err := b.Sync()
	if err != nil {
		return []*Item{}, err
	}

	command := b.Commander.Executor.CommandContext(
		context.Background(),
		BitwardenCommand,
		"list",
		"items",
		"--search",
		BitwardenDefaultPrefix,
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	log.Debugln("Getting items in Bitwarden.")

	output, err := command.Output()
	if err != nil {
		return []*Item{}, ExecutionFailedError{Command: "bw get", Message: stderr.String()}
	}

	type tmpItem struct {
		ID    *string `json:"id"`
		Name  string  `json:"name"`
		Notes string  `json:"notes"`
	}

	var tmpItems []tmpItem

	err = json.Unmarshal(output, &tmpItems)
	if err != nil {
		return []*Item{}, err
	}

	items := make([]*Item, 0)

	for _, item := range tmpItems {
		items = append(items, &Item{
			ID:   *item.ID,
			Name: strings.Replace(item.Name, BitwardenDefaultPrefix, "", 1),
		})
	}

	return items, nil
}

// Sync syncs Bitwarden vault.
func (b Bitwarden) Sync() error {
	command := b.Commander.Executor.CommandContext(
		context.Background(),
		BitwardenCommand,
		"sync",
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	if _, err := command.Output(); err != nil {
		return ExecutionFailedError{Command: "bw sync", Message: stderr.String()}
	}

	log.Debugln("Syncing Bitwarden Vault.")

	return nil
}

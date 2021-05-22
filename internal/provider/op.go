package provider

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	// OnePasswordDefaultPrefix default prefix for OnePasswordItem.
	OnePasswordDefaultPrefix = "SSHKeys__"
	// OnePasswordCommand base command for OnePassword.
	OnePasswordCommand = "op"
)

// OnePassword for connection.
type OnePassword struct {
	Commander Commander
}

// OnePasswordItemOverview is item adapter for provider OnePasswordItem.
type OnePasswordItemOverview struct {
	Title *string `json:"title"`
}

// OnePasswordItemDetails is item adapter for provider Item.
type OnePasswordItemDetails struct {
	NotesPlain *string `json:"notesPlain"`
}

// OnePasswordItem is item adapter for provider OnePasswordItem.
type OnePasswordItem struct {
	UUID     *string                 `json:"uuid"`
	Details  OnePasswordItemDetails  `json:"details"`
	Overview OnePasswordItemOverview `json:"overview"`
}

// Add adds given item to OnePassword.
func (o OnePassword) Add(item *Item) error {
	_, err := o.Get(item.Name)
	if err == nil {
		return ItemAlreadyExists{Name: item.Name}
	}

	encodedValues, err := item.EncodeValues()
	if err != nil {
		return err
	}

	command := o.Commander.Executor.CommandContext(
		context.Background(),
		OnePasswordCommand,
		"create",
		"item",
		"login",
		fmt.Sprintf("notesPlain=%s", encodedValues),
		"--title",
		fmt.Sprintf("%s%s", OnePasswordDefaultPrefix, item.Name),
		"--tags",
		strings.Replace(OnePasswordDefaultPrefix, "__", "", 1),
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	_, err = command.Output()
	if err != nil {
		return ExecutionFailedError{Command: "op create", Message: stderr.String()}
	}

	return nil
}

// Get gets Item from OnePassword with given item name.
func (o OnePassword) Get(name string) (Item, error) {
	command := o.Commander.Executor.CommandContext(
		context.Background(),
		OnePasswordCommand,
		"get",
		"item",
		fmt.Sprintf("%s%s", OnePasswordDefaultPrefix, name),
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	log.Debugln(fmt.Sprintf("Getting item %s in OnePassword.", name))

	output, err := command.Output()
	if err != nil {
		return Item{}, ExecutionFailedError{Command: "op get item", Message: stderr.String()}
	}

	var opItem OnePasswordItem

	err = json.Unmarshal(output, &opItem)
	if err != nil {
		return Item{}, err
	}

	item := Item{
		ID:   *opItem.UUID,
		Name: strings.Replace(*opItem.Overview.Title, OnePasswordDefaultPrefix, "", 1),
	}

	log.Debugln(fmt.Sprintf("Decoding keys for item %s.", name))

	decodedRawNotes, err := base64.StdEncoding.DecodeString(*opItem.Details.NotesPlain)
	if err != nil {
		return Item{}, err
	}

	err = json.Unmarshal(decodedRawNotes, &item.Values)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

// List lists all added SSH keys.
func (o OnePassword) List() ([]Item, error) {
	command := o.Commander.Executor.CommandContext(
		context.Background(),
		OnePasswordCommand,
		"list",
		"items",
		"--categories",
		"login",
		"--tags",
		strings.Replace(OnePasswordDefaultPrefix, "__", "", 1),
	)

	var stderr bytes.Buffer

	command.SetStderr(&stderr)

	log.Debugln("Getting items in OnePassword.")

	output, err := command.Output()
	if err != nil {
		return []Item{}, ExecutionFailedError{Command: "op list", Message: stderr.String()}
	}

	var opItems []OnePasswordItem

	err = json.Unmarshal(output, &opItems)
	if err != nil {
		return []Item{}, err
	}

	var items []Item

	for _, item := range opItems {
		items = append(items, Item{
			ID:   *item.UUID,
			Name: strings.Replace(*item.Overview.Title, OnePasswordDefaultPrefix, "", 1),
		})
	}

	return items, nil
}

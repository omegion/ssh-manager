package bw

import (
	"encoding/json"
	"fmt"
	"os/exec"

	cmd "github.com/omegion/bw-ssh/pkg/exec"
)

const (
	DefaultFolderName = "SSHKeys"
	BitwardenCommand  = "bw"
)

// Bitwarden for connection.
type Bitwarden struct {
	Items     []Item
	FolderID  string
	Options   []string
	Commander cmd.CommanderInterface
}

// Sync updates local cache.
func (l *Bitwarden) Sync() error {
	options := l.syncOptions()

	_, err := l.Commander.Output(BitwardenCommand, options...)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError
		}
	}

	return nil
}

func (l *Bitwarden) GetFolder() (Folder, error) {
	folder := Folder{
		Name: DefaultFolderName,
	}

	options := l.getFolderOptions(folder)

	out, err := l.Commander.Output(BitwardenCommand, options...)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return folder, exitError
		}
	}

	var folders []Folder

	err = json.Unmarshal(out, &folders)
	if err != nil {
		return folder, err
	}

	if len(folders) == 0 {
		// Create a folder.
	} else {
		folder = folders[0]
		l.FolderID = folder.ID
	}

	return folder, nil
}

func (l *Bitwarden) GetItems() error {
	//nolint:nestif // refactor this function.
	if len(l.Items) == 0 || l.FolderID == "" {
		err := l.Sync()
		if err != nil {
			return err
		}

		folder, err := l.GetFolder()
		if err != nil {
			return err
		}

		options := l.listOptions(folder)

		out, err := l.Commander.Output(BitwardenCommand, options...)
		if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				return exitError
			}
		}

		err = json.Unmarshal(out, &l.Items)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get gets item from Bitwarden.
func (l *Bitwarden) Get(name string) (Item, error) {
	err := l.GetItems()
	if err != nil {
		return Item{}, err
	}

	item := Item{}

	for i := range l.Items {
		if l.Items[i].Name == name {
			return l.Items[i], nil
		}
	}

	return item, nil
}

// Add adds new item to Bitwarden.
func (l *Bitwarden) Add(item Item) error {
	receivedItem, err := l.Get(item.Name)
	if err != nil {
		return err
	}

	if receivedItem.IsExists() {
		return BitwardenError{
			//nolint:goerr113 // replace this with custom error.
			Origin: fmt.Errorf("%s is already exists", item.Name),
		}
	}

	item.FolderID = l.FolderID

	encodedItem, err := item.Encode()
	if err != nil {
		return err
	}

	options := l.addOptions(encodedItem)

	_, err = l.Commander.Output(BitwardenCommand, options...)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError
		}

		return err
	}

	return nil
}

func (l *Bitwarden) syncOptions() []string {
	options := []string{
		"sync",
	}

	options = append(options, l.Options...)

	return options
}

func (l *Bitwarden) addOptions(encodedItem string) []string {
	options := []string{
		"create",
		"item",
	}

	options = append(options, l.Options...)
	options = append(options, encodedItem)

	return options
}

func (l *Bitwarden) getFolderOptions(folder Folder) []string {
	options := []string{
		"list",
		"folders",
	}

	options = append(options, l.Options...)
	options = append(options, "--search")
	options = append(options, folder.Name)

	return options
}

func (l *Bitwarden) listOptions(folder Folder) []string {
	options := []string{
		"list",
		"items",
	}

	options = append(options, l.Options...)
	options = append(options, "--folderid")
	options = append(options, folder.ID)

	return options
}

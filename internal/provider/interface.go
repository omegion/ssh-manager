package provider

import (
	"encoding/base64"
	"encoding/json"
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/interface_mock.go -package=mocks github.com/omegion/ssh-manager/internal/provider Interface
// Interface is an interface for all providers.
type Interface interface {
	GetName() string
	Add(item *Item) error
	Get(name string) (*Item, error)
	List() ([]*Item, error)
}

// Field is custom fields under Item.
type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Item is a secure item of provider.
type Item struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name"`
	Values []Field `json:"value"`
}

// EncodeValues encodes Values.
func (i Item) EncodeValues() (string, error) {
	var p []byte

	p, err := json.Marshal(i.Values)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(p), nil
}

package provider

import (
	"encoding/base64"
	"encoding/json"
)

// Interface is an interface for all providers.
//
//go:generate mockgen -destination=mocks/interface_mock.go -package=mocks github.com/omegion/ssh-manager/internal/provider Interface
type Interface interface {
	GetName() string
	Add(item *Item) error
	Get(options GetOptions) (*Item, error)
	List(options ListOptions) ([]*Item, error)
}

// Field is custom fields under Item.
type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetOptions is options for Get method.
type GetOptions struct {
	Name   string
	Bucket *string
}

// ListOptions is options for List method.
type ListOptions struct {
	Bucket *string
}

// Item is a secure item of provider.
type Item struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name"`
	Values []Field `json:"value"`
	Bucket *string
}

// EncodeValues encodes Values.
func (i Item) EncodeValues() (string, error) {
	var values []byte

	values, err := i.MarshalValues()
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(values), nil
}

// MarshalValues encodes the Values to JSON bytes.
func (i Item) MarshalValues() ([]byte, error) {
	var values []byte

	values, err := json.Marshal(i.Values)
	if err != nil {
		return nil, err
	}

	return values, nil
}

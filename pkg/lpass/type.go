package lpass

import (
	"fmt"
)

type Note struct {
	Note string `json:"note"`
}

type SSHKey struct {
	Name       string
	PublicKey  string `yaml:"Public Key"`
	PrivateKey string `yaml:"Private Key"`
	NoteType   string `yaml:"NoteType"`
}

func (k SSHKey) GetPath() string {
	return fmt.Sprintf("SSHKeys/%s", k.Name)
}

func (k SSHKey) Serialize() string {
	return fmt.Sprintf("Public Key:%s\nPrivate Key:%s", k.PublicKey, k.PrivateKey)
}

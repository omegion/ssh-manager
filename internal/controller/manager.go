package controller

import (
	"github.com/omegion/ssh-manager/internal"
	"github.com/omegion/ssh-manager/internal/provider"
)

// Manager is a controller for SSH providers.
type Manager struct {
	Provider provider.Interface
}

// NewManager is a factory for Manager.
func NewManager(providerName *string) *Manager {
	return &Manager{Provider: getProviderByName(providerName)}
}

// Add adds item to provider.
func (c Manager) Add(item *provider.Item) error {
	return c.Provider.Add(item)
}

// Get gets item from provider.
func (c Manager) Get(options provider.GetOptions) (*provider.Item, error) {
	return c.Provider.Get(options)
}

// List lists items from provider.
func (c Manager) List(options provider.ListOptions) ([]*provider.Item, error) {
	return c.Provider.List(options)
}

func getProviderByName(name *string) provider.Interface {
	commander := internal.NewCommander()

	switch *name {
	case provider.BitwardenCommand:
		return provider.Bitwarden{Commander: commander}
	case provider.OnePasswordCommand:
		return provider.OnePassword{Commander: commander}
	case provider.S3ProviderName:
		return provider.NewS3Provider()
	default:
		return provider.Bitwarden{Commander: commander}
	}
}

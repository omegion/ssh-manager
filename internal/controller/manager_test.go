package controller

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/omegion/ssh-manager/internal/provider"
	"github.com/omegion/ssh-manager/internal/provider/mocks"
)

func TestNewManager(t *testing.T) {
	t.Run("bw provider", func(t *testing.T) {
		providerName := provider.BitwardenCommand
		manager := NewManager(&providerName)

		assert.Equal(t, providerName, manager.Provider.GetName())
	})

	t.Run("op provider", func(t *testing.T) {
		providerName := provider.OnePasswordCommand
		manager := NewManager(&providerName)

		assert.Equal(t, providerName, manager.Provider.GetName())
	})

	t.Run("unknown provider", func(t *testing.T) {
		providerName := "unknown"
		manager := NewManager(&providerName)

		assert.Equal(t, provider.BitwardenCommand, manager.Provider.GetName())
	})
}

func TestManager(t *testing.T) {
	ctrl := gomock.NewController(t)
	prvMock := mocks.NewMockInterface(ctrl)

	expectedItem := provider.Item{}

	var expectedItems []*provider.Item

	t.Run("add", func(t *testing.T) {
		prvMock.EXPECT().Add(&expectedItem).Return(nil)

		manager := Manager{Provider: prvMock}

		err := manager.Add(&expectedItem)

		assert.NoError(t, err)
	})

	t.Run("get", func(t *testing.T) {
		prvMock.EXPECT().Get(gomock.Any()).Return(&expectedItem, nil)

		manager := Manager{Provider: prvMock}

		item, err := manager.Get("test")

		assert.NoError(t, err)
		assert.Equal(t, &expectedItem, item)
	})

	t.Run("get", func(t *testing.T) {
		expectedItems = append(expectedItems, &provider.Item{Name: "test"})

		prvMock.EXPECT().List().Return(expectedItems, nil)

		manager := Manager{Provider: prvMock}

		items, err := manager.List()

		assert.NoError(t, err)
		assert.Len(t, items, 1)
	})
}

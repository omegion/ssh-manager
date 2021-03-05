package bw_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/omegion/bw-ssh/pkg/bw"
	"github.com/omegion/bw-ssh/pkg/exec/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBitwarden_Sync(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockCommanderInterface(ctrl)

	m.EXPECT().Output(bw.BitwardenCommand, getSyncCommanderOptions()).Return(nil, nil)

	bitwarden := bw.Bitwarden{
		Commander: m,
	}

	err := bitwarden.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

func TestBitwarden_GetItems(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockCommanderInterface(ctrl)

	getFolderOutput, err := getGetFolderCommanderOutput()
	if err != nil {
		log.Fatal(err)
	}

	m.EXPECT().Output(bw.BitwardenCommand, getSyncCommanderOptions()).Return(nil, nil)
	m.EXPECT().Output(bw.BitwardenCommand, getGetFolderCommanderOptions()).Return(getFolderOutput, nil).Times(1)
	m.EXPECT().Output(bw.BitwardenCommand, getGetItemsCommanderOptions()).Return([]byte("[]"), nil)

	bitwarden := bw.Bitwarden{
		Commander: m,
	}

	err = bitwarden.GetItems()
	if err != nil {
		log.Fatal(err)
	}
}

func TestBitwarden_GetAndDuplicateAddFailure(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockCommanderInterface(ctrl)

	getFolderOutput, err := getGetFolderCommanderOutput()
	if err != nil {
		log.Fatal(err)
	}

	getItemsOutput, err := getGetItemCommanderOutput()
	if err != nil {
		log.Fatal(err)
	}

	m.EXPECT().Output(bw.BitwardenCommand, getSyncCommanderOptions()).Return(nil, nil).Times(1)
	m.EXPECT().Output(bw.BitwardenCommand, getGetFolderCommanderOptions()).Return(getFolderOutput, nil).Times(1)
	m.EXPECT().Output(bw.BitwardenCommand, getGetItemsCommanderOptions()).Return(getItemsOutput, nil).Times(1)

	bitwarden := bw.Bitwarden{
		Commander: m,
	}

	item := bw.Item{
		ID:       "test-ssh-id",
		Type:     1,
		Name:     "test-ssh",
		FolderID: "test-folder-id",
		Notes: []bw.Field{
			{
				Name:  "public_key",
				Value: "X",
			},
			{
				Name:  "private_key",
				Value: "Y",
			},
		},
	}

	expectedItem, err := bitwarden.Get(item.Name)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expectedItem, item)

	err = bitwarden.Add(item)

	assert.EqualError(t, err, fmt.Sprintf(": %s is already exists", item.Name))
}

func TestBitwarden_Add(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mocks.NewMockCommanderInterface(ctrl)

	getFolderOutput, err := getGetFolderCommanderOutput()
	if err != nil {
		log.Fatal(err)
	}

	item := bw.Item{
		ID:       "test-ssh-id",
		Type:     1,
		Name:     "test-ssh",
		FolderID: "test-folder-id",
		Notes: []bw.Field{
			{
				Name:  "public_key",
				Value: "X",
			},
			{
				Name:  "private_key",
				Value: "Y",
			},
		},
	}

	encodedItem, err := item.Encode()
	if err != nil {
		log.Fatal(err)
	}

	m.EXPECT().Output(bw.BitwardenCommand, getSyncCommanderOptions()).Return(nil, nil).Times(1)
	m.EXPECT().Output(bw.BitwardenCommand, getGetFolderCommanderOptions()).Return(getFolderOutput, nil).Times(1)
	m.EXPECT().Output(bw.BitwardenCommand, getGetItemsCommanderOptions()).Return([]byte("[]"), nil).Times(1)
	m.EXPECT().Output(bw.BitwardenCommand, getAddItemCommanderOptions(encodedItem)).Return(nil, nil).Times(1)

	bitwarden := bw.Bitwarden{
		Commander: m,
	}

	err = bitwarden.Add(item)
	if err != nil {
		log.Fatal(err)
	}
}

func getSyncCommanderOptions() []string {
	return []string{
		"sync",
	}
}

func getGetItemsCommanderOptions() []string {
	return []string{
		"list",
		"items",
		"--folderid",
		"test-folder-id",
	}
}

func getGetFolderCommanderOptions() []string {
	return []string{
		"list",
		"folders",
		"--search",
		bw.DefaultFolderName,
	}
}
func getAddItemCommanderOptions(encodedItem string) []string {
	return []string{
		"create",
		"item",
		encodedItem,
	}
}

func getGetFolderCommanderOutput() ([]byte, error) {
	f, err := os.Open("fixtures/list_folder_response.json")
	if err != nil {
		return []byte{}, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func getGetItemCommanderOutput() ([]byte, error) {
	f, err := os.Open("fixtures/simple_response.json")
	if err != nil {
		return []byte{}, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

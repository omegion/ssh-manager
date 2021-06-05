package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadFixture_found(t *testing.T) {
	data, err := LoadFixture("data")

	assert.NoError(t, err)
	assert.Equal(t, "yey\n", string(data))
}

func TestLoadFixture_notFound(t *testing.T) {
	data, err := LoadFixture("nodata")

	assert.Error(t, err)
	assert.Equal(t, "", string(data))
}

func TestMust_noError(t *testing.T) {
	input := []byte("Yey")

	output := Must(input, nil)

	assert.Equal(t, "Yey", string(output))
}

func TestMust_hasError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestMust_hasError should have panicked!")
		}
	}()

	_ = Must([]byte{}, FixtureFileNotFound{Path: "/path", Name: "file"})
}

package test_test

import (
	"testing"

	"github.com/omegion/ssh-manager/test"
)

func TestLoadFixture_found(t *testing.T) {
	data, err := test.LoadFixture("data")
	if err != nil {
		t.Errorf("TestLoadFixture_found should find the 'data' file")
	}

	expected := "yey\n"
	if string(data) != expected {
		t.Errorf("Expected file content: %s, got: %s", expected, data)
	}
}

func TestLoadFixture_notFound(t *testing.T) {
	data, err := test.LoadFixture("nodata")
	if err == nil {
		t.Errorf("TestLoadFixture_notFound should not find the 'nodata' file")
	}

	expected := ""
	if string(data) != expected {
		t.Errorf("Expected file content: %s, got: %s", expected, data)
	}
}

func TestMust_noError(t *testing.T) {
	input := []byte("Yey")
	expected := "Yey"

	output := test.Must(input, nil)
	if string(output) != expected {
		t.Errorf("Must() = %v, want %v", output, expected)
	}
}

func TestMust_hasError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestMust_hasError should have panicked!")
		}
	}()

	_ = test.Must([]byte{}, test.FixtureFileNotFound{Path: "/path", Name: "file"})
}

package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// LoadFixture loads the content of a fixture file.
func LoadFixture(name string) ([]byte, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", name))
	if err != nil {
		path, _ := os.Getwd()

		return []byte{}, FixtureFileNotFound{Path: path, Name: name}
	}

	return content, nil
}

// LoadFixtureReader loads the content of a fixture file and returns with an io.Reader.
func LoadFixtureReader(name string) (io.Reader, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", name))
	if err != nil {
		path, _ := os.Getwd()

		return nil, FixtureFileNotFound{Path: path, Name: name}
	}

	return bytes.NewReader(content), nil
}

// Must be content. Basically go and panic if error happened.
func Must(content []byte, err error) []byte {
	if err != nil {
		panic(err)
	}

	return content
}

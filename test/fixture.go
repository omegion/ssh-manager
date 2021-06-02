package test

import (
	"fmt"
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

// Must be content. Basically go and panic if error happened.
func Must(content []byte, err error) []byte {
	if err != nil {
		panic(err)
	}

	return content
}

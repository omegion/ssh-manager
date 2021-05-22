package test

import "fmt"

// FixtureFileNotFound occurs when the requested fixture file does not exist.
type FixtureFileNotFound struct {
	Path string
	Name string
}

func (e FixtureFileNotFound) Error() string {
	return fmt.Sprintf("Fixture file does not exist: %s/fixtures/%s", e.Path, e.Name)
}

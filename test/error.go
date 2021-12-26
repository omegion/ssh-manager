package test

import "fmt"

// FixtureFileNotFoundError occurs when the requested fixture file does not exist.
type FixtureFileNotFoundError struct {
	Path string
	Name string
}

func (e FixtureFileNotFoundError) Error() string {
	return fmt.Sprintf("Fixture file does not exist: %s/fixtures/%s", e.Path, e.Name)
}

package lpass

import "fmt"

// LastPassError occurs when LastPass command fails.
type LastPassError struct {
	Origin  error
	Message string
}

func (e LastPassError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Origin.Error())
}

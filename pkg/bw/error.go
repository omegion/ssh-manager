package bw

import "fmt"

// BitwardenError occurs when Bitwarden command fails.
type BitwardenError struct {
	Origin  error
	Message string
}

func (e BitwardenError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Origin.Error())
}

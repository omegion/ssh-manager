package provider

import "fmt"

// ProviderNotFound occurs when no provider found.
type ProviderNotFound struct {
	Name *string
}

func (e ProviderNotFound) Error() string {
	return fmt.Sprintf("no provider found for %s", *e.Name)
}

type ExecutionFailedError struct {
	Command string
	Message string
}

func (e ExecutionFailedError) Error() string {
	return fmt.Sprintf("'%s': Execution failed: %s", e.Command, e.Message)
}

type ItemAlreadyExists struct {
	Name string
}

func (e ItemAlreadyExists) Error() string {
	return fmt.Sprintf("item %s already exists", e.Name)
}

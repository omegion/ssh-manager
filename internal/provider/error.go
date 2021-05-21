package provider

import "fmt"

// NotFound occurs when no provider found.
type NotFound struct {
	Name *string
}

func (e NotFound) Error() string {
	return fmt.Sprintf("no provider found for %s", *e.Name)
}

// ExecutionFailedError occurs when an execution fails.
type ExecutionFailedError struct {
	Command string
	Message string
}

func (e ExecutionFailedError) Error() string {
	return fmt.Sprintf("'%s': Execution failed: %s", e.Command, e.Message)
}

// ItemAlreadyExists occurs when given item is not found in the provder.
type ItemAlreadyExists struct {
	Name string
}

func (e ItemAlreadyExists) Error() string {
	return fmt.Sprintf("item %s already exists", e.Name)
}

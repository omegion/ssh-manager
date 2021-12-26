package provider

import "fmt"

// NotFoundError occurs when no provider found.
type NotFoundError struct {
	Name *string
}

func (e NotFoundError) Error() string {
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

// ItemAlreadyExistsError occurs when given item is not found in the provder.
type ItemAlreadyExistsError struct {
	Name string
}

func (e ItemAlreadyExistsError) Error() string {
	return fmt.Sprintf("item %s already exists", e.Name)
}

package errors

import "fmt"

type QueryError struct {
	Err error
}

func (e *QueryError) Error() string {
	return "QueryError: " + e.Err.Error()
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

type QueryExecError struct {
	Value string
	Query string
}

func (e *QueryExecError) Error() string {
	return fmt.Sprintf("QueryExecutionEror: Error: %v Query %v", e.Value, e.Query)
}

type NotFoundError struct {
	Value string
}

func (e *NotFoundError) Error() string {
	return e.Value
}

type ModelAlreadyExists struct {
}

func (e *ModelAlreadyExists) Error() string {
	return "model already exists"
}

type LastModelDeletion struct {
}

func (e *LastModelDeletion) Error() string {
	return "cannot delete last item"
}

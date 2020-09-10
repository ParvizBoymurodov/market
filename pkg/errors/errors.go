package errors

import "fmt"

type QueryError struct {
	Query string
	Err   error
}

func (receiver *QueryError) Unwrap() error {
	return receiver.Err
}
func (receiver *QueryError) Error() string {
	return fmt.Sprintf("can't handle db operation: %v", receiver.Err.Error())
}

func QueryErrors(query string, err error) *QueryError {
	return &QueryError{Query: query, Err: err}
}
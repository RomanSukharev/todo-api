package errs

import "fmt"

type CustomError struct {
	Code    ErrorCode
	Message string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}

func New(code ErrorCode, msg string) *CustomError {
	return &CustomError{Code: code, Message: msg}
}

type ErrorCode int

const (
	CodeUnknown ErrorCode = iota
	CodeTodoNotFound
	CodeInvalidRequest
)

var (
	ErrUnknown        = New(CodeUnknown, "unknown error")
	ErrTodoNotFound   = New(CodeTodoNotFound, "todo not found")
	ErrInvalidRequest = New(CodeInvalidRequest, "invalid request")
)

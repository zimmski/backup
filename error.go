package backup

import (
	"fmt"
)

// ErrorType the error error type which is used everywhere
type ErrorType int

const (
	// AlreadyMounted mount point already mounted
	AlreadyMounted ErrorType = iota
	// AlreadyOpened tunnel already opened
	AlreadyOpened
	// FolderDoesNotExist folder does not exist
	FolderDoesNotExist
	// NotMounted mount point needs to be mounted
	NotMounted
	// NotOpened tunnel needs to be opened
	NotOpened
	// UnexpectedOutput did not expect any output from this execution
	UnexpectedOutput
)

func (e ErrorType) String() string {
	switch e {
	case AlreadyMounted:
		return "mount point already mounted"
	case AlreadyOpened:
		return "tunnel already opened"
	case FolderDoesNotExist:
		return "folder does not exist"
	case NotMounted:
		return "mount point needs to be mounted"
	case NotOpened:
		return "tunnel needs to be opened"
	case UnexpectedOutput:
		return "did not expect any output from this execution"
	default:
		return "unknown error"
	}
}

// Error the error type which is used everywhere
type Error struct {
	Type    ErrorType
	Message string
}

// NewError returns a new backup error
func NewError(errorType ErrorType, errorMessage string) *Error {
	return &Error{
		Type:    errorType,
		Message: errorMessage,
	}
}

func (e Error) Error() string {
	return e.String()
}

func (e Error) String() string {
	return fmt.Sprintf("%s - %s", e.Type, e.Message)
}

package postgres

import (
	"errors"
)

const (
	ErrorInputIsNilMessage  = "input cannot be nil"
	ErrorCarNotFoundMessage = "car not found"
)

var (
	ErrorInputIsNil  = errors.New(ErrorInputIsNilMessage)
	ErrorCarNotFound = errors.New(ErrorCarNotFoundMessage)
)

package postgres

import (
	"errors"
)

const (
	ErrorInputIsNilMessage   = "input cannot be nil"
	ErrorCarNotFoundMessage  = "car not found"
	ErrorUserNotFoundMessage = "user not found"
	ErrorUserExistsMessage   = "user already exists"
)

var (
	ErrorInputIsNil   = errors.New(ErrorInputIsNilMessage)
	ErrorCarNotFound  = errors.New(ErrorCarNotFoundMessage)
	ErrorUserNotFound = errors.New(ErrorUserNotFoundMessage)
	ErrorUserExists   = errors.New(ErrorUserExistsMessage)
)

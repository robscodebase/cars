package util

import (
	"errors"
	"log"
)

// ErrorAndLog avoids sending stack trace errors to the caller
// It prints the error and the message but only returns the message.
func ErrorAndLog(err error, msg string) error {
	log.Println(msg, err.Error())
	return errors.New(msg)
}

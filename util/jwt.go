package util

import (
	"cars/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const defaultSessionDuration = 72 * time.Hour

func GenerateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	d, err := config.Get().SessionDuration()
	if err != nil {
		return "", errors.New("session duration is not set")
	}
	if d == 0 {
		log.Printf("session duration is not set, using default: %v", defaultSessionDuration)
		d = defaultSessionDuration
	}
	claims["exp"] = time.Now().Add(d).Unix() // Token expires after 72 hours

	t, err := token.SignedString([]byte(config.Get().TokenSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

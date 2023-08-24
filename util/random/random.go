package random

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// String generates a random string of random length between 10 and 255 length
func String() string {
	minLen := 10
	maxLen := 255

	// Determine the length of the resulting string
	length, err := rand.Int(rand.Reader, big.NewInt(int64(maxLen-minLen+1)))
	if err != nil {
		log.Printf("using default value: %v, can't generate random string: %v", maxLen, err)
		length = big.NewInt(int64(maxLen))
	}
	strLen := minLen + int(length.Int64())

	result := make([]byte, strLen)
	charsetLength := big.NewInt(int64(len(charset)))
	for i := range result {
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return fmt.Sprintf("cant create random string %d", i)
		}
		result[i] = charset[index.Int64()]
	}

	return string(result)
}

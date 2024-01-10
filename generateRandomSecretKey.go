package main

import (
	"crypto/rand"
	"fmt"
)

func generateRandomSecretKey() (string, error) {
	// Generate a random secret key
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", key), nil
}

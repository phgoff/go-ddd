package storage

import (
	"crypto/rand"
	"fmt"
)

func GetID(prefix string) (string, error) {
	gen := make([]byte, 8)

	_, err := rand.Read(gen)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s_%x", prefix, gen), nil
}

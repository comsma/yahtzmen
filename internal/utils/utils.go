package utils

import (
	"github.com/jaevor/go-nanoid"
)

func GenerateId() (string, error) {
	gen, err := nanoid.Standard(20)
	if err != nil {
		return "", err

	}

	id := gen()

	return id, nil
}

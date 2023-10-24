package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Genrate 16 characters Nanoid

func Nanoid() (string, error) {
	id, err := gonanoid.New(16)
	if err != nil {
		panic(err)
	}
	return id, nil
}

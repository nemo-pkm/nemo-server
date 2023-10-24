package utils

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Genrate Nanoid

func Nanoid() (string, error) {
	id, err := gonanoid.New()
	if err != nil {
		fmt.Println("Generate nanoid error!")
		return "", fmt.Errorf("faild to generate Nanoid: %w", err)
	}
	return id, nil
}

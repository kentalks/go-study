package utils

import (
	"errors"
)

type error interface {
	Error() string
}

func Process(count int) error {
	if count < 1 {
		return errors.New("Invalid count")
	}
	return nil
}

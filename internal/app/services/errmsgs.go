package services

import (
	"errors"
)

var (
	ErrorRNGFail = errors.New("failed to generate random sequence")
)

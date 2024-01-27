package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

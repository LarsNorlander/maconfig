package action

import (
	"errors"
	"fmt"
)

var (
	ErrUnexpectedParameterType = errors.New("unexpected parameter type")
	ErrInvalidControl          = errors.New("invalid control type")
)

type notEnoughValuesError struct {
	expected, provided int
}

func (e *notEnoughValuesError) Error() string {
	return fmt.Sprintf("not enough values provided; expected %d, provided %d", e.expected, e.provided)
}

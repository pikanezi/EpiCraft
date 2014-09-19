package entities

import "fmt"

var (
	// ErrCannotMove is the error when trying to move when
	ErrCannotMove = fmt.Errorf("cannot move")

	// ErrTooFarAway is the error when the target is too far away.
	ErrTooFarAway = fmt.Errorf("target too far away")
)

package sdk

import "errors"

// Error error
type Error struct {
}

func (e Error) Error() string {
	return ""
}

// ErrNotImplement error not implement
var ErrNotImplement = errors.New("not implement")

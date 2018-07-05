package httpErrorCodes

import (
	"errors"
)

// TODO: think about error code struct

var (
	// ErrCode is a config or an internal error
	ErrCode = errors.New("Case statement in code is not correct.")

	// ErrNoResult is a not results error
	ErrNoResult = errors.New("Result not found.")

	// ErrUnavailable is a database not available error
	ErrUnavailable = errors.New("Database is unavailable.")

	// ErrUnauthorized is a permissions violation
	ErrUnauthorized = errors.New("User does not have permission to perform this operation.")
)

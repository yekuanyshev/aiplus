package repository

import "errors"

var (
	ErrInvalidQuery = errors.New("invalid query")
	ErrNotFound     = errors.New("not found")
)

package repository

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrConflict       = errors.New("resource already exists")
	ErrDuplicateOwner = errors.New("this owner already exists")
)

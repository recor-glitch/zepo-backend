package domain

import "errors"

var (
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrRecordNotFound = errors.New("record not found")
	ErrInternal       = errors.New("internal error")
)

package settings

import "errors"

var (
	ErrTaskNameRequired    = errors.New("task name is required")
	ErrIdRequired          = errors.New("id is required")
	ErrAccountNameRequired = errors.New("account name is required")
	ErrFormatRequired      = errors.New("format is required")
)

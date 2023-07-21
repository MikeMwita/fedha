package errors

import "errors"

var (
	ErrDBPortNotSet     = errors.New("database port not set")
	ErrDbPasswordNotSet = errors.New("database password not set")
	ErrDBUsernameNotSet = errors.New("database username not set")
	ErrDBHostNotSet     = errors.New("database host not set")
)

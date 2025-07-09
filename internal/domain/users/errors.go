package users

import "errors"

var (
	ErrInvalidUserCreads = errors.New("invalid creads")
	ErrUserAlreadyExists = errors.New("user alreadi exists")
	ErrUserNotFound      = errors.New("user not found")
)

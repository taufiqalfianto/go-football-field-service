package error

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserAlreadyExists     = errors.New("user already exists")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrInvalidToken          = errors.New("invalid token")
	ErrPasswordIncorrect     = errors.New("password incorrect")
	ErrPasswordDoesNotMatch  = errors.New("password does not match")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

var UserError = []error{
    ErrUserNotFound,
    ErrUserAlreadyExists,
    ErrInvalidCredentials,
    ErrInvalidToken,
    ErrPasswordIncorrect,
    ErrPasswordDoesNotMatch,
    ErrUsernameAlreadyExists,
}
package apperrors

import "errors"

type Kind string
type Code string

const (
	KindBadRequest   Kind = "bad_request"
	KindUnauthorized Kind = "unauthorized"
	KindNotFound     Kind = "not_found"
	KindConflict     Kind = "conflict"
	KindInternal     Kind = "internal"
)

const (
	CodeInvalidRequest     Code = "invalid_request"
	CodeInvalidCredentials Code = "invalid_credentials"
	CodeUserNotFound       Code = "user_not_found"
	CodeEmailAlreadyExists Code = "email_already_exists"
	CodeInternal           Code = "internal_error"
)

type Error struct {
	Kind Kind
	Code Code
	Err  error
}

func (e *Error) Error() string {
	return string(e.Code)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func New(kind Kind, code Code, err error) *Error {
	return &Error{
		Kind: kind,
		Code: code,
		Err:  err,
	}
}

func BadRequest(code Code, err error) *Error {
	return New(KindBadRequest, code, err)
}

func Unauthorized(code Code, err error) *Error {
	return New(KindUnauthorized, code, err)
}

func NotFound(code Code, err error) *Error {
	return New(KindNotFound, code, err)
}

func Conflict(code Code, err error) *Error {
	return New(KindConflict, code, err)
}

func Internal(err error) *Error {
	return New(KindInternal, CodeInternal, err)
}

func From(err error) (*Error, bool) {
	var appErr *Error
	if errors.As(err, &appErr) {
		return appErr, true
	}

	return nil, false
}

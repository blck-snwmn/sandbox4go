package error

import (
	"github.com/blck-snwmn/sandbox4go/20200905_error/code"
	"golang.org/x/xerrors"
)

var _ errResponse = (*ErrResponse)(nil)

type response interface {
	Code() string
	Message() string
}
type errResponse interface {
	response
	xerrors.Wrapper
	xerrors.Formatter
	As(interface{}) bool
	Is(error) bool
}

// ErrResponse is ...
type ErrResponse struct {
	code    string
	message string
	Err     error
}

func (*ErrResponse) Code() string {
	return ""
}
func (*ErrResponse) Message() string {
	return ""
}
func (*ErrResponse) Error() string {
	return ""
}
func (er *ErrResponse) As(err interface{}) bool {
	if other, ok := err.(**ErrResponse); ok {
		(*other).code = er.code
		(*other).message = er.message
		(*other).Err = er.Err
	}
	return true
}
func (er *ErrResponse) Is(err error) bool {
	if ec, ok := err.(*code.ErrCode); ok {
		return er.code == ec.Code()
	}
	if other, ok := err.(*ErrResponse); ok {
		return er.code == other.code && er.message == other.message
	}
	return false
}

func (*ErrResponse) FormatError(p xerrors.Printer) (next error) {
	return nil
}

func (*ErrResponse) Unwrap() error {
	return nil
}

const (
	Expired        = "Expired"
	Insufficient   = "Insufficient"
	Mismatch       = "Mismatch"
	FailedRegister = "Failed to register"
)

var (
	ErrExpired = &ErrResponse{
		code:    code.CodeNotAllowed,
		message: Expired,
	}
	ErrInsufficient = &ErrResponse{
		code:    code.CodeNotAllowed,
		message: Insufficient,
	}
	ErrMismatch = &ErrResponse{
		code:    code.CodeInvalid,
		message: Mismatch,
	}
	ErrRegisterFailed = &ErrResponse{
		code:    code.CodeFailed,
		message: FailedRegister,
	}
)

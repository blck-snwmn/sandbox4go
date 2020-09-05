package error

import (
	"testing"

	"github.com/blck-snwmn/sandbox4go/20200905_error/code"
	"golang.org/x/xerrors"
)

func TestErrResponse_Is(t *testing.T) {
	err := xerrors.Errorf("oops!: %w", ErrExpired)
	if !xerrors.Is(err, ErrExpired) {
		t.Errorf("wrap error")
	}
	sameErr := &ErrResponse{
		code:    ErrExpired.code,
		message: ErrExpired.message,
	}
	if !xerrors.Is(err, sameErr) {
		t.Errorf("error that have same code, message is equal, but not equal")
	}

	differentMessageErr := &ErrResponse{
		code:    ErrExpired.code,
		message: "test",
	}
	if xerrors.Is(err, differentMessageErr) {
		t.Errorf("error that have different message is not equal, but equal")
	}

	if !xerrors.Is(err, code.ErrNotAllowed) {
		t.Errorf("ErrCode that have same code is equal, but not equal")
	}
	xerrors.Is(err, nil)
}

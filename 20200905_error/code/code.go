package code

var _ error = (*ErrCode)(nil)

// ErrCode is ...
type ErrCode struct {
	code string
}

func (ec *ErrCode) Code() string {
	return ec.code
}
func (*ErrCode) Error() string {
	return ""
}
func (*ErrCode) As(interface{}) bool {
	return true
}
func (*ErrCode) Is(error) bool {
	return true
}

const (
	CodeNotAllowed = "01"
	CodeInvalid    = "02"
	CodeFailed     = "03"
)

var (
	ErrNotAllowed = &ErrCode{CodeNotAllowed}
	ErrInvalid    = &ErrCode{CodeInvalid}
	ErrFailed     = &ErrCode{CodeFailed}
)

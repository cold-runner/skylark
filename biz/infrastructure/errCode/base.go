package errCode

type BizErrCode int32
type BizError struct {
	HttpStatus int
	Msg        string
}

const (
	ErrSuccess BizErrCode = iota + 100001
	ErrUnknown
	ErrBind
	ErrValidation
	ErrNotFound
	ErrSmsCode
)

const (
	ErrTokenInvalid = iota + 100101
	ErrPermissionDenied
)

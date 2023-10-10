package code

import (
	"github.com/marmotedu/errors"
	"github.com/novalagung/gubrak"
	"net/http"
)

// ErrCode implements `github.com/marmotedu/errors`.Coder interface.
type ErrCode struct {
	// C refers to the code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Ext string

	// Ref specify the reference document.
	Ref string
}

// 编译器检查
var _ errors.Coder = &ErrCode{}

// Code returns the integer code of ErrCode.
func (c ErrCode) Code() int {
	return c.C
}

// String implements stringer. String returns the external error message,
// if any.
func (c ErrCode) String() string {
	return c.Ext
}

// Reference returns the reference document.
func (c ErrCode) Reference() string {
	return c.Ref
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (c ErrCode) HTTPStatus() int {
	if c.HTTP == 0 {
		return http.StatusInternalServerError
	}

	return c.HTTP
}

func register(code int, httpStatus int, message string, refs ...string) {
	found, _ := gubrak.Includes([]int{200, 400, 401, 403, 404, 500}, httpStatus)
	if !found {
		panic("http code not in `200, 400, 401, 403, 404, 500`")
	}

	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}

	coder := &ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  reference,
	}

	errors.MustRegister(coder)
}

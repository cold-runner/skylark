package code

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/marmotedu/errors"
	"net/http"
)

// Response defines project response format which in marmotedu organization.
type Response struct {
	Code      int         `json:"code,omitempty"`
	Message   string      `json:"message,omitempty"`
	Reference string      `json:"reference,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

// WriteResponse used to write an error and JSON data into response.
func WriteResponse(ctx *app.RequestContext, err error, data interface{}) {
	if err != nil {
		coder := errors.ParseCoder(err)

		ctx.JSON(coder.HTTPStatus(), Response{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
			Data:      data,
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{Data: data})
}

// WriteExtraResponse used to write an error and JSON data into response.
func WriteExtraResponse(ctx *app.RequestContext, err error, extra string, data interface{}) {
	if err != nil {
		coder := errors.ParseCoder(err)

		ctx.JSON(coder.HTTPStatus(), Response{
			Code:      coder.Code(),
			Message:   coder.String() + "," + extra,
			Reference: coder.Reference(),
			Data:      data,
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{Data: data})
}

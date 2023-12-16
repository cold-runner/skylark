package errCode

import "github.com/cloudwego/hertz/pkg/protocol/consts"

var ErrMap = map[BizErrCode]*BizError{
	// 通用错误
	ErrSuccess:            {consts.StatusOK, "ok"},
	ErrUnknown:            {consts.StatusInternalServerError, "server internal error"},
	ErrValidation:         {consts.StatusBadRequest, "validation error"},
	ErrBind:               {consts.StatusInternalServerError, "error occurred while binding the request body"},
	ErrNotFound:           {consts.StatusNotFound, "resource not found"},
	ErrSmsCode:            {consts.StatusBadRequest, "sms code invalid"},
	ErrAlreadySendSmsCode: {consts.StatusBadRequest, "already send smsCode"},

	ErrTokenInvalid:     {consts.StatusUnauthorized, "token invalid"},
	ErrPermissionDenied: {consts.StatusForbidden, "permission denied"},

	// 用户相关
	ErrPasswordIncorrect: {consts.StatusBadRequest, "password is incorrect"},
	ErrUserNotFound:      {consts.StatusBadRequest, "user not found"},
	ErrUserAlreadyExist:  {consts.StatusBadRequest, "user already exist"},

	// 文章相关
}

var ErrBackendMap = map[BizErrCode]*BizError{
	// 文章相关
	ErrSync:        {consts.StatusInternalServerError, "error occurred when sync"},
	ErrSearchPosts: {consts.StatusInternalServerError, "error occurred when search posts"},
}

package entity

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/biz/infrastructure/cache"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	stdErr "github.com/pkg/errors"
)

type SmsCode struct {
}

func (s *SmsCode) Validate(c context.Context, ctx *app.RequestContext, cacheIns cache.Cache, phone, recvSmsCode string) *errors.Error {
	smsCode, err := cacheIns.Get(c, phone)
	if err != nil {
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}

	smsCodeStr := smsCode.(string)
	if smsCodeStr != recvSmsCode {
		errMsg := "smsCode incorrect! received: " + recvSmsCode + " stored: " + smsCodeStr
		return errCode.WrapBizErr(ctx, stdErr.New(errMsg), errCode.ErrSmsCode)
	}

	return nil
}

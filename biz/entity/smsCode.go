package entity

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/biz/infrastructure/cache"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/sms"
	"github.com/cold-runner/skylark/biz/util"
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

func (s *SmsCode) AlreadySend(c context.Context, ctx *app.RequestContext, cacheIns cache.Cache, phone string) error {
	isExpired, err := cacheIns.IsExpired(c, phone)
	if err != nil {
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	if !isExpired {
		return errCode.WrapBizErr(ctx, stdErr.New("already send smsCode"), errCode.ErrAlreadySendSmsCode)
	}
	return nil
}

func (s *SmsCode) SendSmsCode(c context.Context, ctx *app.RequestContext, smsClient sms.Sms, phone, smsCode string, expireTime int) *errors.Error {
	if err := smsClient.SendToSingle(c, phone, []string{smsCode, strconv.Itoa(expireTime)}); err != nil {
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return nil
}

func (s *SmsCode) SetCache(c context.Context, ctx *app.RequestContext, cacheIns cache.Cache, phone, smsCode string, expireTime time.Duration) *errors.Error {
	if err := cacheIns.SetExpiration(c, phone, smsCode, expireTime); err != nil {
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return nil
}

func (s *SmsCode) GenSmsCodeRandom(ctx *app.RequestContext, count int) (string, *errors.Error) {
	smsCode, err := util.RandSmsCode(count)
	if err != nil {
		return "", errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return smsCode, nil
}

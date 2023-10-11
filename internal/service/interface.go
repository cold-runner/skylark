package service

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/oss"
	"github.com/cold-runner/skylark/internal/pkg/sms"
	"github.com/cold-runner/skylark/internal/store"
	"time"
)

// Factory 抽象工厂
type Factory interface {
	NewInstance(cacheIns cache.Cache, ossIns oss.Oss, smsClient sms.Sms, storeIns store.Store) Interface
}

// Interface service层接口
type Interface interface {
	ArticleSrv
	CommentSrv
	LarkSrv
	PublicSrv
	InternalMethod
}

type LarkSrv interface {
}

type CommentSrv interface {
}

type ArticleSrv interface {
}

type PublicSrv interface {
	Register(c context.Context, register *user.Register) error
	SendRegisterSms(c context.Context, phone string, paramSet []string) error
}

type InternalMethod interface {
	SetExpiration(c context.Context, key, value string, expiration time.Duration) error
	GetSmsCode(c context.Context, key string) (value string, err error)
	ValidateSmsCode(c context.Context, phone, inCache, userPass string) error
}

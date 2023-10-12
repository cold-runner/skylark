package service

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/cold-runner/skylark/internal/pkg/cache"
	"github.com/cold-runner/skylark/internal/pkg/oss"
	"github.com/cold-runner/skylark/internal/pkg/sms"
	"github.com/cold-runner/skylark/internal/store"
)

// Factory 抽象工厂
type Factory interface {
	NewInstance(cacheIns cache.Cache, ossIns oss.Oss, smsClient sms.Sms, storeIns store.Store) Interface
}

// Interface service层接口
type Interface interface {
	ArticleSrv
	AuthSrv
	CommentSrv
	LarkSrv
	PublicSrv
}

type AuthSrv interface {
	ProcessLogin(c context.Context, loginUser *user.LoginUser) (loggedUser *user.LoggedUser, err error)
}

type LarkSrv interface {
}

type CommentSrv interface {
}

type ArticleSrv interface {
}

type PublicSrv interface {
	Register(c context.Context, register *user.Register) error
	SendSms(c context.Context, phone string) error
}

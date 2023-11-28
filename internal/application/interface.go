package application

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/domain/entity"
	"github.com/cold-runner/skylark/internal/infrastructure/cache"
	"github.com/cold-runner/skylark/internal/infrastructure/oss"
	"github.com/cold-runner/skylark/internal/infrastructure/sms"
	"github.com/cold-runner/skylark/internal/infrastructure/store"
)

// Factory 抽象工厂
type Factory interface {
	NewInstance(cacheIns cache.Cache, ossIns oss.Oss, smsClient sms.Sms, storeIns store.Store) Interface
}

// Interface service层接口，向controller层暴露
type Interface interface {
	PostSrv
	Middleware
	CommentSrv
	LarkSrv
	PublicSrv
}

type Middleware interface {
	Authenticator(c context.Context, ctx *app.RequestContext, loginUser *entity.LoginUser) (loggedUser *entity.LoggedInUser, err error)
}

type LarkSrv interface {
	GetLarkInfo(c context.Context, ctx *app.RequestContext, stuNum string) (*entity.LoggedInUser, error)
	BindQq(c context.Context, ctx *app.RequestContext, stuNum, qqUnionId string) error
}

type CommentSrv interface {
}

type PostSrv interface {
}

type PublicSrv interface {
	Register(c context.Context, ctx *app.RequestContext, register *entity.Register) error
	SendSmsCode(c context.Context, ctx *app.RequestContext, phone string) error
}

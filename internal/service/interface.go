package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/internal/model/post"
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

// Interface service层接口，向controller层暴露
type Interface interface {
	PostSrv
	Middleware
	CommentSrv
	LarkSrv
	PublicSrv
}

type Middleware interface {
	Authenticator(c context.Context, ctx *app.RequestContext, loginUser *user.LoginUser) (loggedUser *user.LoggedUser, err *errors.Error)
}

type LarkSrv interface {
	BindQq(c context.Context, ctx *app.RequestContext, stuNum, qqUnionId string) *errors.Error
}

type CommentSrv interface {
}

type PostSrv interface {
	CreateDraft(c context.Context, ctx *app.RequestContext) *errors.Error
	SaveDraft(c context.Context, ctx *app.RequestContext, userId string, draft *post.DraftInfo) *errors.Error
}

type PublicSrv interface {
	Register(c context.Context, ctx *app.RequestContext, register *user.Register) *errors.Error
	SendSmsCode(c context.Context, ctx *app.RequestContext, phone string) *errors.Error
}

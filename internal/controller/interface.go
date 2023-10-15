package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/service"
	"github.com/hertz-contrib/jwt"
)

// Factory 抽象工厂
type Factory interface {
	NewInstance(serviceIns service.Interface) Interface
}

// Interface controller层接口
type Interface interface {
	PostController
	CommentController
	LarkController
	PublicController
}

type PostController interface {
	Save(context.Context, *app.RequestContext)
}

type LarkController interface {
	BindQq(context.Context, *app.RequestContext)
	ChangePassword(context.Context, *app.RequestContext)
}

type CommentController interface {
}

type PublicController interface {
	Jwt() *jwt.HertzJWTMiddleware
	Register(context.Context, *app.RequestContext)
	SendSms(context.Context, *app.RequestContext)
}

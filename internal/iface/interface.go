package iface

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/application"
	"github.com/hertz-contrib/jwt"
)

// Factory 抽象工厂
type Factory interface {
	NewInstance(serviceIns application.Interface) UserInterface
}

// UserInterface 用户接口层接口
type UserInterface interface {
	PostInterface
	CommentInterface
	LarkInterface
	PublicInterface
}

type PostInterface interface {
}

type LarkInterface interface {
	BindQq(context.Context, *app.RequestContext)
	GetLarkInfo(context.Context, *app.RequestContext)
}

type CommentInterface interface {
}

type PublicInterface interface {
	Jwt() *jwt.HertzJWTMiddleware
	Register(context.Context, *app.RequestContext)
	SendSms(context.Context, *app.RequestContext)
}

package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/service"
)

// Factory 抽象工厂
type Factory interface {
	NewInstance(context context.Context, serviceIns service.Interface) Interface
}

// Interface controller层接口
type Interface interface {
	ArticleController
	CommentController
	LarkController
	PublicController
}

type ArticleController interface {
}

type LarkController interface {
}

type CommentController interface {
}

type PublicController interface {
	Register(context.Context, *app.RequestContext)
	SendSms(context.Context, *app.RequestContext)
}

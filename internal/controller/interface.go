package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// Controller create a controller used to handle request for user resource.
type Controller interface {
	ArticleController
	CommentController
	LarkController
	PublicController
}

type ArticleController interface {
}

type LarkController interface {
	Register(context.Context, *app.RequestContext)
}

type CommentController interface {
}

type PublicController interface {
	SendSms(context.Context, *app.RequestContext)
}

package core

import (
	"context"
	"github.com/cold-runner/skylark/internal/pkg/code"
	bizErr "github.com/marmotedu/errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/controller"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/marmotedu/iam/pkg/log"
)

func (a *Application) InstallRouter() *Application {
	// 注册自定义校验规则
	controller.Validator()

	// 注册中间件
	a.router.Use(gzip.Gzip(gzip.DefaultCompression), accesslog.New())
	jwt := a.controllerIns.Jwt()

	// 非鉴权路由
	{
		publicRouter := a.router.Group("")
		publicRouter.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
			log.L(c).Info("健康测试通过！")
			code.WriteResponse(ctx, bizErr.WithCode(code.ErrUnknown, "", nil), nil)
		})
		publicRouter.POST("/share", jwt.LoginHandler)
		publicRouter.GET("/sendSms", a.controllerIns.SendSms)
		publicRouter.POST("/register", a.controllerIns.Register)
	}

	// 鉴权路由
	{
		authRouter := a.router.Group("/auth")
		authRouter.GET("/refresh_token", jwt.RefreshHandler)
		// 用户功能路由
		{
			larkRouter := authRouter.Group("/lark")
			larkRouter.GET("/bindQq", a.controllerIns.BindQq)
		}

		// 文章功能路由
		{
			//postRouter := authRouter.Group("/post")
			//postRouter.GET("")
		}

		// 评论功能路由
		{
			//commentRouter := authRouter.Group("/comment")
			//commentRouter.GET("")
		}
	}

	return a
}

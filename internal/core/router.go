package core

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/limiter"
)

func (a *Application) InstallRouter() *Application {
	// 注册中间件
	a.router.Use(gzip.Gzip(gzip.DefaultCompression), limiter.AdaptiveLimit())
	jwt := a.controllerIns.Jwt()
	// pprof
	//pprof.Register(a.router)
	// 非鉴权路由
	{
		publicRouter := a.router.Group("")
		publicRouter.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
			hlog.Debug("调用/ping路由")
			errCode.ResponseOk(ctx, nil)
		})
		publicRouter.POST("/login", jwt.LoginHandler)
		publicRouter.GET("/sendSms", a.controllerIns.SendSms)
		publicRouter.POST("/register", a.controllerIns.Register)
	}

	// 鉴权路由
	{
		authRouter := a.router.Group("/auth")
		authRouter.GET("/refresh_token", jwt.RefreshHandler)
		authRouter.Use(jwt.MiddlewareFunc())
		// 用户功能路由
		{
			larkRouter := authRouter.Group("/lark")
			larkRouter.GET("/bindQq", a.controllerIns.BindQq)
		}

		// 文章功能路由
		{
			postRouter := authRouter.Group("/post")
			postRouter.POST("/createDraft", a.controllerIns.CreateDraft)
			postRouter.POST("/save", a.controllerIns.Save)
		}

		// 评论功能路由
		{
			//commentRouter := authRouter.Group("/comment")
			//commentRouter.GET("")
		}
	}

	return a
}

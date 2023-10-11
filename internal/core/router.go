package core

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/skylark/internal/model/user"
	"github.com/marmotedu/iam/pkg/log"
)

func (a *Application) larkRouter() {
	// 注册自定义校验规则
	user.CheckRegister()

}

func (a *Application) articleRouter() {

}

func (a *Application) commentRouter() {

}

func (a *Application) publicRouter() {
	publicRouter := a.router.Group("/noAuth")

	publicRouter.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		log.L(c).Info("健康测试通过！")
		ctx.JSON(consts.StatusOK, utils.H{"msg": "pong!"})
	})
	publicRouter.GET("/sendSms", a.controllerIns.SendSms)
	publicRouter.POST("/register", a.controllerIns.Register)
}

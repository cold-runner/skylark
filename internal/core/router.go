package core

import (
	"github.com/cold-runner/skylark/internal/model/user"
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

	publicRouter.GET("/sendSms", a.controller.SendSms)
	publicRouter.POST("/register", a.controller.Register)
}

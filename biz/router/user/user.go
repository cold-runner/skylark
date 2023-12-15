// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "github.com/cold-runner/skylark/biz/handler/user"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/getLark", append(_getuserinfobystunumMw(), user.GetUserInfoByStuNum)...)
	root.POST("/register", append(_registerMw(), user.Register)...)
	root.GET("/sendSmsCode", append(_sendsmscodeMw(), user.SendSmsCode)...)
	{
		_login := root.Group("/login", _loginMw()...)
		_login.GET("/password", append(_passwordloginMw(), user.PasswordLogin)...)
		_login.GET("/phone", append(_phoneloginMw(), user.PhoneLogin)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/getInfoById", append(_getuserinfobyidMw(), user.GetUserInfoById)...)
	}
}

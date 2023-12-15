// Code generated by hertz generator.

package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/biz/handler/user"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/limiter"
)

func rootMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		gzip.Gzip(gzip.DefaultCompression),
		limiter.AdaptiveLimit(),
	}
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _sendsmscodeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _passwordloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _phoneloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserinfoMw() []app.HandlerFunc {
	return []app.HandlerFunc{user.JwtMiddleware.MiddlewareFunc()}
}

func _getuserinfobystunumMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserinfobyidMw() []app.HandlerFunc {
	// your code...
	return nil
}

// Code generated by hertz generator.

package comment

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/biz/handler/user"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentpostMw() []app.HandlerFunc {
	return []app.HandlerFunc{user.JwtMiddleware.MiddlewareFunc()}
}

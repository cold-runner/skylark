package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/model/post"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/cold-runner/skylark/internal/pkg/config"
	"github.com/marmotedu/errors"
)

func (c *controllerV1) Save(ctx context.Context, context *app.RequestContext) {
	draft := &post.Draft{}
	if err := context.BindAndValidate(draft); err != nil {
		code.WriteResponse(context, errors.WithCode(code.ErrValidation, "", nil), nil)
		return
	}
	jwtConfig := config.GetConfig().JwtConfig()
	userInfo := context.Value(jwtConfig.IdentityKey).(map[string]interface{})
	userId := userInfo["UserId"].(string)

	if err := c.serviceIns.SaveDraft(ctx, userId, draft); err != nil {
		code.WriteResponse(context, err, nil)
		return
	}
	code.WriteResponse(context, errors.WithCode(code.ErrSuccess, "", nil), nil)
}

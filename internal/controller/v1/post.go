package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/model/post"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
)

func (c *controllerV1) CreateDraft(ctx context.Context, context *app.RequestContext) {
	err := c.serviceIns.CreateDraft(ctx, context)
	if err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, nil)
}

func (c *controllerV1) Save(ctx context.Context, context *app.RequestContext) {
	// TODO 参数校验
	draftInfo := &post.DraftInfo{}
	if err := context.BindAndValidate(draftInfo); err != nil {
		errCode.ResponseWithError(context, errCode.ErrValidation, nil)
		return
	}
	userInfo := context.Value("identity").(map[string]interface{})
	userId := userInfo["UserId"].(string)

	// 移交服务层
	if err := c.serviceIns.SaveDraft(ctx, context, userId, draftInfo); err != nil {
		errCode.ResponseError(context, nil)
		return
	}

	errCode.ResponseOk(context, nil)
}

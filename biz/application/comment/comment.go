package comment

import (
	"context"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"

	entity "github.com/cold-runner/skylark/biz/entity/comment"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/model/comment"

	"github.com/cloudwego/hertz/pkg/app"
)

func Comment(c context.Context, ctx *app.RequestContext, req *comment.CommentPostReq) (*comment.CommentPostResp, error) {
	storeIns := store.GetIns()

	ety := &entity.Comment{}
	// TODO 做一些合法性校验
	if err := ety.Store(c, ctx, storeIns, req); err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return &comment.CommentPostResp{
		Status: errCode.SuccessStatus,
	}, nil
}

package comment

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"github.com/cold-runner/skylark/biz/model/comment"
	"github.com/google/uuid"
)

type Comment struct {
}

func (co *Comment) Store(c context.Context, ctx *app.RequestContext, storeIns store.Store, req *comment.CommentPostReq) *errors.Error {
	ov := &orm_gen.Comment{
		ID:       uuid.New(),
		PostID:   req.PostId,
		UserID:   req.UserId,
		ParentID: req.ParentId,
		Content:  req.Content,
	}
	err := storeIns.CreateComment(c, ov)
	if err != nil {
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return nil
}

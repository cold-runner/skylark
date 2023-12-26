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

func (co *Comment) Store(c context.Context, ctx *app.RequestContext, req *comment.CommentPostReq) *errors.Error {
	storeIns := store.GetIns()

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

func (co *Comment) Delete(c context.Context, ctx *app.RequestContext, req *comment.DeleteCommentReq) *errors.Error {
	storeIns := store.GetIns()
	err := storeIns.DeleteComment(c, req.CommentId)
	if err != nil {
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return nil
}

func (co *Comment) Update(c context.Context, ctx *app.RequestContext, req *comment.UpdateCommentReq) *errors.Error {
	storeIns := store.GetIns()
	err := storeIns.UpdateComment(c, req.CommentId, req.Content)
	if err != nil {
		return errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return nil
}

func (co *Comment) Query(c context.Context, ctx *app.RequestContext, req *comment.GetCommentReq) ([]*comment.ArticleCommentInfo, *errors.Error) {
	storeIns := store.GetIns()
	stored, err := storeIns.QueryComment(c, req.PostId, int(req.Cursor), int(req.Limit))
	if err != nil {
		return nil, errCode.WrapBackendErr(ctx, err, errCode.ErrUnknown)
	}
	return format(stored), nil
}

func (co *Comment) QueryReply(c context.Context, ctx *app.RequestContext, req *comment.GetReplyCommentReq) ([]*comment.ArticleCommentInfo, *errors.Error) {
	storeIns := store.GetIns()
	stored, err := storeIns.QueryReplyComment(c, req.PostId, req.ParentId, int(req.Cursor), int(req.Limit))
	if err != nil {
		return nil, errCode.WrapBackendErr(ctx, err, errCode.ErrUnknown)
	}
	return format(stored), nil
}

func format(vd []*orm_gen.Comment) []*comment.ArticleCommentInfo {
	res := make([]*comment.ArticleCommentInfo, len(vd))
	for i := range res {
		res[i] = &comment.ArticleCommentInfo{}
		res[i].Id = vd[i].ID.String()
		res[i].CreatedAt = vd[i].CreatedAt.String()
		res[i].PostId = vd[i].PostID
		res[i].UserId = vd[i].UserID
		res[i].ParentId = vd[i].ParentID
		res[i].Like = int32(vd[i].Like)
		res[i].Content = vd[i].Content
	}

	return res
}

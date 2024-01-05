package comment

import (
	"context"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"

	entity "github.com/cold-runner/skylark/biz/entity/comment"
	"github.com/cold-runner/skylark/biz/model/comment"

	"github.com/cloudwego/hertz/pkg/app"
)

func Comment(c context.Context, ctx *app.RequestContext, req *comment.CommentPostReq) (*comment.CommentPostResp, error) {
	ety := &entity.Comment{}
	// TODO 做一些合法性校验
	if err := ety.Store(c, ctx, req); err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return &comment.CommentPostResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
	}, nil
}

func DeleteComment(c context.Context, ctx *app.RequestContext, req *comment.DeleteCommentReq) (*comment.DeleteCommentResp, error) {
	ety := &entity.Comment{}
	if err := ety.Delete(c, ctx, req); err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return &comment.DeleteCommentResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
	}, nil
}

func UpdateComment(c context.Context, ctx *app.RequestContext, req *comment.UpdateCommentReq) (interface{}, error) {
	ety := &entity.Comment{}
	if err := ety.Update(c, ctx, req); err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return &comment.UpdateCommentResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
	}, nil
}

func GetComments(c context.Context, ctx *app.RequestContext, req *comment.GetCommentReq) (*comment.GetCommentResp, error) {
	ety := &entity.Comment{}
	comments, err := ety.Query(c, ctx, req)
	if err != nil {
		return nil, err
	}

	return &comment.GetCommentResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
		Data: &comment.GetCommentResp_Data{Comments: comments},
	}, nil
}

func GetReplyComments(c context.Context, ctx *app.RequestContext, req *comment.GetReplyCommentReq) (*comment.GetReplyCommentResp, error) {
	ety := &entity.Comment{}
	comments, err := ety.QueryReply(c, ctx, req)
	if err != nil {
		return nil, err
	}

	return &comment.GetReplyCommentResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
		Data: &comment.GetReplyCommentResp_Data{Comments: comments},
	}, nil
}

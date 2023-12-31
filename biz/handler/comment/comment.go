// Code generated by hertz generator.

package comment

import (
	"context"
	service "github.com/cold-runner/skylark/biz/application/comment"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	comment "github.com/cold-runner/skylark/biz/model/comment"
)

// CreateComment .
// @router /comment/create [POST]
func CreateComment(ctx context.Context, c *app.RequestContext) {
	routerPath := string(c.URI().Path())

	hlog.Debugf(log.ROUTE_PATH, routerPath)
	var err error
	var req comment.CommentPostReq
	err = c.BindAndValidate(&req)
	if err != nil {
		errCode.ResponseValidationFailed(c, err)
		return
	}

	resp, err := service.Comment(ctx, c, &req)
	if err != nil {
		errCode.ResponseFailed(c)
		hlog.Warnf(log.REQUEST_FAILED+log.EXTRA_ERROR_INFO, routerPath, c.Errors.Last())
		return
	}

	hlog.Debugf(log.REQUEST_SUCCESSFUL, routerPath)
	c.JSON(consts.StatusOK, resp)
}

// DeleteComment .
// @router /comment/delete [DELETE]
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	routerPath := string(c.URI().Path())

	hlog.Debugf(log.ROUTE_PATH, routerPath)
	var err error
	var req comment.DeleteCommentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		errCode.ResponseValidationFailed(c, err)
		return
	}

	resp, err := service.DeleteComment(ctx, c, &req)
	if err != nil {
		errCode.ResponseFailed(c)
		hlog.Warnf(log.REQUEST_FAILED+log.EXTRA_ERROR_INFO, routerPath, c.Errors.Last())
		return
	}

	hlog.Debugf(log.REQUEST_SUCCESSFUL, routerPath)
	c.JSON(consts.StatusOK, resp)
}

// UpdateComment .
// @router /comment/update [POST]
func UpdateComment(ctx context.Context, c *app.RequestContext) {
	routerPath := string(c.URI().Path())

	hlog.Debugf(log.ROUTE_PATH, routerPath)
	var err error
	var req comment.UpdateCommentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		errCode.ResponseValidationFailed(c, err)
		return
	}

	resp, err := service.UpdateComment(ctx, c, &req)
	if err != nil {
		errCode.ResponseFailed(c)
		hlog.Warnf(log.REQUEST_FAILED+log.EXTRA_ERROR_INFO, routerPath, c.Errors.Last())
		return
	}

	hlog.Debugf(log.REQUEST_SUCCESSFUL, routerPath)
	c.JSON(consts.StatusOK, resp)
}

// GetComments .
// @router /comment/get [POST]
func GetComments(ctx context.Context, c *app.RequestContext) {
	routerPath := string(c.URI().Path())

	hlog.Debugf(log.ROUTE_PATH, routerPath)
	var err error
	var req comment.GetCommentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		errCode.ResponseValidationFailed(c, err)
		return
	}

	resp, err := service.GetComments(ctx, c, &req)
	if err != nil {
		errCode.ResponseFailed(c)
		hlog.Warnf(log.REQUEST_FAILED+log.EXTRA_ERROR_INFO, routerPath, c.Errors.Last())
		return
	}

	hlog.Debugf(log.REQUEST_SUCCESSFUL, routerPath)
	c.JSON(consts.StatusOK, resp)
}

// GetReplyComments .
// @router /comment/getReply [GET]
func GetReplyComments(ctx context.Context, c *app.RequestContext) {
	routerPath := string(c.URI().Path())

	hlog.Debugf(log.ROUTE_PATH, routerPath)
	var err error
	var req comment.GetReplyCommentReq
	err = c.BindAndValidate(&req)
	if err != nil {
		errCode.ResponseValidationFailed(c, err)
		return
	}

	resp, err := service.GetReplyComments(ctx, c, &req)
	if err != nil {
		errCode.ResponseFailed(c)
		hlog.Warnf(log.REQUEST_FAILED+log.EXTRA_ERROR_INFO, routerPath, c.Errors.Last())
		return
	}

	hlog.Debugf(log.REQUEST_SUCCESSFUL, routerPath)
	c.JSON(consts.StatusOK, resp)
}

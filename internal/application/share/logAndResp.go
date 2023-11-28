package share

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func WarnfAndResp(ctx *app.RequestContext, err error, bizErr *errors.Error, msg string, v ...any) *errors.Error {
	if err == nil {
		hlog.Warnf(msg, v)
		return ctx.Error(bizErr)
	}
	ctx.Error(err)
	hlog.Warnf(msg+"\nerr: %v", v, ctx.Errors.Errors())
	return ctx.Error(bizErr)
}

func WarnAndResp(ctx *app.RequestContext, err error, bizErr *errors.Error, msg string) *errors.Error {
	if err != nil {
		ctx.Error(err)
	}
	hlog.Warn(msg)
	return ctx.Error(bizErr)
}

func DebugfAndResp(ctx *app.RequestContext, err error, bizErr *errors.Error, msg string, v ...any) *errors.Error {
	if err == nil {
		hlog.Debugf(msg, v)
		return ctx.Error(bizErr)
	}
	ctx.Error(err)
	hlog.Debugf(msg+"\nerr: %v", v, ctx.Errors.Errors())
	return ctx.Error(bizErr)
}

func DebugAndResp(ctx *app.RequestContext, err error, bizErr *errors.Error, msg string) *errors.Error {
	if err != nil {
		ctx.Error(err)
	}
	hlog.Debug(msg)
	return ctx.Error(bizErr)
}

func ErrorfAndResp(ctx *app.RequestContext, err error, bizErr *errors.Error, msg string, v ...any) *errors.Error {
	if err == nil {
		hlog.Errorf(msg, v)
		return ctx.Error(bizErr)
	}
	ctx.Error(err)
	hlog.Errorf(msg+"\nerr: %v", v, ctx.Errors.Errors())
	return ctx.Error(bizErr)
}

func ErrorAndResp(ctx *app.RequestContext, err error, bizErr *errors.Error, msg string) *errors.Error {
	if err != nil {
		ctx.Error(err)
	}
	hlog.Error(msg)
	return ctx.Error(bizErr)
}

func InfofAndResp(ctx *app.RequestContext, err error, bizErr *errors.Error, msg string, v ...any) *errors.Error {
	if err == nil {
		hlog.Infof(msg, v)
		return ctx.Error(bizErr)
	}
	ctx.Error(err)
	hlog.Infof(msg+"\nerr: %v", v, ctx.Errors.Errors())
	return ctx.Error(bizErr)
}

func InfoAndResp(ctx *app.RequestContext, err error, bizErr *errors.Error, msg string) *errors.Error {
	if err != nil {
		ctx.Error(err)
	}
	hlog.Info(msg)
	return ctx.Error(bizErr)
}

package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/internal/model"
	"github.com/cold-runner/skylark/internal/model/post"
	"github.com/cold-runner/skylark/internal/pkg/errCode"
)

func (s *serviceV1) CreateDraft(c context.Context, ctx *app.RequestContext) *errors.Error {
	storedDraft := &model.Draft{}
	err := s.storeIns.CreateDraft(c, storedDraft)
	if err != nil {
		ctx.Error(err)
		return ctx.Error(errCode.ErrCreateDraft)
	}
	return nil
}

func (s *serviceV1) SaveDraft(c context.Context, ctx *app.RequestContext, userId string, draftInfo *post.DraftInfo) *errors.Error {
	panic("尚未实现")
}

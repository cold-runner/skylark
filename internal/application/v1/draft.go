package v1

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cold-runner/skylark/internal/domain/entity"
	"github.com/cold-runner/skylark/internal/domain/model_gen"
	"github.com/cold-runner/skylark/internal/infrastructure/errCode"
)

func (s *serviceV1) CreateDraft(c context.Context, ctx *app.RequestContext, draft *entity.Draft) error {
	storedDraft := &model_gen.Draft{}
	err := s.storeIns.CreateDraft(c, storedDraft)
	if err != nil {
		ctx.Error(err)
		return ctx.Error(errCode.ErrCreateDraft)
	}
	return nil
}

func (s *serviceV1) SaveDraft(c context.Context, ctx *app.RequestContext, userId string, draftInfo *entity.DraftInfo) error {
	panic("尚未实现")
}

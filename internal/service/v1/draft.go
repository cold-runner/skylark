package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/post"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/marmotedu/errors"
)

func (s *serviceV1) SaveDraft(c context.Context, userId string, draft *post.Draft) error {
	err := s.cacheIns.Set(c, userId, draft)
	if err != nil {
		return errors.WithCode(code.ErrSavePost, "", nil)
	}
	return nil
}

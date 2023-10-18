package v1

import (
	"context"
	"github.com/cold-runner/skylark/internal/model/post"
	"github.com/cold-runner/skylark/internal/pkg/code"
	"github.com/marmotedu/errors"
)

func (s *serviceV1) SaveDraft(c context.Context, userId, draftId string, draft *post.Draft) error {
	// 查缓存

	// 没缓存从数据库同步到缓存

	// 更新缓存

	// 返回前端

	// 保存到数据库
	err := s.cacheIns.SetHash(c, userId, draftId, draft)
	if err != nil {
		return errors.WithCode(code.ErrSavePost, "", nil)
	}
	return nil
}

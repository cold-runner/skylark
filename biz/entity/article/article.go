package article

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"github.com/cold-runner/skylark/biz/model/article"
)

type ArticleEntity struct {
}

func (e *ArticleEntity) QueryStore(c context.Context, ctx *app.RequestContext, storeIns store.Store) ([]*orm_gen.Categorie, *errors.Error) {
	list, err := storeIns.GetCategoryList(c)
	if err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return list, nil
}

func (e *ArticleEntity) Format(stored []*orm_gen.Categorie) ([]*article.Category, *errors.Error) {
	resp := make([]*article.Category, len(stored))
	for i := range resp {
		resp[i] = &article.Category{
			Id:          stored[i].ID.String(),
			Name:        stored[i].Name,
			CategoryUrl: stored[i].URL,
			Rank:        int32(stored[i].Rank),
			BackGround:  stored[i].BackgroundURL,
			Icon:        stored[i].Icon,
		}
	}
	return resp, nil
}

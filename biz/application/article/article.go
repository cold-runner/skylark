package article

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	entity "github.com/cold-runner/skylark/biz/entity/article"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/model/article"
)

func GetArticleFeed(c context.Context, ctx *app.RequestContext, req *article.GetArticleFeedReq) (*article.GetArticleFeedRes, error) {
	//根据热度排序从全文搜索引擎中检索出文章列表
	//zinc := searchEngine.GetZincClient()
	return nil, nil
}

func GetCategories(c context.Context, ctx *app.RequestContext, req *article.CategoriesReq) (*article.CategoriesResp, error) {
	storeIns := store.GetStoreIns()

	categoryEntity := &entity.ArticleEntity{}
	stored, err := categoryEntity.QueryStore(c, ctx, storeIns)
	if err != nil {
		return nil, err
	}

	resp, err := categoryEntity.Format(stored)
	if err != nil {
		return nil, err
	}

	return &article.CategoriesResp{
		Status:     errCode.SuccessStatus,
		Categories: resp,
	}, nil
}

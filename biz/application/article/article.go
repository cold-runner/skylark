package article

import (
	"context"

	entity "github.com/cold-runner/skylark/biz/entity/article"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/model/article"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetArticleFeed(c context.Context, ctx *app.RequestContext, req *article.GetArticleFeedReq) (*article.GetArticleFeedResp, error) {
	//根据热度排序从全文搜索引擎中检索出文章列表
	articleEntity := &entity.ArticleEntity{}
	articles, err := articleEntity.Feed(c, ctx, req)
	if err != nil {
		return nil, err
	}

	return &article.GetArticleFeedResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
		Data: &article.GetArticleFeedResp_Data{Articles: articles},
	}, nil
}

func GetCategories(c context.Context, ctx *app.RequestContext, req *article.CategoriesReq) (*article.CategoriesResp, error) {
	categoryEntity := &entity.ArticleEntity{}
	stored, err := categoryEntity.GetCategoryList(c, ctx)
	if err != nil {
		return nil, err
	}

	resp := categoryEntity.FormatCategoryList(stored)

	return &article.CategoriesResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
		Data: &article.CategoriesResp_Data{Categories: resp},
	}, nil
}

func GetArticleById(c context.Context, ctx *app.RequestContext, req *article.GetArticleReq) (*article.GetArticleResp, error) {
	articleEntity := &entity.ArticleEntity{}
	res, err := articleEntity.GetSingle(c, ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SearchArticle(c context.Context, ctx *app.RequestContext, req *article.SearchArticleReq) (*article.SearchArticleResp, error) {
	articleEntity := &entity.ArticleEntity{}
	articleList, err := articleEntity.Search(c, ctx, req)
	if err != nil {
		return nil, err
	}

	return &article.SearchArticleResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
		Data: &article.SearchArticleResp_Data{Articles: articleList},
	}, nil
}

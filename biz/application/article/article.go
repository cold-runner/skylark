package article

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	entity "github.com/cold-runner/skylark/biz/entity/article"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/model/article"
)

func GetArticleFeed(c context.Context, ctx *app.RequestContext, req *article.GetArticleFeedReq) (*article.GetArticleFeedResp, error) {
	//根据热度排序从全文搜索引擎中检索出文章列表
	articleEntity := &entity.ArticleEntity{}
	articles, err := articleEntity.Feed(c, ctx, req)
	if err != nil {
		return nil, err
	}

	return &article.GetArticleFeedResp{
		Status:   errCode.SuccessStatus,
		Articles: articles,
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
		Status:     errCode.SuccessStatus,
		Categories: resp,
	}, nil
}

func GetArticleById(c context.Context, ctx *app.RequestContext, req *article.GetArticleReq) (*article.GetArticleResp, error) {
	articleEntity := &entity.ArticleEntity{}
	res, err := articleEntity.GetSingle(c, ctx, req)
	if err != nil {
		return nil, err
	}
	return &article.GetArticleResp{
		Status:      errCode.SuccessStatus,
		Content:     res.Content,
		AuthorInfo:  res.AuthorInfo,
		Interaction: res.Interaction,
	}, nil
}

func SearchArticle(c context.Context, ctx *app.RequestContext, req *article.SearchArticleReq) (*article.SearchArticleResp, error) {
	articleEntity := &entity.ArticleEntity{}
	articleList, err := articleEntity.Search(c, ctx, req)
	if err != nil {
		return nil, err
	}

	return &article.SearchArticleResp{
		Status:   errCode.SuccessStatus,
		Articles: articleList,
	}, nil
}

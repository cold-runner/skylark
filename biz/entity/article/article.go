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

func (e *ArticleEntity) GetCategoryList(c context.Context, ctx *app.RequestContext, storeIns store.Store) ([]*orm_gen.Categorie, *errors.Error) {
	list, err := storeIns.GetCategoryList(c)
	if err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return list, nil
}

func (e *ArticleEntity) FormatCategoryList(stored []*orm_gen.Categorie) ([]*article.Category, *errors.Error) {
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

func (e *ArticleEntity) GetArticle(c context.Context, ctx *app.RequestContext,
	storeIns store.Store,
	req *article.GetArticleReq) (*orm_gen.Post, *orm_gen.UserInteraction, *orm_gen.Lark, *errors.Error) {

	//post, err := storeIns.GetPost(c, store.PostById(storeIns, req.ArticleId))
	//lark, err := storeIns.GetLark(c, store.LarkById(storeIns, post.UserID))
	//interaction, err := storeIns.GetLarkInteraction(c, mysql.LarkInteraction(storeIns, lark.ID))
	//if err != nil {
	//	return nil, nil, nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	//}
	return nil, nil, nil, nil
}

func (e *ArticleEntity) FormatArticle(stored *orm_gen.Post) (*article.ArticlesInfo, *errors.Error) {

	//resp := &article.ArticlesInfo{
	//	Basic:       errCode.SuccessStatus,
	//	AuthorInfo:  nil,
	//	Interaction: nil,
	//}

	return nil, nil
}

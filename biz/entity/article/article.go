package article

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cold-runner/skylark/biz/infrastructure/searchEngine"
	"github.com/cold-runner/skylark/biz/model/user"
	"time"

	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"github.com/cold-runner/skylark/biz/model/article"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
)

type ArticleEntity struct {
}
type searchReqBody struct {
	SearchType string   `json:"search_type"`
	Query      query    `json:"query"`
	SortFields []string `json:"sort_fields"`
	From       int32    `json:"from"`
	MaxResults int32    `json:"max_results"`
	Source     []string `json:"_source"`
}

type query struct {
	Term      string    `json:"term"`
	Field     string    `json:"field"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (e *ArticleEntity) GetCategoryList(c context.Context, ctx *app.RequestContext) ([]*orm_gen.Categorie, *errors.Error) {
	storeIns := store.GetIns()
	list, err := storeIns.GetCategoryList(c)
	if err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrUnknown)
	}
	return list, nil
}

func (e *ArticleEntity) FormatCategoryList(stored []*orm_gen.Categorie) []*article.Category {
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
	return resp
}

func (e *ArticleEntity) Search(c context.Context, ctx *app.RequestContext, req *article.SearchArticleReq) ([]*article.ArticlesInfo, *errors.Error) {
	searchIns := searchEngine.GetIns()
	// TODO 封装
	startTime, _ := time.Parse(time.RFC3339, req.StartTime)
	endTime, _ := time.Parse(time.RFC3339, req.EndTime)
	reqBody := &searchReqBody{
		SearchType: "match",
		Query: query{
			Term:      req.Keyword,
			Field:     "_all",
			StartTime: startTime,
			EndTime:   endTime,
		},
		SortFields: []string{"temperature", "-@timestamp"},
		From:       req.From,
		MaxResults: req.MaxResults,
		Source:     []string{},
	}
	res, err := searchIns.ApiQuery(c, "post", reqBody)
	if err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrSearchPosts)
	}

	return transfer(res.Hits.Hits), nil
}

func (e *ArticleEntity) GetSingle(c context.Context, ctx *app.RequestContext, req *article.GetArticleReq) (*article.GetArticleResp, *errors.Error) {
	searchIns := searchEngine.GetIns()
	storeIns := store.GetIns()
	reqBody := &searchReqBody{
		SearchType: "match",
		Query: query{
			Term:      req.ArticleId,
			Field:     "_all",
			StartTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2033, time.December, 10, 23, 0, 0, 0, time.UTC),
		},
		From:       0,
		MaxResults: 1,
		Source:     []string{},
	}
	res, err := searchIns.ApiQuery(c, "post", reqBody)
	if err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrSearchPosts)
	}
	origin := res.Hits.Hits[0].Source.(map[string]any)
	userId := origin["user_info"].(map[string]any)["user_id"].(string)
	authorInfo, err := storeIns.GetLarkAllDetail("id", userId)
	if err != nil {
		return nil, nil
	}
	u := &user.Basic{}
	data, _ := json.Marshal(authorInfo)
	_ = json.Unmarshal(data, u)

	return &article.GetArticleResp{
		Code: errCode.SuccessStatus.Code,
		Msg:  errCode.SuccessStatus.Msg,
		Data: &article.GetArticleResp_Data{
			Content:    origin["content"].(string),
			AuthorInfo: u,
		},
	}, nil
}

func (e *ArticleEntity) Feed(c context.Context, ctx *app.RequestContext, req *article.GetArticleFeedReq) ([]*article.ArticlesInfo, *errors.Error) {
	searchIns := searchEngine.GetIns()
	reqBody := &searchReqBody{
		SearchType: "alldocuments",
		From:       req.Cursor,
		MaxResults: req.Limit,
		Source:     []string{}, // TODO 考虑是否返回content
	}
	res, err := searchIns.ApiQuery(c, "post", reqBody)
	if err != nil {
		return nil, errCode.WrapBizErr(ctx, err, errCode.ErrSearchPosts)
	}
	return transfer(res.Hits.Hits), nil
}

func transfer(hits []*searchEngine.HitItem) []*article.ArticlesInfo {
	articles := make([]*article.ArticlesInfo, len(hits))
	for i := range hits {
		m := hits[i].Source.(map[string]any)
		var tags []*article.TagBasicInfo
		ts := m["tags"].([]interface{})
		for idx := range ts {
			tags = append(tags, &article.TagBasicInfo{
				Name: ts[idx].(string),
			})
		}
		userInfo := hits[i].Source.(map[string]any)["user_info"].(map[string]any)
		articles[i] = &article.ArticlesInfo{
			Basic: &article.ArticleBasicInfo{
				ArticleId:    m["id"].(string),
				Tags:         tags,
				LinkUrl:      "",
				CoverImage:   m["cover_image"].(string),
				Title:        m["title"].(string),
				BriefContent: m["summary"].(string),
				Content:      m["content"].(string),
				ViewCount:    int64(m["view_count"].(float64)),
				CollectCount: int64(m["star_count"].(float64)),
				LikeCount:    int64(m["like_count"].(float64)),
				CommentCount: int64(m["comment_count"].(float64)),
				ShareCount:   int64(m["share_count"].(float64)),
				Temperature:  int64(m["temperature"].(float64)),
			},
			UserInfo: &article.ArticlesInfo_UserInfo{
				Name:   userInfo["name"].(string),
				UserId: userInfo["user_id"].(string),
			},
		}
	}

	return articles
}

package event

import (
	"context"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/pkg/errors"

	"github.com/cold-runner/skylark/biz/infrastructure/searchEngine"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
)

func SyncToEngine(c context.Context, storeIns store.Store, searchIns searchEngine.SearchEngine) {
	_ = searchIns.DelIndex(c, "post")
	err := searchIns.CreateIndex(c, "post", &searchEngine.IndexProperty{
		"id": {
			Type:          "text",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"title": {
			Type:          "text",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"content": {
			Type:          "text", // text, keyword, date, numeric, boolean, geo_point
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"summary": {
			Type:          "text",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"picture_url": {
			Type:          "text",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"created_at": {
			Type:          "date",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"user_info": {
			Type:          "text",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"tags": {
			Type:          "keyword",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"watch": {
			Type:          "numeric",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"like": {
			Type:          "numeric",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"star": {
			Type:          "numeric",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		//"share":{},
	})
	if err != nil {
		panic(err)
	}
	data, err := storeIns.GetPostListAllDetail(c)
	if err != nil {
		panic(errors.Errorf("同步搜索引擎时查询文章列表错误！err: %v err code: %v", err, errCode.ErrSyncPosts))
	}

	dataMap := make([]map[string]any, len(data)*2)
	for i := 0; i < len(dataMap); i += 2 {
		dataMap[i], dataMap[i+1] = make(map[string]any), make(map[string]any)

		dataMap[i] = map[string]any{
			"index": map[string]string{
				"_index": "post",
			},
		}
		dataMap[i+1]["user_info"] = make(map[string]any)
		idx := i / 2
		dataMap[i+1] = map[string]any{
			"id":          data[idx].ID,
			"title":       data[idx].Title,
			"content":     data[idx].Content,
			"summary":     data[idx].Summary,
			"picture_url": data[idx].PictureURL,
			"created_at":  data[idx].CreatedAt,
			"user_info": map[string]any{
				"user_id":   data[idx].UserId,
				"stu_num":   data[idx].StuNum,
				"name":      data[idx].Name,
				"gender":    data[idx].Gender,
				"college":   data[idx].College,
				"major":     data[idx].Major,
				"grade":     data[idx].Grade,
				"province":  data[idx].Province,
				"age":       data[idx].Age,
				"photo_url": data[idx].PhotoURL,
				"introduce": data[idx].Introduce,
				"avatar":    data[idx].Avatar,
			},
			"temperature": data[idx].Temperature,
			"watch":       data[idx].Watch,
			"like":        data[idx].Like,
			"star":        data[idx].Star,
			"tag":         data[idx].Tag,
		}
	}
	err = searchIns.BulkPushDoc(c, dataMap)
	if err != nil {
		panic(errors.Errorf("同步搜索引擎时错误！err: %v err code: %v", err, errCode.ErrSyncPosts))
	}
}

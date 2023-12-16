package event

import (
	"context"
	"github.com/bytedance/gopkg/util/gopool"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cold-runner/skylark/biz/infrastructure/errCode"
	"github.com/cold-runner/skylark/biz/infrastructure/searchEngine"
	"github.com/cold-runner/skylark/biz/infrastructure/store"
	"github.com/pkg/errors"
)

func SyncToEngine(c context.Context, storeIns store.Store, searchIns searchEngine.SearchEngine) {
	preSync(c, searchIns)

	gopool.Go(func() {
		data, err := storeIns.GetPostListAllDetail(c)
		if err != nil {
			panic(errors.Errorf("同步搜索引擎时查询文章列表错误！err: %v err code: %v", err, errCode.ErrSync))
		}

		formatted := appendIndex("post", data)

		err = searchIns.BulkPushDoc(c, formatted)
		if err != nil {
			panic(errors.Errorf("同步搜索引擎时错误！err: %v err code: %v", err, errCode.ErrSync))
		}
	})
}

func preSync(c context.Context, searchIns searchEngine.SearchEngine) {
	var err error

	_ = searchIns.DelIndex(c, "post")
	err = searchIns.CreateIndex(c, "post", &searchEngine.IndexProperty{
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
		"created_at": {
			Type:          "date",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"cover_image": {
			Type:          "text",
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
		"content": {
			Type:          "text", // text, keyword, date, numeric, boolean, geo_point
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"temperature": {
			Type:          "numeric",
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
		"view_count": {
			Type:          "numeric",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"like_count": {
			Type:          "numeric",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"star_count": {
			Type:          "numeric",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"comment_count": {
			Type:          "numeric",
			Index:         true,
			Store:         true,
			Sortable:      true,
			Aggregatable:  true,
			Highlightable: true},
		"share_count": {
			Type:          "numeric",
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
	})
	if err != nil {
		panic(err)
	}

}

func appendIndex(indexName string, data []byte) []map[string]any {
	var dataMap []map[string]any
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		panic(err)
	}

	newDataMap := make([]map[string]any, len(dataMap)*2)
	for i := 0; i < len(dataMap); i++ {
		newDataMap[i*2] = map[string]any{
			"index": map[string]any{
				"_index": indexName,
			},
		}
		newDataMap[i*2+1] = dataMap[i]
	}

	return newDataMap
}

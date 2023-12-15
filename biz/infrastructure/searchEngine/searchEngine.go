package searchEngine

import (
	"context"
	"github.com/cold-runner/skylark/biz/config"
	"github.com/pkg/errors"
	"time"
)

type EngineType string

func (e EngineType) string() string {
	return string(e)
}

const (
	ZINC EngineType = "zinc"
)

var (
	e SearchEngine
)

func Init(c config.Conf) {
	se := c.GetServer().SearchEngine
	switch se {
	case ZINC.string():
		e = newZincClient(c.GetZinc())
	default:
		panic(errors.Errorf("无效的全文搜索引擎类型，支持的类型：zinc。当前传入类型为：%s", c.GetServer().SearchEngine))
	}
}

func GetIns() SearchEngine {
	if e == nil {
		panic("searchEngine is not init!")
	}
	return e
}

type SearchEngine interface {
	Health() error
	CreateIndex(c context.Context, name string, p *IndexProperty) error
	DelIndex(c context.Context, idxName string) error
	ExistIndex(c context.Context, name string) (bool, error)
	PutDoc(c context.Context, name string, id int64, doc any) error
	BulkPushDoc(c context.Context, docs []map[string]any) error
	EsQuery(c context.Context, indexName string, q any) (*QueryResultT, error)
	ApiQuery(c context.Context, indexName string, q any) (*QueryResultT, error)
	DelDoc(c context.Context, indexName, id string) error
}

type Index struct {
	Name        string         `json:"name"`
	StorageType string         `json:"storage_type"`
	Mappings    *IndexMappings `json:"mappings"`
}

type IndexMappings struct {
	Properties *IndexProperty `json:"properties"`
}

type IndexProperty map[string]*IndexPropertyT

type IndexPropertyT struct {
	Type           string `json:"type"`
	Index          bool   `json:"index"`
	Store          bool   `json:"store"`
	Sortable       bool   `json:"sortable"`
	Aggregatable   bool   `json:"aggregatable"`
	Highlightable  bool   `json:"highlightable"`
	Analyzer       string `json:"analyzer"`
	SearchAnalyzer string `json:"search_analyzer"`
	Format         string `json:"format"`
}

type QueryResultT struct {
	Took     int          `json:"took"`
	TimedOut bool         `json:"timed_out"`
	Hits     *HitsResultT `json:"hits"`
}

type HitsResultT struct {
	Total    *HitsResultTotalT `json:"total"`
	MaxScore float64           `json:"max_score"`
	Hits     []*HitItem        `json:"hits"`
}

type HitsResultTotalT struct {
	Value int64 `json:"value"`
}

type HitItem struct {
	Index     string    `json:"_index"`
	Type      string    `json:"_type"`
	ID        string    `json:"_id"`
	Score     float64   `json:"_score"`
	Timestamp time.Time `json:"@timestamp"`
	Source    any       `json:"_source"`
}

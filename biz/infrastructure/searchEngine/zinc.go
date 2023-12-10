package searchEngine

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/skylark/biz/config"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

var (
	Client *ZincClient
)

type ZincClient struct {
	Host     string
	User     string
	Password string
}

type ZincIndex struct {
	Name        string             `json:"name"`
	StorageType string             `json:"storage_type"`
	Mappings    *ZincIndexMappings `json:"mappings"`
}

type ZincIndexMappings struct {
	Properties *ZincIndexProperty `json:"properties"`
}

type ZincIndexProperty map[string]*ZincIndexPropertyT

type ZincIndexPropertyT struct {
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

// Init 初始化默认实例
func Init(cfg config.Conf) {
	opt := cfg.GetZinc()
	Client = &ZincClient{
		Host:     opt.Host,
		User:     opt.User,
		Password: opt.Password,
	}
	if err := Client.Health(); err != nil {
		panic(err)
	}
}

// GetZincClient 获取全局zinc客户端实例
func GetZincClient() *ZincClient {
	if Client == nil {
		panic("zinc还未被初始化")
	}
	return Client
}

// Health 健康检查
func (c *ZincClient) Health() error {
	resp, err := c.request().SetBody(nil).Get("/healthz")
	if err != nil || resp.StatusCode() != consts.StatusOK {
		return errors.Errorf("health ping error! err: %v", err)
	}
	return nil
}

// CreateIndex 创建索引
func (c *ZincClient) CreateIndex(name string, p *ZincIndexProperty) bool {
	data := &ZincIndex{
		Name:        name,
		StorageType: "disk",
		Mappings: &ZincIndexMappings{
			Properties: p,
		},
	}
	resp, err := c.request().SetBody(data).Put("/api/index")

	if err != nil || resp.StatusCode() != http.StatusOK {
		return false
	}

	return true
}

// ExistIndex 检查索引是否存在
func (c *ZincClient) ExistIndex(name string) bool {
	resp, err := c.request().Get("/api/index")

	if err != nil || resp.StatusCode() != http.StatusOK {
		return false
	}

	retData := &map[string]any{}
	err = json.Unmarshal([]byte(resp.String()), retData)
	if err != nil {
		return false
	}

	if _, ok := (*retData)[name]; ok {
		return true
	}

	return false
}

// PutDoc 新增/更新文档
func (c *ZincClient) PutDoc(name string, id int64, doc any) (bool, error) {
	resp, err := c.request().SetBody(doc).Put(fmt.Sprintf("/api/%s/_doc/%d", name, id))

	if err != nil {
		return false, err
	}

	if resp.StatusCode() != http.StatusOK {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

// BulkPushDoc 批量新增文档
func (c *ZincClient) BulkPushDoc(docs []map[string]any) (bool, error) {
	dataStr := ""
	for _, doc := range docs {
		str, err := json.Marshal(doc)
		if err == nil {
			dataStr = dataStr + string(str) + "\n"
		}
	}

	resp, err := c.request().SetBody(dataStr).Post("/api/_bulk")
	if err != nil {
		return false, err
	}

	if resp.StatusCode() != http.StatusOK {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func (c *ZincClient) EsQuery(indexName string, q any) (*QueryResultT, error) {
	resp, err := c.request().SetBody(q).Post(fmt.Sprintf("/es/%s/_search", indexName))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	result := &QueryResultT{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ZincClient) ApiQuery(indexName string, q any) (*QueryResultT, error) {
	resp, err := c.request().SetBody(q).Post(fmt.Sprintf("/api/%s/_search", indexName))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	result := &QueryResultT{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ZincClient) DelDoc(indexName, id string) error {
	resp, err := c.request().Delete(fmt.Sprintf("/api/%s/_doc/%s", indexName, id))
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}

	return nil
}

func (c *ZincClient) request() *resty.Request {
	client := resty.New()
	client.DisableWarn = true
	client.SetBaseURL(c.Host)
	client.SetBasicAuth(c.User, c.Password)

	return client.R()
}

package searchEngine

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/skylark/biz/config"
	"github.com/pkg/errors"
)

type zincClient struct {
	Host     string
	User     string
	Password string
	Client   *client.Client
}

func newZincClient(cfg *config.ZincClientConfig) SearchEngine {
	hzClient, _ := client.NewClient()
	zincClient := &zincClient{
		Host:     cfg.Host,
		User:     cfg.User,
		Password: cfg.Password,
		Client:   hzClient,
	}
	if err := zincClient.Health(); err != nil {
		panic(err)
	}
	return zincClient
}

// Health 健康检查
func (c *zincClient) Health() error {
	resp, err := c.do(context.Background(), consts.MethodGet, nil, "/healthz")

	if err != nil || resp.StatusCode() != consts.StatusOK {
		return errors.Errorf("health ping error! err: %v", err)
	}
	return nil
}

// CreateIndex 创建索引
func (c *zincClient) CreateIndex(ctx context.Context, name string, p *IndexProperty) bool {
	data := &Index{
		Name:        name,
		StorageType: "disk",
		Mappings: &IndexMappings{
			Properties: p,
		},
	}
	resp, err := c.do(ctx, consts.MethodPut, data, "/api/index")

	if err != nil || resp.StatusCode() != consts.StatusOK {
		return false
	}

	return true
}

// ExistIndex 检查索引是否存在
func (c *zincClient) ExistIndex(ctx context.Context, name string) bool {
	resp, err := c.do(ctx, consts.MethodGet, nil, "/api/index")

	if err != nil || resp.StatusCode() != consts.StatusOK {
		return false
	}

	retData := &map[string]any{}
	err = json.Unmarshal(resp.Body(), retData)
	if err != nil {
		return false
	}

	if _, ok := (*retData)[name]; ok {
		return true
	}

	return false
}

// PutDoc 新增/更新文档
func (c *zincClient) PutDoc(ctx context.Context, name string, id int64, doc any) (bool, error) {
	resp, err := c.do(ctx, consts.MethodPut, doc, fmt.Sprintf("/api/%s/_doc/%d", name, id))

	if err != nil {
		return false, err
	}

	if resp.StatusCode() != consts.StatusOK {
		return false, errors.New(strconv.Itoa(resp.StatusCode()))
	}

	return true, nil
}

// BulkPushDoc 批量新增文档
func (c *zincClient) BulkPushDoc(ctx context.Context, docs []map[string]any) (bool, error) {
	dataStr := ""
	for _, doc := range docs {
		str, err := json.Marshal(doc)
		if err == nil {
			dataStr = dataStr + string(str) + "\n"
		}
	}

	resp, err := c.do(ctx, consts.MethodPost, dataStr, "/api/_bulk")
	if err != nil {
		return false, err
	}

	if resp.StatusCode() != consts.StatusOK {
		return false, errors.New(strconv.Itoa(resp.StatusCode()))
	}

	return true, nil
}

func (c *zincClient) EsQuery(ctx context.Context, indexName string, q any) (*QueryResultT, error) {
	resp, err := c.do(ctx, consts.MethodPost, q, fmt.Sprintf("/es/%s/_search", indexName))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != consts.StatusOK {
		return nil, errors.New(strconv.Itoa(resp.StatusCode()))
	}

	result := &QueryResultT{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *zincClient) ApiQuery(ctx context.Context, indexName string, q any) (*QueryResultT, error) {
	resp, err := c.do(ctx, consts.MethodPost, q, fmt.Sprintf("/api/%s/_search", indexName))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != consts.StatusOK {
		return nil, errors.New(strconv.Itoa(resp.StatusCode()))
	}

	result := &QueryResultT{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *zincClient) DelDoc(ctx context.Context, indexName, id string) error {
	resp, err := c.do(ctx, consts.MethodDelete, nil, fmt.Sprintf("/api/%s/_doc/%s", indexName, id))

	if err != nil {
		return err
	}

	if resp.StatusCode() != consts.StatusOK {
		return errors.New(strconv.Itoa(resp.StatusCode()))
	}

	return nil
}

func (c *zincClient) do(ctx context.Context, method string, body any, url string) (*protocol.Response, error) {
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()
	req.SetRequestURI(c.Host + url)
	req.SetBasicAuth(c.User, c.Password)
	req.SetMethod(method)
	if body != nil {
		req.SetBody(reqBody)
	}

	if err := c.Client.Do(ctx, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

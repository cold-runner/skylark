package searchEngine

import (
	"context"
	"strconv"

	"github.com/cold-runner/skylark/biz/config"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
	c := &zincClient{
		Host:     cfg.Host,
		User:     cfg.User,
		Password: cfg.Password,
		Client:   hzClient,
	}
	if err := c.Health(); err != nil {
		panic(err)
	}
	return c
}

// Health 健康检查
func (c *zincClient) Health() error {
	resp, err := c.do(context.Background(), consts.MethodGet, nil, "/healthz")

	if err != nil {
		return err
	}
	if resp.StatusCode() != consts.StatusOK {
		return errors.Errorf("健康检查时http状态码不为200！status code: %v", resp.StatusCode())
	}
	return nil
}

// CreateIndex 创建索引
func (c *zincClient) CreateIndex(ctx context.Context, name string, p *IndexProperty) error {
	data := &Index{
		Name:        name,
		StorageType: "disk",
		Mappings: &IndexMappings{
			Properties: p,
		},
	}
	formatted, err := json.Marshal(data)
	if err != nil {
		return err
	}
	resp, err := c.do(ctx, consts.MethodPut, formatted, "/api/index")

	if err != nil {
		return err
	}
	if resp.StatusCode() != consts.StatusOK {
		return errors.Errorf("创建索引时http状态码不为200！status code: %v, resp data: %v", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

func (c *zincClient) DelIndex(ctx context.Context, idxName string) error {
	resp, err := c.do(ctx, consts.MethodDelete, nil, "/api/index/"+idxName)
	if err != nil {
		return err
	}
	if resp.StatusCode() != consts.StatusOK {
		return errors.Errorf("删除索引时http状态码不为200！status code: %v, resp data: %v", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

// ExistIndex 检查索引是否存在
func (c *zincClient) ExistIndex(ctx context.Context, name string) (bool, error) {
	resp, err := c.do(ctx, consts.MethodGet, nil, "/api/index")

	if err != nil {
		return false, err
	}
	if resp.StatusCode() != consts.StatusOK {
		return false, errors.Errorf("查找索引时http状态码不为200！status code: %v, resp data: %v", resp.StatusCode(), string(resp.Body()))
	}
	retData := &map[string]any{}
	err = json.Unmarshal(resp.Body(), retData)
	if err != nil {
		return false, err
	}

	if _, ok := (*retData)[name]; ok {
		return true, nil
	}

	return false, nil
}

// PutDoc 新增/更新文档
func (c *zincClient) PutDoc(ctx context.Context, name string, id int64, doc any) error {
	formatted, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	resp, err := c.do(ctx, consts.MethodPut, formatted, "/api/"+name+"/_doc/"+strconv.FormatInt(id, 10))

	if err != nil {
		return err
	}

	if resp.StatusCode() != consts.StatusOK {
		return errors.Errorf("新增文档时http状态码不为200！status code: %v, resp data: %v", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

// BulkPushDoc 批量新增文档
func (c *zincClient) BulkPushDoc(ctx context.Context, docs []map[string]any) error {
	var dataBytes []byte
	for _, doc := range docs {
		bytes, err := json.Marshal(doc)
		if err == nil {
			bytes = append(bytes, '\n')
			dataBytes = append(dataBytes, bytes...)
		}
	}
	resp, err := c.do(ctx, consts.MethodPost, dataBytes, "/api/_bulk")
	if err != nil {
		return err
	}

	if resp.StatusCode() != consts.StatusOK {
		return errors.Errorf("批量新增文档时http状态码不为200！status code: %v, resp data: %v", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

func (c *zincClient) EsQuery(ctx context.Context, indexName string, q any) (*QueryResultT, error) {
	formatted, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(ctx, consts.MethodPost, formatted, "/es/"+indexName+"/_search")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != consts.StatusOK {
		return nil, errors.Errorf("搜索文档时http状态码不为200！status code: %v, resp data: %v", resp.StatusCode(), string(resp.Body()))
	}

	result := &QueryResultT{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *zincClient) ApiQuery(ctx context.Context, indexName string, q any) (*QueryResultT, error) {
	formatted, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(ctx, consts.MethodPost, formatted, "/api/"+indexName+"/_search")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != consts.StatusOK {
		return nil, errors.Errorf("搜索索引时http状态码不为200！status code: %v, resp data: %v", resp.StatusCode(), string(resp.Body()))
	}

	result := &QueryResultT{}
	err = json.Unmarshal(resp.Body(), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *zincClient) DelDoc(ctx context.Context, indexName, id string) error {
	resp, err := c.do(ctx, consts.MethodDelete, nil, "/api/"+indexName+"/_doc/"+id)

	if err != nil {
		return err
	}

	if resp.StatusCode() != consts.StatusOK {
		return errors.Errorf("删除文档时http状态码不为200！status code: %v, resp data: %v", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

func (c *zincClient) do(ctx context.Context, method string, body []byte, url string) (*protocol.Response, error) {
	req, resp := protocol.AcquireRequest(), protocol.AcquireResponse()

	req.SetRequestURI(c.Host + url)
	req.SetBasicAuth(c.User, c.Password)
	req.SetMethod(method)
	req.SetBody(body)

	if err := c.Client.Do(ctx, req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

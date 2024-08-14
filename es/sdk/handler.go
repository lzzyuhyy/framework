package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

func (c *EsClient) CreateIndex(index string) error {
	create, err := c.cli.Indices.Create(index)
	if err != nil {
		return err
	}

	if create.StatusCode != 200 && create.StatusCode != 201 {
		return fmt.Errorf("error create index")
	}

	return nil
}

func (c *EsClient) IndexDoc(index string, id string, body any) (*R, error) {
	data, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	response, err := c.cli.Index(index, bytes.NewReader(data), func(request *esapi.IndexRequest) {
		request.DocumentID = id
		request.DocumentType = "_doc"
	})
	if err != nil {
		return nil, err
	}

	return handler(response)
}

func (c *EsClient) GetDocById(index string, id string) (*R, error) {
	response, err := c.cli.Get(index, id)
	if err != nil {
		return nil, err
	}

	return handler(response)
}

func (c *EsClient) SearchDoc(index string, query string) (*R, error) {
	if query == "" {
		query = `{ "query": { "match_all": {} } }`
	}
	response, err := c.cli.Search(
		c.cli.Search.WithIndex(index),
		c.cli.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}

	return handler(response)
}

func (c *EsClient) UpdateDoc(index, id string, data any) (*R, error) {
	var req struct {
		Doc interface{} `json:"doc"`
	}

	req.Doc = data
	reqStr, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	response, err := c.cli.Update(index, id, strings.NewReader(string(reqStr)))
	if err != nil {
		log.Println("update err:", err)
		return nil, err
	}

	return handler(response)
}

func (c *EsClient) DeleteDoc(index, id string) error {
	response, err := c.cli.Delete(index, id)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 && response.StatusCode != 201 {
		return fmt.Errorf("delete doc err")
	}

	return nil
}

func (c *EsClient) DeleteIndex(index ...string) error {
	response, err := c.cli.Indices.Delete(index)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 && response.StatusCode != 201 {
		return fmt.Errorf(`delete index err`)
	}

	return nil
}

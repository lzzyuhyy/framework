package esgo

import (
	"encoding/json"
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/spf13/viper"
	"strings"
)

func NewClient(addr ...string) (*es.Client, error) {
	client, err := es.NewClient(es.Config{
		//Addresses: []string{viper.GetString("es.addr")},
		Addresses: []string{"http://149.104.26.163:9200"},
	})
	if err != nil {
		return nil, err
	}

	return client, err
}

func NewClientFunc(handler func(client *es.Client) (R, error)) (R, error) {
	cli, err := es.NewClient(es.Config{
		Addresses: []string{viper.GetString("es.addr")},
	})
	if err != nil {
		return nil, err
	}

	return handler(cli)
}

type R map[string]interface{}

var data = make(R)

func CreateIndex(index string) (R, error) {
	cli, err := NewClient()
	if err != nil {
		return nil, err
	}

	response, err := cli.Indices.Create(index)
	if err != nil {
		return nil, err
	}

	return esGoHandler(response)
}

func GetInfoById(index string, id string) (R, error) {
	cli, err := NewClient()
	if err != nil {
		return nil, err
	}

	response, err := cli.Get(index, id)
	if err != nil {
		return nil, err
	}

	return esGoHandler(response)
}

func GetInfo(index string, info any) (R, error) {
	cli, err := NewClient()
	if err != nil {
		return nil, err
	}

	query, err := json.Marshal(&info)
	if err != nil {
		return nil, err
	}

	response, err := cli.Search(
		cli.Search.WithIndex(index),
		cli.Search.WithBody(strings.NewReader(string(query))),
	)

	if err != nil {
		return nil, err
	}

	return esGoHandler(response)
}

func PutInfo(index string, info any) (R, error) {
	cli, err := NewClient()
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(&info)
	if err != nil {
		return nil, err
	}

	response, err := cli.Index(index, strings.NewReader(string(marshal)))
	if err != nil {
		return nil, err
	}

	return esGoHandler(response)
}

// 改
func PostInfo(index, id string, info any) (R, error) {
	cli, err := NewClient()
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(&info)
	if err != nil {
		return nil, err
	}

	response, err := cli.Update(index, id, strings.NewReader(string(marshal)))
	if err != nil {
		return nil, err
	}

	return esGoHandler(response)
}

func DelInfo(index, id string) (R, error) {
	cli, err := NewClient()
	if err != nil {
		return nil, err
	}

	response, err := cli.Delete(index, id)
	if err != nil {
		return nil, err
	}

	return esGoHandler(response)
}

func DelIndex(index ...string) (R, error) {
	cli, err := NewClient()
	if err != nil {
		return nil, err
	}

	response, err := cli.Indices.Delete(index)
	if err != nil {
		return nil, err
	}

	return esGoHandler(response)
}

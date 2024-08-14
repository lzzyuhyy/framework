package sdk

import (
	"encoding/json"
	"fmt"
	"framework/nacos"
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type R map[string]any

type Config struct {
	Es
}

type Es struct {
	Addr string `yaml:"addr"`
}

type EsClient struct {
	cli *es.Client
}

func NewEsClient() (*EsClient, error) {
	var conf Config
	config, err := nacos.GetNacosConfig()
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(config), &conf)
	if err != nil {
		return nil, err
	}

	cli, err := es.NewClient(es.Config{
		Addresses: []string{conf.Es.Addr},
	})

	return &EsClient{
		cli,
	}, err
}

func handler(res *esapi.Response) (*R, error) {
	var r = make(R)

	if res.StatusCode != 200 && res.StatusCode != 201 {
		return nil, fmt.Errorf("error actions")
	}

	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(all, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

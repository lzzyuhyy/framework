package httpes

import (
	"encoding/json"
	"fmt"
	"github.com/lzzyuhyy/framework/es"
	"github.com/spf13/viper"
)

func GetInfo(index string, id string, data R) (R, error) {
	url := viper.GetString("es.addr") + "/"
	if index == "" {
		return nil, fmt.Errorf("params is not null")
	} else {
		url = url + index + "/_search"
	}

	if id != "" {
		url = url + index + "/_doc/" + id
	}

	marshal, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	return HttpHandler(es.GET_REQ, url, string(marshal))
}

func PutInfo(index, id string, data R) (R, error) {
	url := viper.GetString("es.addr") + "/"
	if index == "" || id == "" {
		return nil, fmt.Errorf("索引或id不能为空")
	}

	url = url + index + "/_doc/" + id

	marshal, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	return HttpHandler(es.PUT_REQ, url, string(marshal))
}

// 添加
func PostInfo(index, id string, data R) (R, error) {
	url := viper.GetString("es.addr") + "/"
	if index == "" || id == "" {
		return nil, fmt.Errorf("索引或id不能为空")
	}

	url = url + index + "/_doc/" + id

	marshal, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	return HttpHandler(es.POST_REQ, url, string(marshal))
}

func DelInfo(index, id string) (R, error) {
	url := viper.GetString("es.addr") + "/"
	if index == "" {
		return nil, fmt.Errorf("params is not null")
	}

	if id != "" {
		url = url + index + "/_doc/" + id
	}

	return HttpHandler("PUT", url, "")
}

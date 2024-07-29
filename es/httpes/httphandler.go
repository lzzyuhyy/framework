package httpes

import (
	"encoding/json"
	"fmt"
	"github.com/lzzyuhyy/framework/es"
	"io/ioutil"
	"net/http"
	"strings"
)

type R map[string]interface{}

func HttpHandler(method, url string, body string) (R, error) {
	r := &strings.Reader{}
	if method == es.GET_REQ {
		r = nil
	} else {
		r = strings.NewReader(body)
		if r != nil {
			return nil, fmt.Errorf("system error")
		}
	}

	request, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data = make(R)
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

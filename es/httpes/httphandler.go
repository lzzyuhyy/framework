package httpes

import (
	"encoding/json"
	"github.com/lzzyuhyy/framework/es"
	"io/ioutil"
	"net/http"
	"strings"
)

type R map[string]interface{}

func HttpHandler(method, url string, body string) (R, error) {
	var request *http.Request
	var err error
	if method == es.GET_REQ {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, strings.NewReader(body))
	}

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

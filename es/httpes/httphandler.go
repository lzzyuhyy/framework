package httpes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type R map[string]interface{}

func HttpHandler(method, url string, body string) (R, error) {
	r := &strings.Reader{}
	if body == "" || body == "null" {
		r = nil
	}

	r = strings.NewReader(body)
	if r == nil {
		fmt.Println(1)
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

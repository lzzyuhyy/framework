package esgo

import (
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
)

func esGoHandler(response *esapi.Response) (R, error) {
	if response.StatusCode != 200 && response.StatusCode != 201 {
		return nil, errors.New("")
	}

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

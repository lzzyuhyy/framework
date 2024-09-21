package httpresponse

import (
	"fmt"
	"io"
	"net/http"
)

func ResponseHand(r *http.Response) ([]byte, error) {
	if r.StatusCode < 200 && r.StatusCode > 300 {
		return nil, fmt.Errorf("请求失败")
	}

	return io.ReadAll(r.Body)
}

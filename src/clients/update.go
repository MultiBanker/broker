package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func OrderUpdateState(url string, order OrderUpdate) (Response, error) {
	c := http.Client{}

	b, err := json.Marshal(&order)
	if err != nil {
		return Response{}, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return Response{}, err
	}

	res, err := c.Do(req)
	if err != nil {
		return Response{}, err
	}

	if res.StatusCode != http.StatusOK {
		var errResp Response
		if err := json.NewDecoder(res.Body).Decode(&errResp); err != nil {
			return Response{}, err
		}
		return Response{}, fmt.Errorf("[ERROR] Unauthorized client %v", errResp.Message)
	}

	var resp Response

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return Response{}, err
	}
	return resp, nil
}

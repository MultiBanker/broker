package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func MoneyTransfer(url string, id string) (OrderResponse, error) {
	c := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return OrderResponse{}, err
	}

	req.Header.Add("referenceId", id)

	res, err := c.Do(req)
	if err != nil {
		return OrderResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		var errResp Response
		if err := json.NewDecoder(res.Body).Decode(&errResp); err != nil {
			return OrderResponse{}, err
		}
		return OrderResponse{}, fmt.Errorf("[ERROR] Unauthorized client %v", errResp.Message)
	}

	var resp OrderResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return OrderResponse{}, err
	}
	return resp, nil
}
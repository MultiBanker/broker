package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/MultiBanker/broker/src/servers/http/dto"
)

type Clienter interface {
	RequestOrder(ctx context.Context, order interface{}, count int, err error) ([]byte, error)
}

type Client struct {
	URL   string
	Token string
	cli   *http.Client
}

func NewClient(URL, Token string) *Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	cli := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	return &Client{
		URL:   URL,
		Token: Token,
		cli:   cli,
	}
}
func (p Client) RequestOrder(ctx context.Context, order interface{}, count int, err error) ([]byte, error) {
	if count == 0 {
		return nil, err
	}

	b, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, p.URL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+p.Token)
	res, err := p.cli.Do(req.WithContext(ctx))
	if err != nil {
		count--
		time.Sleep(2 * time.Second)
		err = fmt.Errorf("[ERROR] request order %v", err)
		return p.RequestOrder(ctx, order, count, err)
	}

	if res.StatusCode != http.StatusOK {
		var errResp dto.Response
		if err := json.NewDecoder(res.Body).Decode(&errResp); err != nil {
			return nil, err
		}
		count--
		time.Sleep(2 * time.Second)
		err = fmt.Errorf("[ERROR] request order %v", errResp.Message)
		return p.RequestOrder(ctx, order, count, err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

package httpconnector

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	client *resty.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: resty.New(),
	}
}

func (h *HttpClient) DoGet(url string, out any) error {
	resp, err := h.client.R().
		SetHeader("Accept", "application/json").
		SetResult(out).
		Get(url)

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("HTTP error: %s", resp.Status())
	}

	return nil
}

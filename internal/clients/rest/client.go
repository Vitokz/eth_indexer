package rest

import (
	"context"
	"github.com/go-resty/resty/v2"
	"time"
)

const (
	httpTimeout = 30 * time.Second
)

type Clients interface {
	GetCosmos() CosmosClient
}

type clients struct {
	cosmos CosmosClient
}

func NewRestClients(cosmos CosmosClient) Clients {
	return &clients{
		cosmos: cosmos,
	}
}

func (r *clients) GetCosmos() CosmosClient { return r.cosmos }

type client struct {
	rc *resty.Client
}

func NewClient(addr string) *client {
	cl := resty.New().
		SetBaseURL(addr).
		SetTimeout(httpTimeout)

	return &client{
		rc: cl,
		//url: addr,
	}
}

func (c *client) Get(ctx context.Context, out interface{}, endpoint string, params map[string]string) error {
	req := c.rc.R().
		SetPathParams(params).
		SetContext(ctx).
		SetResult(out).
		SetHeader("Accept", "application/json")
	//TODO: SetError

	_, err := req.Get(endpoint)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Post(ctx context.Context, out interface{}, endpoint string, body interface{}, params map[string]string) error {
	req := c.rc.R().
		SetBody(body).
		SetPathParams(params).
		SetContext(ctx).
		SetResult(out).
		SetHeader("Content-Type", "application/json")
	//TODO: SetError

	_, err := req.Get(endpoint)
	if err != nil {
		return err
	}
	return nil
}

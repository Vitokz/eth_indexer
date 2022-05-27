package jsonrpc

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ybbus/jsonrpc/v3"
	"net/http"
	"time"
)

const (
	httpTimeout = 30 * time.Second
)

type Clients interface {
}

type clients struct {
}

func NewJsonRpcClients() Clients {
	return &clients{}
}

type client struct {
	conn       jsonrpc.RPCClient
	httpClient *http.Client
}

// newClient start dial to jsonrpc server
func newClient(addr string) (client, error) {
	httpClient := &http.Client{
		Timeout: httpTimeout,
	}

	rpcClient := jsonrpc.NewClientWithOpts(addr, &jsonrpc.RPCClientOpts{
		HTTPClient: httpClient,
	})

	// check if the client is connected
	_, err := rpcClient.Call(context.Background(), "addNumbers", 1, 2)
	switch err.(type) {
	case nil:
	case *jsonrpc.HTTPError:
	default:
		// any other error
		return client{}, err
	}

	return client{
		httpClient: httpClient,
		conn:       rpcClient,
	}, nil
}

// Call is calling rpc method
func (c *client) Call(ctx context.Context, method string, params ...any) (any, error) {
	resp, err := c.conn.Call(ctx, method, params...)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, errors.Errorf("failed to send request: %s", resp.Error.Message)
	}

	return resp.Result, nil
}

// CallFor is calling rpc method with output struct
func (c *client) CallFor(ctx context.Context, out any, method string, params ...any) (err error) {
	err = c.conn.CallFor(ctx, out, method, params)
	return
}

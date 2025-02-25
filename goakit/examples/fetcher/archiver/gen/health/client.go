// Code generated by goa v3.0.6, DO NOT EDIT.
//
// health client
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/archiver/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/archiver

package health

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Client is the "health" service client.
type Client struct {
	ShowEndpoint endpoint.Endpoint
}

// NewClient initializes a "health" service client given the endpoints.
func NewClient(show endpoint.Endpoint) *Client {
	return &Client{
		ShowEndpoint: show,
	}
}

// Show calls the "show" endpoint of the "health" service.
func (c *Client) Show(ctx context.Context) (res string, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(string), nil
}

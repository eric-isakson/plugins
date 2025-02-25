// Code generated by goa v3.0.6, DO NOT EDIT.
//
// fetcher client HTTP transport
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/fetcher/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/fetcher

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
)

// Client lists the fetcher service endpoint HTTP clients.
type Client struct {
	// Fetch Doer is the HTTP client used to make requests to the fetch endpoint.
	FetchDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the fetcher service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		FetchDoer:           doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Fetch returns an endpoint that makes HTTP requests to the fetcher service
// fetch server.
func (c *Client) Fetch() endpoint.Endpoint {
	var (
		decodeResponse = DecodeFetchResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildFetchRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.FetchDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("fetcher", "fetch", err)
		}
		return decodeResponse(resp)
	}
}

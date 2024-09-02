// This file was auto-generated by Fern from our API Definition.

package client

import (
	http "net/http"

	core "github.com/Method-Security/methodazure/generated/go/core"
	option "github.com/Method-Security/methodazure/generated/go/option"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/client/client.go
package client

import (
	"net/http"
)

type Client struct {
	host       string
	apiKey     string
	httpClient *http.Client
}

func NewClient(baseUrl, apiKey string) (*Client, error) {
	return &Client{
		host:       baseUrl,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}, nil
}

func (r *Request) AddQueryParam(key, value string) *Request {
	r.queryParams[key] = value
	return r
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	for key, value := range params {
		r.queryParams[key] = value
	}
	return r
}

// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": EntrySubmitTrx Resource Client
//
// Command:
// $ go generate

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetEntrySubmitTrxPath computes a request path to the get action of EntrySubmitTrx.
func GetEntrySubmitTrxPath(trxID string) string {
	param0 := trxID

	return fmt.Sprintf("/sao/v1/entry-submit-transaction/%s", param0)
}

// Get submitted entry transaction metadata for the given ID
func (c *Client) GetEntrySubmitTrx(ctx context.Context, path string, trxD *int) (*http.Response, error) {
	req, err := c.NewGetEntrySubmitTrxRequest(ctx, path, trxD)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetEntrySubmitTrxRequest create the request corresponding to the get action endpoint of the EntrySubmitTrx resource.
func (c *Client) NewGetEntrySubmitTrxRequest(ctx context.Context, path string, trxD *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if trxD != nil {
		tmp39 := strconv.Itoa(*trxD)
		values.Set("trxD", tmp39)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowEntrySubmitTrxPath computes a request path to the show action of EntrySubmitTrx.
func ShowEntrySubmitTrxPath() string {

	return fmt.Sprintf("/sao/v1/entry-submit-transaction/")
}

// ShowEntrySubmitTrx makes a request to the show action endpoint of the EntrySubmitTrx resource
func (c *Client) ShowEntrySubmitTrx(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowEntrySubmitTrxRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowEntrySubmitTrxRequest create the request corresponding to the show action endpoint of the EntrySubmitTrx resource.
func (c *Client) NewShowEntrySubmitTrxRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
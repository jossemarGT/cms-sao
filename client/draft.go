// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": draft Resource Client
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.4.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetDraftPath computes a request path to the get action of draft.
func GetDraftPath(draftID int) string {
	param0 := strconv.Itoa(draftID)

	return fmt.Sprintf("/sao/v1/drafts/%s", param0)
}

// Get the complete Entry Draft metadata (excluding the associated sources) for the given ID
func (c *Client) GetDraft(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetDraftRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetDraftRequest create the request corresponding to the get action endpoint of the draft resource.
func (c *Client) NewGetDraftRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowDraftPath computes a request path to the show action of draft.
func ShowDraftPath() string {

	return fmt.Sprintf("/sao/v1/drafts/")
}

// List the entry drafts without their sources.
func (c *Client) ShowDraft(ctx context.Context, path string, contest *int, page *int, pageSize *int, sort *string, task *int, user *int) (*http.Response, error) {
	req, err := c.NewShowDraftRequest(ctx, path, contest, page, pageSize, sort, task, user)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowDraftRequest create the request corresponding to the show action endpoint of the draft resource.
func (c *Client) NewShowDraftRequest(ctx context.Context, path string, contest *int, page *int, pageSize *int, sort *string, task *int, user *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if contest != nil {
		tmp6 := strconv.Itoa(*contest)
		values.Set("contest", tmp6)
	}
	if page != nil {
		tmp7 := strconv.Itoa(*page)
		values.Set("page", tmp7)
	}
	if pageSize != nil {
		tmp8 := strconv.Itoa(*pageSize)
		values.Set("page_size", tmp8)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	if task != nil {
		tmp9 := strconv.Itoa(*task)
		values.Set("task", tmp9)
	}
	if user != nil {
		tmp10 := strconv.Itoa(*user)
		values.Set("user", tmp10)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

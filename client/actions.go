// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": actions Resource Client
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.4.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// SubmitEntryActionsPath computes a request path to the submitEntry action of actions.
func SubmitEntryActionsPath() string {

	return fmt.Sprintf("/sao/v1/submit-entry")
}

// Orchestrates the resource creation related to a entry submit process (Entry, Token, Result and Score).
func (c *Client) SubmitEntryActions(ctx context.Context, path string, payload *EntryPayload, contentType string) (*http.Response, error) {
	req, err := c.NewSubmitEntryActionsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSubmitEntryActionsRequest create the request corresponding to the submitEntry action endpoint of the actions resource.
func (c *Client) NewSubmitEntryActionsRequest(ctx context.Context, path string, payload *EntryPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// SubmitEntryDraftActionsPath computes a request path to the submitEntryDraft action of actions.
func SubmitEntryDraftActionsPath() string {

	return fmt.Sprintf("/sao/v1/submit-entry-draft")
}

// Orchestrates the resource creation related to a entry draft submit process (Draft and Result).
func (c *Client) SubmitEntryDraftActions(ctx context.Context, path string, payload *EntryPayload, contentType string) (*http.Response, error) {
	req, err := c.NewSubmitEntryDraftActionsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSubmitEntryDraftActionsRequest create the request corresponding to the submitEntryDraft action endpoint of the actions resource.
func (c *Client) NewSubmitEntryDraftActionsRequest(ctx context.Context, path string, payload *EntryPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// SummarizeScoreActionsPath computes a request path to the summarizeScore action of actions.
func SummarizeScoreActionsPath() string {

	return fmt.Sprintf("/sao/v1/summarize-score")
}

// List scores and its total grouped and filter by contest, task or user
func (c *Client) SummarizeScoreActions(ctx context.Context, path string, contest *int, groupBy *string, page *int, pageSize *int, sort *string, task *int, user *int) (*http.Response, error) {
	req, err := c.NewSummarizeScoreActionsRequest(ctx, path, contest, groupBy, page, pageSize, sort, task, user)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSummarizeScoreActionsRequest create the request corresponding to the summarizeScore action endpoint of the actions resource.
func (c *Client) NewSummarizeScoreActionsRequest(ctx context.Context, path string, contest *int, groupBy *string, page *int, pageSize *int, sort *string, task *int, user *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if contest != nil {
		tmp1 := strconv.Itoa(*contest)
		values.Set("contest", tmp1)
	}
	if groupBy != nil {
		values.Set("groupBy", *groupBy)
	}
	if page != nil {
		tmp2 := strconv.Itoa(*page)
		values.Set("page", tmp2)
	}
	if pageSize != nil {
		tmp3 := strconv.Itoa(*pageSize)
		values.Set("page_size", tmp3)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	if task != nil {
		tmp4 := strconv.Itoa(*task)
		values.Set("task", tmp4)
	}
	if user != nil {
		tmp5 := strconv.Itoa(*user)
		values.Set("user", tmp5)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

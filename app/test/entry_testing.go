// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": entry TestHelpers
//
// Command:
// $ go generate

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/cms-orbits/cms-sao/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// GetEntryBadRequest runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEntryBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, entryID int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/entries/%v", entryID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["entryID"] = []string{fmt.Sprintf("%v", entryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	getCtx, _err := app.NewGetEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// GetEntryNotFound runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEntryNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, entryID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/entries/%v", entryID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["entryID"] = []string{fmt.Sprintf("%v", entryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	getCtx, _err := app.NewGetEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// GetEntryOK runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEntryOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, entryID int) (http.ResponseWriter, *app.ComJossemargtSaoEntry) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/entries/%v", entryID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["entryID"] = []string{fmt.Sprintf("%v", entryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	getCtx, _err := app.NewGetEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoEntry
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoEntry)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoEntry", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetEntryOKFull runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEntryOKFull(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, entryID int) (http.ResponseWriter, *app.ComJossemargtSaoEntryFull) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/entries/%v", entryID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["entryID"] = []string{fmt.Sprintf("%v", entryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	getCtx, _err := app.NewGetEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoEntryFull
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoEntryFull)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoEntryFull", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetEntryOKLink runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetEntryOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, entryID int) (http.ResponseWriter, *app.ComJossemargtSaoEntryLink) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/entries/%v", entryID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["entryID"] = []string{fmt.Sprintf("%v", entryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	getCtx, _err := app.NewGetEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoEntryLink
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoEntryLink)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoEntryLink", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowEntryBadRequest runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowEntryBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, contest int, contestSlug string, page int, pageSize int, sort string, task int, taskSlug string, user int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		query["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		query["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		query["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		query["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/entries/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		prms["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		prms["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		prms["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		prms["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	showCtx, _err := app.NewShowEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowEntryOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowEntryOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, contest int, contestSlug string, page int, pageSize int, sort string, task int, taskSlug string, user int) (http.ResponseWriter, app.ComJossemargtSaoEntryCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		query["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		query["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		query["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		query["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/entries/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		prms["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		prms["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		prms["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		prms["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	showCtx, _err := app.NewShowEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoEntryCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoEntryCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoEntryCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowEntryOKFull runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowEntryOKFull(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, contest int, contestSlug string, page int, pageSize int, sort string, task int, taskSlug string, user int) (http.ResponseWriter, app.ComJossemargtSaoEntryFullCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		query["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		query["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		query["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		query["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/entries/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		prms["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		prms["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		prms["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		prms["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	showCtx, _err := app.NewShowEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoEntryFullCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoEntryFullCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoEntryFullCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowEntryOKLink runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowEntryOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.EntryController, contest int, contestSlug string, page int, pageSize int, sort string, task int, taskSlug string, user int) (http.ResponseWriter, app.ComJossemargtSaoEntryLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		query["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		query["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		query["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		query["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/entries/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		prms["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		prms["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		prms["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		prms["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "EntryTest"), rw, req, prms)
	showCtx, _err := app.NewShowEntryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoEntryLinkCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoEntryLinkCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoEntryLinkCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

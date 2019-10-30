// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ExactMatchDataStoreRequestBuilder is request builder for ExactMatchDataStore
type ExactMatchDataStoreRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExactMatchDataStoreRequest
func (b *ExactMatchDataStoreRequestBuilder) Request() *ExactMatchDataStoreRequest {
	return &ExactMatchDataStoreRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExactMatchDataStoreRequest is request for ExactMatchDataStore
type ExactMatchDataStoreRequest struct{ BaseRequest }

// Get performs GET request for ExactMatchDataStore
func (r *ExactMatchDataStoreRequest) Get(ctx context.Context) (resObj *ExactMatchDataStore, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExactMatchDataStore
func (r *ExactMatchDataStoreRequest) Update(ctx context.Context, reqObj *ExactMatchDataStore) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExactMatchDataStore
func (r *ExactMatchDataStoreRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Sessions returns request builder for ExactMatchSession collection
func (b *ExactMatchDataStoreRequestBuilder) Sessions() *ExactMatchDataStoreSessionsCollectionRequestBuilder {
	bb := &ExactMatchDataStoreSessionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sessions"
	return bb
}

// ExactMatchDataStoreSessionsCollectionRequestBuilder is request builder for ExactMatchSession collection
type ExactMatchDataStoreSessionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExactMatchSession collection
func (b *ExactMatchDataStoreSessionsCollectionRequestBuilder) Request() *ExactMatchDataStoreSessionsCollectionRequest {
	return &ExactMatchDataStoreSessionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ExactMatchSession item
func (b *ExactMatchDataStoreSessionsCollectionRequestBuilder) ID(id string) *ExactMatchSessionRequestBuilder {
	bb := &ExactMatchSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ExactMatchDataStoreSessionsCollectionRequest is request for ExactMatchSession collection
type ExactMatchDataStoreSessionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExactMatchSession collection
func (r *ExactMatchDataStoreSessionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ExactMatchSession, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ExactMatchSession
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ExactMatchSession
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ExactMatchSession collection
func (r *ExactMatchDataStoreSessionsCollectionRequest) Get(ctx context.Context) ([]ExactMatchSession, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ExactMatchSession collection
func (r *ExactMatchDataStoreSessionsCollectionRequest) Add(ctx context.Context, reqObj *ExactMatchSession) (resObj *ExactMatchSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

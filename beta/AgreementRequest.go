// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AgreementRequestBuilder is request builder for Agreement
type AgreementRequestBuilder struct{ BaseRequestBuilder }

// Request returns AgreementRequest
func (b *AgreementRequestBuilder) Request() *AgreementRequest {
	return &AgreementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AgreementRequest is request for Agreement
type AgreementRequest struct{ BaseRequest }

// Get performs GET request for Agreement
func (r *AgreementRequest) Get(ctx context.Context) (resObj *Agreement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Agreement
func (r *AgreementRequest) Update(ctx context.Context, reqObj *Agreement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Agreement
func (r *AgreementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Files returns request builder for AgreementFile collection
func (b *AgreementRequestBuilder) Files() *AgreementFilesCollectionRequestBuilder {
	bb := &AgreementFilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/files"
	return bb
}

// AgreementFilesCollectionRequestBuilder is request builder for AgreementFile collection
type AgreementFilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AgreementFile collection
func (b *AgreementFilesCollectionRequestBuilder) Request() *AgreementFilesCollectionRequest {
	return &AgreementFilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AgreementFile item
func (b *AgreementFilesCollectionRequestBuilder) ID(id string) *AgreementFileRequestBuilder {
	bb := &AgreementFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AgreementFilesCollectionRequest is request for AgreementFile collection
type AgreementFilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AgreementFile collection
func (r *AgreementFilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AgreementFile, error) {
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
	var values []AgreementFile
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
			value  []AgreementFile
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

// Get performs GET request for AgreementFile collection
func (r *AgreementFilesCollectionRequest) Get(ctx context.Context) ([]AgreementFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AgreementFile collection
func (r *AgreementFilesCollectionRequest) Add(ctx context.Context, reqObj *AgreementFile) (resObj *AgreementFile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

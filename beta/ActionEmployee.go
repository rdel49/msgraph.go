// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rdel49/msgraph.go/jsonx"
)

// Picture returns request builder for Picture collection
func (b *EmployeeRequestBuilder) Picture() *EmployeePictureCollectionRequestBuilder {
	bb := &EmployeePictureCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/picture"
	return bb
}

// EmployeePictureCollectionRequestBuilder is request builder for Picture collection
type EmployeePictureCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Picture collection
func (b *EmployeePictureCollectionRequestBuilder) Request() *EmployeePictureCollectionRequest {
	return &EmployeePictureCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Picture item
func (b *EmployeePictureCollectionRequestBuilder) ID(id string) *PictureRequestBuilder {
	bb := &PictureRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmployeePictureCollectionRequest is request for Picture collection
type EmployeePictureCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Picture collection
func (r *EmployeePictureCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Picture, error) {
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
	var values []Picture
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Picture
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
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

// GetN performs GET request for Picture collection, max N pages
func (r *EmployeePictureCollectionRequest) GetN(ctx context.Context, n int) ([]Picture, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Picture collection
func (r *EmployeePictureCollectionRequest) Get(ctx context.Context) ([]Picture, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Picture collection
func (r *EmployeePictureCollectionRequest) Add(ctx context.Context, reqObj *Picture) (resObj *Picture, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

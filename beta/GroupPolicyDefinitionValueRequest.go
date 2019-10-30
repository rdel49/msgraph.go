// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// GroupPolicyDefinitionValueRequestBuilder is request builder for GroupPolicyDefinitionValue
type GroupPolicyDefinitionValueRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupPolicyDefinitionValueRequest
func (b *GroupPolicyDefinitionValueRequestBuilder) Request() *GroupPolicyDefinitionValueRequest {
	return &GroupPolicyDefinitionValueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupPolicyDefinitionValueRequest is request for GroupPolicyDefinitionValue
type GroupPolicyDefinitionValueRequest struct{ BaseRequest }

// Get performs GET request for GroupPolicyDefinitionValue
func (r *GroupPolicyDefinitionValueRequest) Get(ctx context.Context) (resObj *GroupPolicyDefinitionValue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupPolicyDefinitionValue
func (r *GroupPolicyDefinitionValueRequest) Update(ctx context.Context, reqObj *GroupPolicyDefinitionValue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupPolicyDefinitionValue
func (r *GroupPolicyDefinitionValueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Definition is navigation property
func (b *GroupPolicyDefinitionValueRequestBuilder) Definition() *GroupPolicyDefinitionRequestBuilder {
	bb := &GroupPolicyDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/definition"
	return bb
}

// PresentationValues returns request builder for GroupPolicyPresentationValue collection
func (b *GroupPolicyDefinitionValueRequestBuilder) PresentationValues() *GroupPolicyDefinitionValuePresentationValuesCollectionRequestBuilder {
	bb := &GroupPolicyDefinitionValuePresentationValuesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/presentationValues"
	return bb
}

// GroupPolicyDefinitionValuePresentationValuesCollectionRequestBuilder is request builder for GroupPolicyPresentationValue collection
type GroupPolicyDefinitionValuePresentationValuesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GroupPolicyPresentationValue collection
func (b *GroupPolicyDefinitionValuePresentationValuesCollectionRequestBuilder) Request() *GroupPolicyDefinitionValuePresentationValuesCollectionRequest {
	return &GroupPolicyDefinitionValuePresentationValuesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GroupPolicyPresentationValue item
func (b *GroupPolicyDefinitionValuePresentationValuesCollectionRequestBuilder) ID(id string) *GroupPolicyPresentationValueRequestBuilder {
	bb := &GroupPolicyPresentationValueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GroupPolicyDefinitionValuePresentationValuesCollectionRequest is request for GroupPolicyPresentationValue collection
type GroupPolicyDefinitionValuePresentationValuesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GroupPolicyPresentationValue collection
func (r *GroupPolicyDefinitionValuePresentationValuesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]GroupPolicyPresentationValue, error) {
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
	var values []GroupPolicyPresentationValue
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
			value  []GroupPolicyPresentationValue
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

// Get performs GET request for GroupPolicyPresentationValue collection
func (r *GroupPolicyDefinitionValuePresentationValuesCollectionRequest) Get(ctx context.Context) ([]GroupPolicyPresentationValue, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for GroupPolicyPresentationValue collection
func (r *GroupPolicyDefinitionValuePresentationValuesCollectionRequest) Add(ctx context.Context, reqObj *GroupPolicyPresentationValue) (resObj *GroupPolicyPresentationValue, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// List undocumented
type List struct {
	// BaseItem is the base model of List
	BaseItem
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// List undocumented
	List *ListInfo `json:"list,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// System undocumented
	System *SystemFacet `json:"system,omitempty"`
	// Columns undocumented
	Columns []ColumnDefinition `json:"columns,omitempty"`
	// ContentTypes undocumented
	ContentTypes []ContentType `json:"contentTypes,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Items undocumented
	Items []ListItem `json:"items,omitempty"`
}

// ListRequestBuilder is request builder for List
type ListRequestBuilder struct{ BaseRequestBuilder }

// Request returns ListRequest
func (b *ListRequestBuilder) Request() *ListRequest {
	return &ListRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ListRequest is request for List
type ListRequest struct{ BaseRequest }

// Get performs GET request for List
func (r *ListRequest) Get(ctx context.Context) (resObj *List, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for List
func (r *ListRequest) Update(ctx context.Context, reqObj *List) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for List
func (r *ListRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Columns returns request builder for ColumnDefinition collection
func (b *ListRequestBuilder) Columns() *ListColumnsCollectionRequestBuilder {
	bb := &ListColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columns"
	return bb
}

// ListColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection
type ListColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *ListColumnsCollectionRequestBuilder) Request() *ListColumnsCollectionRequest {
	return &ListColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *ListColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListColumnsCollectionRequest is request for ColumnDefinition collection
type ListColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ColumnDefinition, error) {
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
	var values []ColumnDefinition
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
			value  []ColumnDefinition
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

// Get performs GET request for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ContentTypes returns request builder for ContentType collection
func (b *ListRequestBuilder) ContentTypes() *ListContentTypesCollectionRequestBuilder {
	bb := &ListContentTypesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contentTypes"
	return bb
}

// ListContentTypesCollectionRequestBuilder is request builder for ContentType collection
type ListContentTypesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ContentType collection
func (b *ListContentTypesCollectionRequestBuilder) Request() *ListContentTypesCollectionRequest {
	return &ListContentTypesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ContentType item
func (b *ListContentTypesCollectionRequestBuilder) ID(id string) *ContentTypeRequestBuilder {
	bb := &ContentTypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListContentTypesCollectionRequest is request for ContentType collection
type ListContentTypesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ContentType collection
func (r *ListContentTypesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ContentType, error) {
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
	var values []ContentType
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
			value  []ContentType
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

// Get performs GET request for ContentType collection
func (r *ListContentTypesCollectionRequest) Get(ctx context.Context) ([]ContentType, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ContentType collection
func (r *ListContentTypesCollectionRequest) Add(ctx context.Context, reqObj *ContentType) (resObj *ContentType, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Drive is navigation property
func (b *ListRequestBuilder) Drive() *DriveRequestBuilder {
	bb := &DriveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/drive"
	return bb
}

// Items returns request builder for ListItem collection
func (b *ListRequestBuilder) Items() *ListItemsCollectionRequestBuilder {
	bb := &ListItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// ListItemsCollectionRequestBuilder is request builder for ListItem collection
type ListItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ListItem collection
func (b *ListItemsCollectionRequestBuilder) Request() *ListItemsCollectionRequest {
	return &ListItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ListItem item
func (b *ListItemsCollectionRequestBuilder) ID(id string) *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemsCollectionRequest is request for ListItem collection
type ListItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ListItem collection
func (r *ListItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ListItem, error) {
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
	var values []ListItem
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
			value  []ListItem
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

// Get performs GET request for ListItem collection
func (r *ListItemsCollectionRequest) Get(ctx context.Context) ([]ListItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ListItem collection
func (r *ListItemsCollectionRequest) Add(ctx context.Context, reqObj *ListItem) (resObj *ListItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

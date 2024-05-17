// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rdel49/msgraph.go/jsonx"
)

// ListItemVersionRestoreVersionRequestParameter undocumented
type ListItemVersionRestoreVersionRequestParameter struct {
}

// Activities returns request builder for ItemActivityOLD collection
func (b *ListRequestBuilder) Activities() *ListActivitiesCollectionRequestBuilder {
	bb := &ListActivitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activities"
	return bb
}

// ListActivitiesCollectionRequestBuilder is request builder for ItemActivityOLD collection
type ListActivitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivityOLD collection
func (b *ListActivitiesCollectionRequestBuilder) Request() *ListActivitiesCollectionRequest {
	return &ListActivitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivityOLD item
func (b *ListActivitiesCollectionRequestBuilder) ID(id string) *ItemActivityOLDRequestBuilder {
	bb := &ItemActivityOLDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListActivitiesCollectionRequest is request for ItemActivityOLD collection
type ListActivitiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ItemActivityOLD collection
func (r *ListActivitiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ItemActivityOLD, error) {
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
	var values []ItemActivityOLD
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
			value  []ItemActivityOLD
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

// GetN performs GET request for ItemActivityOLD collection, max N pages
func (r *ListActivitiesCollectionRequest) GetN(ctx context.Context, n int) ([]ItemActivityOLD, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ItemActivityOLD collection
func (r *ListActivitiesCollectionRequest) Get(ctx context.Context) ([]ItemActivityOLD, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ItemActivityOLD collection
func (r *ListActivitiesCollectionRequest) Add(ctx context.Context, reqObj *ItemActivityOLD) (resObj *ItemActivityOLD, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
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
func (r *ListColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnDefinition, error) {
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
			value  []ColumnDefinition
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

// GetN performs GET request for ColumnDefinition collection, max N pages
func (r *ListColumnsCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	return r.GetN(ctx, 0)
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
func (r *ListContentTypesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ContentType, error) {
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
			value  []ContentType
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

// GetN performs GET request for ContentType collection, max N pages
func (r *ListContentTypesCollectionRequest) GetN(ctx context.Context, n int) ([]ContentType, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ContentType collection
func (r *ListContentTypesCollectionRequest) Get(ctx context.Context) ([]ContentType, error) {
	return r.GetN(ctx, 0)
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
func (r *ListItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ListItem, error) {
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
			value  []ListItem
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

// GetN performs GET request for ListItem collection, max N pages
func (r *ListItemsCollectionRequest) GetN(ctx context.Context, n int) ([]ListItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ListItem collection
func (r *ListItemsCollectionRequest) Get(ctx context.Context) ([]ListItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ListItem collection
func (r *ListItemsCollectionRequest) Add(ctx context.Context, reqObj *ListItem) (resObj *ListItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Subscriptions returns request builder for Subscription collection
func (b *ListRequestBuilder) Subscriptions() *ListSubscriptionsCollectionRequestBuilder {
	bb := &ListSubscriptionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/subscriptions"
	return bb
}

// ListSubscriptionsCollectionRequestBuilder is request builder for Subscription collection
type ListSubscriptionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Subscription collection
func (b *ListSubscriptionsCollectionRequestBuilder) Request() *ListSubscriptionsCollectionRequest {
	return &ListSubscriptionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Subscription item
func (b *ListSubscriptionsCollectionRequestBuilder) ID(id string) *SubscriptionRequestBuilder {
	bb := &SubscriptionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListSubscriptionsCollectionRequest is request for Subscription collection
type ListSubscriptionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Subscription, error) {
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
	var values []Subscription
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
			value  []Subscription
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

// GetN performs GET request for Subscription collection, max N pages
func (r *ListSubscriptionsCollectionRequest) GetN(ctx context.Context, n int) ([]Subscription, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Get(ctx context.Context) ([]Subscription, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Add(ctx context.Context, reqObj *Subscription) (resObj *Subscription, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Activities returns request builder for ItemActivityOLD collection
func (b *ListItemRequestBuilder) Activities() *ListItemActivitiesCollectionRequestBuilder {
	bb := &ListItemActivitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activities"
	return bb
}

// ListItemActivitiesCollectionRequestBuilder is request builder for ItemActivityOLD collection
type ListItemActivitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivityOLD collection
func (b *ListItemActivitiesCollectionRequestBuilder) Request() *ListItemActivitiesCollectionRequest {
	return &ListItemActivitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivityOLD item
func (b *ListItemActivitiesCollectionRequestBuilder) ID(id string) *ItemActivityOLDRequestBuilder {
	bb := &ItemActivityOLDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemActivitiesCollectionRequest is request for ItemActivityOLD collection
type ListItemActivitiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ItemActivityOLD collection
func (r *ListItemActivitiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ItemActivityOLD, error) {
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
	var values []ItemActivityOLD
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
			value  []ItemActivityOLD
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

// GetN performs GET request for ItemActivityOLD collection, max N pages
func (r *ListItemActivitiesCollectionRequest) GetN(ctx context.Context, n int) ([]ItemActivityOLD, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ItemActivityOLD collection
func (r *ListItemActivitiesCollectionRequest) Get(ctx context.Context) ([]ItemActivityOLD, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ItemActivityOLD collection
func (r *ListItemActivitiesCollectionRequest) Add(ctx context.Context, reqObj *ItemActivityOLD) (resObj *ItemActivityOLD, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Analytics is navigation property
func (b *ListItemRequestBuilder) Analytics() *ItemAnalyticsRequestBuilder {
	bb := &ItemAnalyticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/analytics"
	return bb
}

// DriveItem is navigation property
func (b *ListItemRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}

// Fields is navigation property
func (b *ListItemRequestBuilder) Fields() *FieldValueSetRequestBuilder {
	bb := &FieldValueSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fields"
	return bb
}

// Versions returns request builder for ListItemVersion collection
func (b *ListItemRequestBuilder) Versions() *ListItemVersionsCollectionRequestBuilder {
	bb := &ListItemVersionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/versions"
	return bb
}

// ListItemVersionsCollectionRequestBuilder is request builder for ListItemVersion collection
type ListItemVersionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ListItemVersion collection
func (b *ListItemVersionsCollectionRequestBuilder) Request() *ListItemVersionsCollectionRequest {
	return &ListItemVersionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ListItemVersion item
func (b *ListItemVersionsCollectionRequestBuilder) ID(id string) *ListItemVersionRequestBuilder {
	bb := &ListItemVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemVersionsCollectionRequest is request for ListItemVersion collection
type ListItemVersionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ListItemVersion, error) {
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
	var values []ListItemVersion
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
			value  []ListItemVersion
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

// GetN performs GET request for ListItemVersion collection, max N pages
func (r *ListItemVersionsCollectionRequest) GetN(ctx context.Context, n int) ([]ListItemVersion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Get(ctx context.Context) ([]ListItemVersion, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Add(ctx context.Context, reqObj *ListItemVersion) (resObj *ListItemVersion, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Fields is navigation property
func (b *ListItemVersionRequestBuilder) Fields() *FieldValueSetRequestBuilder {
	bb := &FieldValueSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fields"
	return bb
}

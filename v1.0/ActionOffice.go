// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rdel49/msgraph.go/jsonx"
)

// Shared returns request builder for SharedInsight collection
func (b *OfficeGraphInsightsRequestBuilder) Shared() *OfficeGraphInsightsSharedCollectionRequestBuilder {
	bb := &OfficeGraphInsightsSharedCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shared"
	return bb
}

// OfficeGraphInsightsSharedCollectionRequestBuilder is request builder for SharedInsight collection
type OfficeGraphInsightsSharedCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SharedInsight collection
func (b *OfficeGraphInsightsSharedCollectionRequestBuilder) Request() *OfficeGraphInsightsSharedCollectionRequest {
	return &OfficeGraphInsightsSharedCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SharedInsight item
func (b *OfficeGraphInsightsSharedCollectionRequestBuilder) ID(id string) *SharedInsightRequestBuilder {
	bb := &SharedInsightRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsSharedCollectionRequest is request for SharedInsight collection
type OfficeGraphInsightsSharedCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SharedInsight, error) {
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
	var values []SharedInsight
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
			value  []SharedInsight
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

// GetN performs GET request for SharedInsight collection, max N pages
func (r *OfficeGraphInsightsSharedCollectionRequest) GetN(ctx context.Context, n int) ([]SharedInsight, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Get(ctx context.Context) ([]SharedInsight, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Add(ctx context.Context, reqObj *SharedInsight) (resObj *SharedInsight, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Trending returns request builder for Trending collection
func (b *OfficeGraphInsightsRequestBuilder) Trending() *OfficeGraphInsightsTrendingCollectionRequestBuilder {
	bb := &OfficeGraphInsightsTrendingCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/trending"
	return bb
}

// OfficeGraphInsightsTrendingCollectionRequestBuilder is request builder for Trending collection
type OfficeGraphInsightsTrendingCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Trending collection
func (b *OfficeGraphInsightsTrendingCollectionRequestBuilder) Request() *OfficeGraphInsightsTrendingCollectionRequest {
	return &OfficeGraphInsightsTrendingCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Trending item
func (b *OfficeGraphInsightsTrendingCollectionRequestBuilder) ID(id string) *TrendingRequestBuilder {
	bb := &TrendingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsTrendingCollectionRequest is request for Trending collection
type OfficeGraphInsightsTrendingCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Trending, error) {
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
	var values []Trending
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
			value  []Trending
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

// GetN performs GET request for Trending collection, max N pages
func (r *OfficeGraphInsightsTrendingCollectionRequest) GetN(ctx context.Context, n int) ([]Trending, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Get(ctx context.Context) ([]Trending, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Add(ctx context.Context, reqObj *Trending) (resObj *Trending, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Used returns request builder for UsedInsight collection
func (b *OfficeGraphInsightsRequestBuilder) Used() *OfficeGraphInsightsUsedCollectionRequestBuilder {
	bb := &OfficeGraphInsightsUsedCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/used"
	return bb
}

// OfficeGraphInsightsUsedCollectionRequestBuilder is request builder for UsedInsight collection
type OfficeGraphInsightsUsedCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UsedInsight collection
func (b *OfficeGraphInsightsUsedCollectionRequestBuilder) Request() *OfficeGraphInsightsUsedCollectionRequest {
	return &OfficeGraphInsightsUsedCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UsedInsight item
func (b *OfficeGraphInsightsUsedCollectionRequestBuilder) ID(id string) *UsedInsightRequestBuilder {
	bb := &UsedInsightRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsUsedCollectionRequest is request for UsedInsight collection
type OfficeGraphInsightsUsedCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UsedInsight, error) {
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
	var values []UsedInsight
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
			value  []UsedInsight
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

// GetN performs GET request for UsedInsight collection, max N pages
func (r *OfficeGraphInsightsUsedCollectionRequest) GetN(ctx context.Context, n int) ([]UsedInsight, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Get(ctx context.Context) ([]UsedInsight, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Add(ctx context.Context, reqObj *UsedInsight) (resObj *UsedInsight, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PlannerPlan undocumented
type PlannerPlan struct {
	// PlannerDelta is the base model of PlannerPlan
	PlannerDelta
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Owner undocumented
	Owner *string `json:"owner,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// Contexts undocumented
	Contexts *PlannerPlanContextCollection `json:"contexts,omitempty"`
	// Tasks undocumented
	Tasks []PlannerTask `json:"tasks,omitempty"`
	// Buckets undocumented
	Buckets []PlannerBucket `json:"buckets,omitempty"`
	// Details undocumented
	Details *PlannerPlanDetails `json:"details,omitempty"`
}

// PlannerPlanRequestBuilder is request builder for PlannerPlan
type PlannerPlanRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerPlanRequest
func (b *PlannerPlanRequestBuilder) Request() *PlannerPlanRequest {
	return &PlannerPlanRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerPlanRequest is request for PlannerPlan
type PlannerPlanRequest struct{ BaseRequest }

// Get performs GET request for PlannerPlan
func (r *PlannerPlanRequest) Get(ctx context.Context) (resObj *PlannerPlan, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PlannerPlan
func (r *PlannerPlanRequest) Update(ctx context.Context, reqObj *PlannerPlan) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PlannerPlan
func (r *PlannerPlanRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Buckets returns request builder for PlannerBucket collection
func (b *PlannerPlanRequestBuilder) Buckets() *PlannerPlanBucketsCollectionRequestBuilder {
	bb := &PlannerPlanBucketsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/buckets"
	return bb
}

// PlannerPlanBucketsCollectionRequestBuilder is request builder for PlannerBucket collection
type PlannerPlanBucketsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerBucket collection
func (b *PlannerPlanBucketsCollectionRequestBuilder) Request() *PlannerPlanBucketsCollectionRequest {
	return &PlannerPlanBucketsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerBucket item
func (b *PlannerPlanBucketsCollectionRequestBuilder) ID(id string) *PlannerBucketRequestBuilder {
	bb := &PlannerBucketRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerPlanBucketsCollectionRequest is request for PlannerBucket collection
type PlannerPlanBucketsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]PlannerBucket, error) {
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
	var values []PlannerBucket
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
			value  []PlannerBucket
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

// Get performs GET request for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Get(ctx context.Context) ([]PlannerBucket, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Add(ctx context.Context, reqObj *PlannerBucket) (resObj *PlannerBucket, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Details is navigation property
func (b *PlannerPlanRequestBuilder) Details() *PlannerPlanDetailsRequestBuilder {
	bb := &PlannerPlanDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/details"
	return bb
}

// Tasks returns request builder for PlannerTask collection
func (b *PlannerPlanRequestBuilder) Tasks() *PlannerPlanTasksCollectionRequestBuilder {
	bb := &PlannerPlanTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// PlannerPlanTasksCollectionRequestBuilder is request builder for PlannerTask collection
type PlannerPlanTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerTask collection
func (b *PlannerPlanTasksCollectionRequestBuilder) Request() *PlannerPlanTasksCollectionRequest {
	return &PlannerPlanTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerTask item
func (b *PlannerPlanTasksCollectionRequestBuilder) ID(id string) *PlannerTaskRequestBuilder {
	bb := &PlannerTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerPlanTasksCollectionRequest is request for PlannerTask collection
type PlannerPlanTasksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]PlannerTask, error) {
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
	var values []PlannerTask
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
			value  []PlannerTask
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

// Get performs GET request for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Get(ctx context.Context) ([]PlannerTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Add(ctx context.Context, reqObj *PlannerTask) (resObj *PlannerTask, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

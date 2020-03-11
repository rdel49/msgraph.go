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

// WindowsDefenderApplicationControlSupplementalPolicy undocumented
type WindowsDefenderApplicationControlSupplementalPolicy struct {
	// Entity is the base model of WindowsDefenderApplicationControlSupplementalPolicy
	Entity
	// DisplayName The display name of WindowsDefenderApplicationControl supplemental policy.
	DisplayName *string `json:"displayName,omitempty"`
	// Description The description of WindowsDefenderApplicationControl supplemental policy.
	Description *string `json:"description,omitempty"`
	// Content The WindowsDefenderApplicationControl supplemental policy content in byte array format.
	Content *Binary `json:"content,omitempty"`
	// ContentFileName The WindowsDefenderApplicationControl supplemental policy content's file name.
	ContentFileName *string `json:"contentFileName,omitempty"`
	// Version The WindowsDefenderApplicationControl supplemental policy's version.
	Version *string `json:"version,omitempty"`
	// CreationDateTime The date and time when the WindowsDefenderApplicationControl supplemental policy was uploaded.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
	// LastModifiedDateTime The date and time when the WindowsDefenderApplicationControl supplemental policy was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// RoleScopeTagIDs List of Scope Tags for this WindowsDefenderApplicationControl supplemental policy entity.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// Assignments undocumented
	Assignments []WindowsDefenderApplicationControlSupplementalPolicyAssignment `json:"assignments,omitempty"`
	// DeploySummary undocumented
	DeploySummary *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary `json:"deploySummary,omitempty"`
	// DeviceStatuses undocumented
	DeviceStatuses []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus `json:"deviceStatuses,omitempty"`
}

// WindowsDefenderApplicationControlSupplementalPolicyAssignRequestParameter undocumented
type WindowsDefenderApplicationControlSupplementalPolicyAssignRequestParameter struct {
	// WdacPolicyAssignments undocumented
	WdacPolicyAssignments []WindowsDefenderApplicationControlSupplementalPolicyAssignment `json:"wdacPolicyAssignments,omitempty"`
}

// WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder is request builder for WindowsDefenderApplicationControlSupplementalPolicy
type WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsDefenderApplicationControlSupplementalPolicyRequest
func (b *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Request() *WindowsDefenderApplicationControlSupplementalPolicyRequest {
	return &WindowsDefenderApplicationControlSupplementalPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsDefenderApplicationControlSupplementalPolicyRequest is request for WindowsDefenderApplicationControlSupplementalPolicy
type WindowsDefenderApplicationControlSupplementalPolicyRequest struct{ BaseRequest }

// Get performs GET request for WindowsDefenderApplicationControlSupplementalPolicy
func (r *WindowsDefenderApplicationControlSupplementalPolicyRequest) Get(ctx context.Context) (resObj *WindowsDefenderApplicationControlSupplementalPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsDefenderApplicationControlSupplementalPolicy
func (r *WindowsDefenderApplicationControlSupplementalPolicyRequest) Update(ctx context.Context, reqObj *WindowsDefenderApplicationControlSupplementalPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsDefenderApplicationControlSupplementalPolicy
func (r *WindowsDefenderApplicationControlSupplementalPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for WindowsDefenderApplicationControlSupplementalPolicyAssignment collection
func (b *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Assignments() *WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequestBuilder {
	bb := &WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequestBuilder is request builder for WindowsDefenderApplicationControlSupplementalPolicyAssignment collection
type WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsDefenderApplicationControlSupplementalPolicyAssignment collection
func (b *WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequestBuilder) Request() *WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequest {
	return &WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsDefenderApplicationControlSupplementalPolicyAssignment item
func (b *WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequestBuilder) ID(id string) *WindowsDefenderApplicationControlSupplementalPolicyAssignmentRequestBuilder {
	bb := &WindowsDefenderApplicationControlSupplementalPolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequest is request for WindowsDefenderApplicationControlSupplementalPolicyAssignment collection
type WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsDefenderApplicationControlSupplementalPolicyAssignment collection
func (r *WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WindowsDefenderApplicationControlSupplementalPolicyAssignment, error) {
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
	var values []WindowsDefenderApplicationControlSupplementalPolicyAssignment
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
			value  []WindowsDefenderApplicationControlSupplementalPolicyAssignment
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

// Get performs GET request for WindowsDefenderApplicationControlSupplementalPolicyAssignment collection
func (r *WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequest) Get(ctx context.Context) ([]WindowsDefenderApplicationControlSupplementalPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WindowsDefenderApplicationControlSupplementalPolicyAssignment collection
func (r *WindowsDefenderApplicationControlSupplementalPolicyAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *WindowsDefenderApplicationControlSupplementalPolicyAssignment) (resObj *WindowsDefenderApplicationControlSupplementalPolicyAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeploySummary is navigation property
func (b *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) DeploySummary() *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequestBuilder {
	bb := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deploySummary"
	return bb
}

// DeviceStatuses returns request builder for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus collection
func (b *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) DeviceStatuses() *WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequestBuilder {
	bb := &WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatuses"
	return bb
}

// WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequestBuilder is request builder for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus collection
type WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus collection
func (b *WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequestBuilder) Request() *WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequest {
	return &WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus item
func (b *WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequestBuilder) ID(id string) *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder {
	bb := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequest is request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus collection
type WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus collection
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus, error) {
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
	var values []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus
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
			value  []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus
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

// Get performs GET request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus collection
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequest) Get(ctx context.Context) ([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus collection
func (r *WindowsDefenderApplicationControlSupplementalPolicyDeviceStatusesCollectionRequest) Add(ctx context.Context, reqObj *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) (resObj *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

//
type WindowsDefenderApplicationControlSupplementalPolicyAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Assign(reqObj *WindowsDefenderApplicationControlSupplementalPolicyAssignRequestParameter) *WindowsDefenderApplicationControlSupplementalPolicyAssignRequestBuilder {
	bb := &WindowsDefenderApplicationControlSupplementalPolicyAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsDefenderApplicationControlSupplementalPolicyAssignRequest struct{ BaseRequest }

//
func (b *WindowsDefenderApplicationControlSupplementalPolicyAssignRequestBuilder) Request() *WindowsDefenderApplicationControlSupplementalPolicyAssignRequest {
	return &WindowsDefenderApplicationControlSupplementalPolicyAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsDefenderApplicationControlSupplementalPolicyAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

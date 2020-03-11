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

// DeviceConfiguration Device Configuration.
type DeviceConfiguration struct {
	// Entity is the base model of DeviceConfiguration
	Entity
	// LastModifiedDateTime DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// CreatedDateTime DateTime the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description Admin provided description of the Device Configuration.
	Description *string `json:"description,omitempty"`
	// DisplayName Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`
	// Version Version of the device configuration.
	Version *int `json:"version,omitempty"`
	// Assignments undocumented
	Assignments []DeviceConfigurationAssignment `json:"assignments,omitempty"`
	// DeviceStatuses undocumented
	DeviceStatuses []DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`
	// UserStatuses undocumented
	UserStatuses []DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`
	// DeviceStatusOverview undocumented
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`
	// UserStatusOverview undocumented
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`
	// DeviceSettingStateSummaries undocumented
	DeviceSettingStateSummaries []SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`
}

// DeviceConfigurationAssignRequestParameter undocumented
type DeviceConfigurationAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []DeviceConfigurationAssignment `json:"assignments,omitempty"`
}

// DeviceConfigurationRequestBuilder is request builder for DeviceConfiguration
type DeviceConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationRequest
func (b *DeviceConfigurationRequestBuilder) Request() *DeviceConfigurationRequest {
	return &DeviceConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceConfigurationRequest is request for DeviceConfiguration
type DeviceConfigurationRequest struct{ BaseRequest }

// Get performs GET request for DeviceConfiguration
func (r *DeviceConfigurationRequest) Get(ctx context.Context) (resObj *DeviceConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceConfiguration
func (r *DeviceConfigurationRequest) Update(ctx context.Context, reqObj *DeviceConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceConfiguration
func (r *DeviceConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for DeviceConfigurationAssignment collection
func (b *DeviceConfigurationRequestBuilder) Assignments() *DeviceConfigurationAssignmentsCollectionRequestBuilder {
	bb := &DeviceConfigurationAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// DeviceConfigurationAssignmentsCollectionRequestBuilder is request builder for DeviceConfigurationAssignment collection
type DeviceConfigurationAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceConfigurationAssignment collection
func (b *DeviceConfigurationAssignmentsCollectionRequestBuilder) Request() *DeviceConfigurationAssignmentsCollectionRequest {
	return &DeviceConfigurationAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceConfigurationAssignment item
func (b *DeviceConfigurationAssignmentsCollectionRequestBuilder) ID(id string) *DeviceConfigurationAssignmentRequestBuilder {
	bb := &DeviceConfigurationAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceConfigurationAssignmentsCollectionRequest is request for DeviceConfigurationAssignment collection
type DeviceConfigurationAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceConfigurationAssignment collection
func (r *DeviceConfigurationAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceConfigurationAssignment, error) {
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
	var values []DeviceConfigurationAssignment
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
			value  []DeviceConfigurationAssignment
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

// Get performs GET request for DeviceConfigurationAssignment collection
func (r *DeviceConfigurationAssignmentsCollectionRequest) Get(ctx context.Context) ([]DeviceConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceConfigurationAssignment collection
func (r *DeviceConfigurationAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *DeviceConfigurationAssignment) (resObj *DeviceConfigurationAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceSettingStateSummaries returns request builder for SettingStateDeviceSummary collection
func (b *DeviceConfigurationRequestBuilder) DeviceSettingStateSummaries() *DeviceConfigurationDeviceSettingStateSummariesCollectionRequestBuilder {
	bb := &DeviceConfigurationDeviceSettingStateSummariesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceSettingStateSummaries"
	return bb
}

// DeviceConfigurationDeviceSettingStateSummariesCollectionRequestBuilder is request builder for SettingStateDeviceSummary collection
type DeviceConfigurationDeviceSettingStateSummariesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SettingStateDeviceSummary collection
func (b *DeviceConfigurationDeviceSettingStateSummariesCollectionRequestBuilder) Request() *DeviceConfigurationDeviceSettingStateSummariesCollectionRequest {
	return &DeviceConfigurationDeviceSettingStateSummariesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SettingStateDeviceSummary item
func (b *DeviceConfigurationDeviceSettingStateSummariesCollectionRequestBuilder) ID(id string) *SettingStateDeviceSummaryRequestBuilder {
	bb := &SettingStateDeviceSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceConfigurationDeviceSettingStateSummariesCollectionRequest is request for SettingStateDeviceSummary collection
type DeviceConfigurationDeviceSettingStateSummariesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SettingStateDeviceSummary collection
func (r *DeviceConfigurationDeviceSettingStateSummariesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SettingStateDeviceSummary, error) {
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
	var values []SettingStateDeviceSummary
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
			value  []SettingStateDeviceSummary
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

// Get performs GET request for SettingStateDeviceSummary collection
func (r *DeviceConfigurationDeviceSettingStateSummariesCollectionRequest) Get(ctx context.Context) ([]SettingStateDeviceSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SettingStateDeviceSummary collection
func (r *DeviceConfigurationDeviceSettingStateSummariesCollectionRequest) Add(ctx context.Context, reqObj *SettingStateDeviceSummary) (resObj *SettingStateDeviceSummary, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceStatusOverview is navigation property
func (b *DeviceConfigurationRequestBuilder) DeviceStatusOverview() *DeviceConfigurationDeviceOverviewRequestBuilder {
	bb := &DeviceConfigurationDeviceOverviewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatusOverview"
	return bb
}

// DeviceStatuses returns request builder for DeviceConfigurationDeviceStatus collection
func (b *DeviceConfigurationRequestBuilder) DeviceStatuses() *DeviceConfigurationDeviceStatusesCollectionRequestBuilder {
	bb := &DeviceConfigurationDeviceStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStatuses"
	return bb
}

// DeviceConfigurationDeviceStatusesCollectionRequestBuilder is request builder for DeviceConfigurationDeviceStatus collection
type DeviceConfigurationDeviceStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceConfigurationDeviceStatus collection
func (b *DeviceConfigurationDeviceStatusesCollectionRequestBuilder) Request() *DeviceConfigurationDeviceStatusesCollectionRequest {
	return &DeviceConfigurationDeviceStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceConfigurationDeviceStatus item
func (b *DeviceConfigurationDeviceStatusesCollectionRequestBuilder) ID(id string) *DeviceConfigurationDeviceStatusRequestBuilder {
	bb := &DeviceConfigurationDeviceStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceConfigurationDeviceStatusesCollectionRequest is request for DeviceConfigurationDeviceStatus collection
type DeviceConfigurationDeviceStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceConfigurationDeviceStatus collection
func (r *DeviceConfigurationDeviceStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceConfigurationDeviceStatus, error) {
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
	var values []DeviceConfigurationDeviceStatus
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
			value  []DeviceConfigurationDeviceStatus
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

// Get performs GET request for DeviceConfigurationDeviceStatus collection
func (r *DeviceConfigurationDeviceStatusesCollectionRequest) Get(ctx context.Context) ([]DeviceConfigurationDeviceStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceConfigurationDeviceStatus collection
func (r *DeviceConfigurationDeviceStatusesCollectionRequest) Add(ctx context.Context, reqObj *DeviceConfigurationDeviceStatus) (resObj *DeviceConfigurationDeviceStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserStatusOverview is navigation property
func (b *DeviceConfigurationRequestBuilder) UserStatusOverview() *DeviceConfigurationUserOverviewRequestBuilder {
	bb := &DeviceConfigurationUserOverviewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userStatusOverview"
	return bb
}

// UserStatuses returns request builder for DeviceConfigurationUserStatus collection
func (b *DeviceConfigurationRequestBuilder) UserStatuses() *DeviceConfigurationUserStatusesCollectionRequestBuilder {
	bb := &DeviceConfigurationUserStatusesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userStatuses"
	return bb
}

// DeviceConfigurationUserStatusesCollectionRequestBuilder is request builder for DeviceConfigurationUserStatus collection
type DeviceConfigurationUserStatusesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceConfigurationUserStatus collection
func (b *DeviceConfigurationUserStatusesCollectionRequestBuilder) Request() *DeviceConfigurationUserStatusesCollectionRequest {
	return &DeviceConfigurationUserStatusesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceConfigurationUserStatus item
func (b *DeviceConfigurationUserStatusesCollectionRequestBuilder) ID(id string) *DeviceConfigurationUserStatusRequestBuilder {
	bb := &DeviceConfigurationUserStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceConfigurationUserStatusesCollectionRequest is request for DeviceConfigurationUserStatus collection
type DeviceConfigurationUserStatusesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceConfigurationUserStatus collection
func (r *DeviceConfigurationUserStatusesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceConfigurationUserStatus, error) {
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
	var values []DeviceConfigurationUserStatus
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
			value  []DeviceConfigurationUserStatus
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

// Get performs GET request for DeviceConfigurationUserStatus collection
func (r *DeviceConfigurationUserStatusesCollectionRequest) Get(ctx context.Context) ([]DeviceConfigurationUserStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceConfigurationUserStatus collection
func (r *DeviceConfigurationUserStatusesCollectionRequest) Add(ctx context.Context, reqObj *DeviceConfigurationUserStatus) (resObj *DeviceConfigurationUserStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

//
type DeviceConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *DeviceConfigurationRequestBuilder) Assign(reqObj *DeviceConfigurationAssignRequestParameter) *DeviceConfigurationAssignRequestBuilder {
	bb := &DeviceConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceConfigurationAssignRequest struct{ BaseRequest }

//
func (b *DeviceConfigurationAssignRequestBuilder) Request() *DeviceConfigurationAssignRequest {
	return &DeviceConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceConfigurationAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceConfigurationAssignment, error) {
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
	var values []DeviceConfigurationAssignment
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
			value  []DeviceConfigurationAssignment
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

//
func (r *DeviceConfigurationAssignRequest) Post(ctx context.Context) ([]DeviceConfigurationAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}

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

// EmbeddedSIMActivationCodePool A pool represents a group of embedded SIM activation codes.
type EmbeddedSIMActivationCodePool struct {
	// Entity is the base model of EmbeddedSIMActivationCodePool
	Entity
	// DisplayName The admin defined name of the embedded SIM activation code pool.
	DisplayName *string `json:"displayName,omitempty"`
	// CreatedDateTime The time the embedded SIM activation code pool was created. Generated service side.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedDateTime The time the embedded SIM activation code pool was last modified. Updated service side.
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// ActivationCodes The activation codes which belong to this pool. This navigation property is used to post activation codes to Intune but cannot be used to read activation codes from Intune.
	ActivationCodes []EmbeddedSIMActivationCode `json:"activationCodes,omitempty"`
	// ActivationCodeCount The total count of activation codes which belong to this pool.
	ActivationCodeCount *int `json:"activationCodeCount,omitempty"`
	// Assignments undocumented
	Assignments []EmbeddedSIMActivationCodePoolAssignment `json:"assignments,omitempty"`
	// DeviceStates undocumented
	DeviceStates []EmbeddedSIMDeviceState `json:"deviceStates,omitempty"`
}

// EmbeddedSIMActivationCodePoolAssignRequestParameter undocumented
type EmbeddedSIMActivationCodePoolAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []EmbeddedSIMActivationCodePoolAssignment `json:"assignments,omitempty"`
}

// EmbeddedSIMActivationCodePoolRequestBuilder is request builder for EmbeddedSIMActivationCodePool
type EmbeddedSIMActivationCodePoolRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmbeddedSIMActivationCodePoolRequest
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Request() *EmbeddedSIMActivationCodePoolRequest {
	return &EmbeddedSIMActivationCodePoolRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmbeddedSIMActivationCodePoolRequest is request for EmbeddedSIMActivationCodePool
type EmbeddedSIMActivationCodePoolRequest struct{ BaseRequest }

// Get performs GET request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Get(ctx context.Context) (resObj *EmbeddedSIMActivationCodePool, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Update(ctx context.Context, reqObj *EmbeddedSIMActivationCodePool) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmbeddedSIMActivationCodePool
func (r *EmbeddedSIMActivationCodePoolRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for EmbeddedSIMActivationCodePoolAssignment collection
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Assignments() *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder is request builder for EmbeddedSIMActivationCodePoolAssignment collection
type EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmbeddedSIMActivationCodePoolAssignment collection
func (b *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder) Request() *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest {
	return &EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmbeddedSIMActivationCodePoolAssignment item
func (b *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequestBuilder) ID(id string) *EmbeddedSIMActivationCodePoolAssignmentRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest is request for EmbeddedSIMActivationCodePoolAssignment collection
type EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
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
	var values []EmbeddedSIMActivationCodePoolAssignment
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
			value  []EmbeddedSIMActivationCodePoolAssignment
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

// Get performs GET request for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Get(ctx context.Context) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for EmbeddedSIMActivationCodePoolAssignment collection
func (r *EmbeddedSIMActivationCodePoolAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *EmbeddedSIMActivationCodePoolAssignment) (resObj *EmbeddedSIMActivationCodePoolAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceStates returns request builder for EmbeddedSIMDeviceState collection
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) DeviceStates() *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceStates"
	return bb
}

// EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder is request builder for EmbeddedSIMDeviceState collection
type EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmbeddedSIMDeviceState collection
func (b *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder) Request() *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest {
	return &EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmbeddedSIMDeviceState item
func (b *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequestBuilder) ID(id string) *EmbeddedSIMDeviceStateRequestBuilder {
	bb := &EmbeddedSIMDeviceStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest is request for EmbeddedSIMDeviceState collection
type EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]EmbeddedSIMDeviceState, error) {
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
	var values []EmbeddedSIMDeviceState
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
			value  []EmbeddedSIMDeviceState
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

// Get performs GET request for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Get(ctx context.Context) ([]EmbeddedSIMDeviceState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for EmbeddedSIMDeviceState collection
func (r *EmbeddedSIMActivationCodePoolDeviceStatesCollectionRequest) Add(ctx context.Context, reqObj *EmbeddedSIMDeviceState) (resObj *EmbeddedSIMDeviceState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

//
type EmbeddedSIMActivationCodePoolAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *EmbeddedSIMActivationCodePoolRequestBuilder) Assign(reqObj *EmbeddedSIMActivationCodePoolAssignRequestParameter) *EmbeddedSIMActivationCodePoolAssignRequestBuilder {
	bb := &EmbeddedSIMActivationCodePoolAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type EmbeddedSIMActivationCodePoolAssignRequest struct{ BaseRequest }

//
func (b *EmbeddedSIMActivationCodePoolAssignRequestBuilder) Request() *EmbeddedSIMActivationCodePoolAssignRequest {
	return &EmbeddedSIMActivationCodePoolAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *EmbeddedSIMActivationCodePoolAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
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
	var values []EmbeddedSIMActivationCodePoolAssignment
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
			value  []EmbeddedSIMActivationCodePoolAssignment
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
func (r *EmbeddedSIMActivationCodePoolAssignRequest) Post(ctx context.Context) ([]EmbeddedSIMActivationCodePoolAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}

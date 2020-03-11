// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// CloudCommunications undocumented
type CloudCommunications struct {
	// Entity is the base model of CloudCommunications
	Entity
	// Calls undocumented
	Calls []Call `json:"calls,omitempty"`
	// OnlineMeetings undocumented
	OnlineMeetings []OnlineMeeting `json:"onlineMeetings,omitempty"`
}

// CloudCommunicationsGetPresencesByUserIDRequestParameter undocumented
type CloudCommunicationsGetPresencesByUserIDRequestParameter struct {
	// IDs undocumented
	IDs []string `json:"ids,omitempty"`
}

// CloudCommunicationsRequestBuilder is request builder for CloudCommunications
type CloudCommunicationsRequestBuilder struct{ BaseRequestBuilder }

// Request returns CloudCommunicationsRequest
func (b *CloudCommunicationsRequestBuilder) Request() *CloudCommunicationsRequest {
	return &CloudCommunicationsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CloudCommunicationsRequest is request for CloudCommunications
type CloudCommunicationsRequest struct{ BaseRequest }

// Get performs GET request for CloudCommunications
func (r *CloudCommunicationsRequest) Get(ctx context.Context) (resObj *CloudCommunications, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CloudCommunications
func (r *CloudCommunicationsRequest) Update(ctx context.Context, reqObj *CloudCommunications) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CloudCommunications
func (r *CloudCommunicationsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Calls returns request builder for Call collection
func (b *CloudCommunicationsRequestBuilder) Calls() *CloudCommunicationsCallsCollectionRequestBuilder {
	bb := &CloudCommunicationsCallsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/calls"
	return bb
}

// CloudCommunicationsCallsCollectionRequestBuilder is request builder for Call collection
type CloudCommunicationsCallsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Call collection
func (b *CloudCommunicationsCallsCollectionRequestBuilder) Request() *CloudCommunicationsCallsCollectionRequest {
	return &CloudCommunicationsCallsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Call item
func (b *CloudCommunicationsCallsCollectionRequestBuilder) ID(id string) *CallRequestBuilder {
	bb := &CallRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CloudCommunicationsCallsCollectionRequest is request for Call collection
type CloudCommunicationsCallsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Call, error) {
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
	var values []Call
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
			value  []Call
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

// Get performs GET request for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Get(ctx context.Context) ([]Call, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Call collection
func (r *CloudCommunicationsCallsCollectionRequest) Add(ctx context.Context, reqObj *Call) (resObj *Call, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OnlineMeetings returns request builder for OnlineMeeting collection
func (b *CloudCommunicationsRequestBuilder) OnlineMeetings() *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder {
	bb := &CloudCommunicationsOnlineMeetingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/onlineMeetings"
	return bb
}

// CloudCommunicationsOnlineMeetingsCollectionRequestBuilder is request builder for OnlineMeeting collection
type CloudCommunicationsOnlineMeetingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnlineMeeting collection
func (b *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder) Request() *CloudCommunicationsOnlineMeetingsCollectionRequest {
	return &CloudCommunicationsOnlineMeetingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnlineMeeting item
func (b *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder) ID(id string) *OnlineMeetingRequestBuilder {
	bb := &OnlineMeetingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CloudCommunicationsOnlineMeetingsCollectionRequest is request for OnlineMeeting collection
type CloudCommunicationsOnlineMeetingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OnlineMeeting, error) {
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
	var values []OnlineMeeting
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
			value  []OnlineMeeting
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

// Get performs GET request for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Get(ctx context.Context) ([]OnlineMeeting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OnlineMeeting collection
func (r *CloudCommunicationsOnlineMeetingsCollectionRequest) Add(ctx context.Context, reqObj *OnlineMeeting) (resObj *OnlineMeeting, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

//
type CloudCommunicationsGetPresencesByUserIDRequestBuilder struct{ BaseRequestBuilder }

// GetPresencesByUserID action undocumented
func (b *CloudCommunicationsRequestBuilder) GetPresencesByUserID(reqObj *CloudCommunicationsGetPresencesByUserIDRequestParameter) *CloudCommunicationsGetPresencesByUserIDRequestBuilder {
	bb := &CloudCommunicationsGetPresencesByUserIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getPresencesByUserId"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CloudCommunicationsGetPresencesByUserIDRequest struct{ BaseRequest }

//
func (b *CloudCommunicationsGetPresencesByUserIDRequestBuilder) Request() *CloudCommunicationsGetPresencesByUserIDRequest {
	return &CloudCommunicationsGetPresencesByUserIDRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CloudCommunicationsGetPresencesByUserIDRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Presence, error) {
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
	var values []Presence
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
			value  []Presence
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
func (r *CloudCommunicationsGetPresencesByUserIDRequest) Post(ctx context.Context) ([]Presence, error) {
	return r.Paging(ctx, "POST", "", r.requestObject)
}

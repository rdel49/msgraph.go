// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// RoleAssignmentRequestBuilder is request builder for RoleAssignment
type RoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleAssignmentRequest
func (b *RoleAssignmentRequestBuilder) Request() *RoleAssignmentRequest {
	return &RoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleAssignmentRequest is request for RoleAssignment
type RoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for RoleAssignment
func (r *RoleAssignmentRequest) Get(ctx context.Context) (resObj *RoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleAssignment
func (r *RoleAssignmentRequest) Update(ctx context.Context, reqObj *RoleAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleAssignment
func (r *RoleAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RoleDefinitionRequestBuilder is request builder for RoleDefinition
type RoleDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleDefinitionRequest
func (b *RoleDefinitionRequestBuilder) Request() *RoleDefinitionRequest {
	return &RoleDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleDefinitionRequest is request for RoleDefinition
type RoleDefinitionRequest struct{ BaseRequest }

// Get performs GET request for RoleDefinition
func (r *RoleDefinitionRequest) Get(ctx context.Context) (resObj *RoleDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleDefinition
func (r *RoleDefinitionRequest) Update(ctx context.Context, reqObj *RoleDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleDefinition
func (r *RoleDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RoleManagementRequestBuilder is request builder for RoleManagement
type RoleManagementRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleManagementRequest
func (b *RoleManagementRequestBuilder) Request() *RoleManagementRequest {
	return &RoleManagementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleManagementRequest is request for RoleManagement
type RoleManagementRequest struct{ BaseRequest }

// Get performs GET request for RoleManagement
func (r *RoleManagementRequest) Get(ctx context.Context) (resObj *RoleManagement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleManagement
func (r *RoleManagementRequest) Update(ctx context.Context, reqObj *RoleManagement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleManagement
func (r *RoleManagementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RoleScopeTagRequestBuilder is request builder for RoleScopeTag
type RoleScopeTagRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleScopeTagRequest
func (b *RoleScopeTagRequestBuilder) Request() *RoleScopeTagRequest {
	return &RoleScopeTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleScopeTagRequest is request for RoleScopeTag
type RoleScopeTagRequest struct{ BaseRequest }

// Get performs GET request for RoleScopeTag
func (r *RoleScopeTagRequest) Get(ctx context.Context) (resObj *RoleScopeTag, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleScopeTag
func (r *RoleScopeTagRequest) Update(ctx context.Context, reqObj *RoleScopeTag) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleScopeTag
func (r *RoleScopeTagRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RoleScopeTagAutoAssignmentRequestBuilder is request builder for RoleScopeTagAutoAssignment
type RoleScopeTagAutoAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleScopeTagAutoAssignmentRequest
func (b *RoleScopeTagAutoAssignmentRequestBuilder) Request() *RoleScopeTagAutoAssignmentRequest {
	return &RoleScopeTagAutoAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleScopeTagAutoAssignmentRequest is request for RoleScopeTagAutoAssignment
type RoleScopeTagAutoAssignmentRequest struct{ BaseRequest }

// Get performs GET request for RoleScopeTagAutoAssignment
func (r *RoleScopeTagAutoAssignmentRequest) Get(ctx context.Context) (resObj *RoleScopeTagAutoAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleScopeTagAutoAssignment
func (r *RoleScopeTagAutoAssignmentRequest) Update(ctx context.Context, reqObj *RoleScopeTagAutoAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleScopeTagAutoAssignment
func (r *RoleScopeTagAutoAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder struct{ BaseRequestBuilder }

// GetRoleScopeTagsByID action undocumented
func (b *DeviceAndAppManagementRoleAssignmentRoleScopeTagsCollectionRequestBuilder) GetRoleScopeTagsByID(reqObj *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestParameter) *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder {
	bb := &RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getRoleScopeTagsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// GetRoleScopeTagsByID action undocumented
func (b *DeviceManagementRoleScopeTagsCollectionRequestBuilder) GetRoleScopeTagsByID(reqObj *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestParameter) *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder {
	bb := &RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getRoleScopeTagsById"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RoleScopeTagCollectionGetRoleScopeTagsByIDRequest struct{ BaseRequest }

//
func (b *RoleScopeTagCollectionGetRoleScopeTagsByIDRequestBuilder) Request() *RoleScopeTagCollectionGetRoleScopeTagsByIDRequest {
	return &RoleScopeTagCollectionGetRoleScopeTagsByIDRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *RoleScopeTagCollectionGetRoleScopeTagsByIDRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RoleScopeTag, error) {
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
	var values []RoleScopeTag
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
			value  []RoleScopeTag
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

//
func (r *RoleScopeTagCollectionGetRoleScopeTagsByIDRequest) PostN(ctx context.Context, n int) ([]RoleScopeTag, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *RoleScopeTagCollectionGetRoleScopeTagsByIDRequest) Post(ctx context.Context) ([]RoleScopeTag, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type RoleScopeTagAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *RoleScopeTagRequestBuilder) Assign(reqObj *RoleScopeTagAssignRequestParameter) *RoleScopeTagAssignRequestBuilder {
	bb := &RoleScopeTagAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RoleScopeTagAssignRequest struct{ BaseRequest }

//
func (b *RoleScopeTagAssignRequestBuilder) Request() *RoleScopeTagAssignRequest {
	return &RoleScopeTagAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *RoleScopeTagAssignRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RoleScopeTagAutoAssignment, error) {
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
	var values []RoleScopeTagAutoAssignment
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
			value  []RoleScopeTagAutoAssignment
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

//
func (r *RoleScopeTagAssignRequest) PostN(ctx context.Context, n int) ([]RoleScopeTagAutoAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *RoleScopeTagAssignRequest) Post(ctx context.Context) ([]RoleScopeTagAutoAssignment, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

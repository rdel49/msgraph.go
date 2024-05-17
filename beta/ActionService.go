// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rdel49/msgraph.go/jsonx"
)

// ServicePrincipalCreatePasswordSingleSignOnCredentialsRequestParameter undocumented
type ServicePrincipalCreatePasswordSingleSignOnCredentialsRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Credentials undocumented
	Credentials []Credential `json:"credentials,omitempty"`
}

// ServicePrincipalGetPasswordSingleSignOnCredentialsRequestParameter undocumented
type ServicePrincipalGetPasswordSingleSignOnCredentialsRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

// ServicePrincipalDeletePasswordSingleSignOnCredentialsRequestParameter undocumented
type ServicePrincipalDeletePasswordSingleSignOnCredentialsRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

// ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequestParameter undocumented
type ServicePrincipalUpdatePasswordSingleSignOnCredentialsRequestParameter struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Credentials undocumented
	Credentials []Credential `json:"credentials,omitempty"`
}

// AppRoleAssignedTo returns request builder for AppRoleAssignment collection
func (b *ServicePrincipalRequestBuilder) AppRoleAssignedTo() *ServicePrincipalAppRoleAssignedToCollectionRequestBuilder {
	bb := &ServicePrincipalAppRoleAssignedToCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appRoleAssignedTo"
	return bb
}

// ServicePrincipalAppRoleAssignedToCollectionRequestBuilder is request builder for AppRoleAssignment collection
type ServicePrincipalAppRoleAssignedToCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppRoleAssignment collection
func (b *ServicePrincipalAppRoleAssignedToCollectionRequestBuilder) Request() *ServicePrincipalAppRoleAssignedToCollectionRequest {
	return &ServicePrincipalAppRoleAssignedToCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppRoleAssignment item
func (b *ServicePrincipalAppRoleAssignedToCollectionRequestBuilder) ID(id string) *AppRoleAssignmentRequestBuilder {
	bb := &AppRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalAppRoleAssignedToCollectionRequest is request for AppRoleAssignment collection
type ServicePrincipalAppRoleAssignedToCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignedToCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AppRoleAssignment, error) {
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
	var values []AppRoleAssignment
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
			value  []AppRoleAssignment
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

// GetN performs GET request for AppRoleAssignment collection, max N pages
func (r *ServicePrincipalAppRoleAssignedToCollectionRequest) GetN(ctx context.Context, n int) ([]AppRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignedToCollectionRequest) Get(ctx context.Context) ([]AppRoleAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignedToCollectionRequest) Add(ctx context.Context, reqObj *AppRoleAssignment) (resObj *AppRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AppRoleAssignments returns request builder for AppRoleAssignment collection
func (b *ServicePrincipalRequestBuilder) AppRoleAssignments() *ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder {
	bb := &ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appRoleAssignments"
	return bb
}

// ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder is request builder for AppRoleAssignment collection
type ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppRoleAssignment collection
func (b *ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder) Request() *ServicePrincipalAppRoleAssignmentsCollectionRequest {
	return &ServicePrincipalAppRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppRoleAssignment item
func (b *ServicePrincipalAppRoleAssignmentsCollectionRequestBuilder) ID(id string) *AppRoleAssignmentRequestBuilder {
	bb := &AppRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalAppRoleAssignmentsCollectionRequest is request for AppRoleAssignment collection
type ServicePrincipalAppRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AppRoleAssignment, error) {
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
	var values []AppRoleAssignment
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
			value  []AppRoleAssignment
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

// GetN performs GET request for AppRoleAssignment collection, max N pages
func (r *ServicePrincipalAppRoleAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]AppRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignmentsCollectionRequest) Get(ctx context.Context) ([]AppRoleAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AppRoleAssignment collection
func (r *ServicePrincipalAppRoleAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *AppRoleAssignment) (resObj *AppRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// CreatedObjects returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) CreatedObjects() *ServicePrincipalCreatedObjectsCollectionRequestBuilder {
	bb := &ServicePrincipalCreatedObjectsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/createdObjects"
	return bb
}

// ServicePrincipalCreatedObjectsCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalCreatedObjectsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalCreatedObjectsCollectionRequestBuilder) Request() *ServicePrincipalCreatedObjectsCollectionRequest {
	return &ServicePrincipalCreatedObjectsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalCreatedObjectsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalCreatedObjectsCollectionRequest is request for DirectoryObject collection
type ServicePrincipalCreatedObjectsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalCreatedObjectsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ServicePrincipalCreatedObjectsCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalCreatedObjectsCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalCreatedObjectsCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// LicenseDetails returns request builder for LicenseDetails collection
func (b *ServicePrincipalRequestBuilder) LicenseDetails() *ServicePrincipalLicenseDetailsCollectionRequestBuilder {
	bb := &ServicePrincipalLicenseDetailsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/licenseDetails"
	return bb
}

// ServicePrincipalLicenseDetailsCollectionRequestBuilder is request builder for LicenseDetails collection
type ServicePrincipalLicenseDetailsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for LicenseDetails collection
func (b *ServicePrincipalLicenseDetailsCollectionRequestBuilder) Request() *ServicePrincipalLicenseDetailsCollectionRequest {
	return &ServicePrincipalLicenseDetailsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for LicenseDetails item
func (b *ServicePrincipalLicenseDetailsCollectionRequestBuilder) ID(id string) *LicenseDetailsRequestBuilder {
	bb := &LicenseDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalLicenseDetailsCollectionRequest is request for LicenseDetails collection
type ServicePrincipalLicenseDetailsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for LicenseDetails collection
func (r *ServicePrincipalLicenseDetailsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]LicenseDetails, error) {
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
	var values []LicenseDetails
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
			value  []LicenseDetails
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

// GetN performs GET request for LicenseDetails collection, max N pages
func (r *ServicePrincipalLicenseDetailsCollectionRequest) GetN(ctx context.Context, n int) ([]LicenseDetails, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for LicenseDetails collection
func (r *ServicePrincipalLicenseDetailsCollectionRequest) Get(ctx context.Context) ([]LicenseDetails, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for LicenseDetails collection
func (r *ServicePrincipalLicenseDetailsCollectionRequest) Add(ctx context.Context, reqObj *LicenseDetails) (resObj *LicenseDetails, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MemberOf returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) MemberOf() *ServicePrincipalMemberOfCollectionRequestBuilder {
	bb := &ServicePrincipalMemberOfCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/memberOf"
	return bb
}

// ServicePrincipalMemberOfCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalMemberOfCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalMemberOfCollectionRequestBuilder) Request() *ServicePrincipalMemberOfCollectionRequest {
	return &ServicePrincipalMemberOfCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalMemberOfCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalMemberOfCollectionRequest is request for DirectoryObject collection
type ServicePrincipalMemberOfCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalMemberOfCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ServicePrincipalMemberOfCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalMemberOfCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalMemberOfCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OAuth2PermissionGrants returns request builder for OAuth2PermissionGrant collection
func (b *ServicePrincipalRequestBuilder) OAuth2PermissionGrants() *ServicePrincipalOAuth2PermissionGrantsCollectionRequestBuilder {
	bb := &ServicePrincipalOAuth2PermissionGrantsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/oauth2PermissionGrants"
	return bb
}

// ServicePrincipalOAuth2PermissionGrantsCollectionRequestBuilder is request builder for OAuth2PermissionGrant collection
type ServicePrincipalOAuth2PermissionGrantsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OAuth2PermissionGrant collection
func (b *ServicePrincipalOAuth2PermissionGrantsCollectionRequestBuilder) Request() *ServicePrincipalOAuth2PermissionGrantsCollectionRequest {
	return &ServicePrincipalOAuth2PermissionGrantsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OAuth2PermissionGrant item
func (b *ServicePrincipalOAuth2PermissionGrantsCollectionRequestBuilder) ID(id string) *OAuth2PermissionGrantRequestBuilder {
	bb := &OAuth2PermissionGrantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalOAuth2PermissionGrantsCollectionRequest is request for OAuth2PermissionGrant collection
type ServicePrincipalOAuth2PermissionGrantsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OAuth2PermissionGrant collection
func (r *ServicePrincipalOAuth2PermissionGrantsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OAuth2PermissionGrant, error) {
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
	var values []OAuth2PermissionGrant
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
			value  []OAuth2PermissionGrant
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

// GetN performs GET request for OAuth2PermissionGrant collection, max N pages
func (r *ServicePrincipalOAuth2PermissionGrantsCollectionRequest) GetN(ctx context.Context, n int) ([]OAuth2PermissionGrant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OAuth2PermissionGrant collection
func (r *ServicePrincipalOAuth2PermissionGrantsCollectionRequest) Get(ctx context.Context) ([]OAuth2PermissionGrant, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OAuth2PermissionGrant collection
func (r *ServicePrincipalOAuth2PermissionGrantsCollectionRequest) Add(ctx context.Context, reqObj *OAuth2PermissionGrant) (resObj *OAuth2PermissionGrant, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OwnedObjects returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) OwnedObjects() *ServicePrincipalOwnedObjectsCollectionRequestBuilder {
	bb := &ServicePrincipalOwnedObjectsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/ownedObjects"
	return bb
}

// ServicePrincipalOwnedObjectsCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalOwnedObjectsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalOwnedObjectsCollectionRequestBuilder) Request() *ServicePrincipalOwnedObjectsCollectionRequest {
	return &ServicePrincipalOwnedObjectsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalOwnedObjectsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalOwnedObjectsCollectionRequest is request for DirectoryObject collection
type ServicePrincipalOwnedObjectsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalOwnedObjectsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ServicePrincipalOwnedObjectsCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalOwnedObjectsCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalOwnedObjectsCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Owners returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) Owners() *ServicePrincipalOwnersCollectionRequestBuilder {
	bb := &ServicePrincipalOwnersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/owners"
	return bb
}

// ServicePrincipalOwnersCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalOwnersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalOwnersCollectionRequestBuilder) Request() *ServicePrincipalOwnersCollectionRequest {
	return &ServicePrincipalOwnersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalOwnersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalOwnersCollectionRequest is request for DirectoryObject collection
type ServicePrincipalOwnersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalOwnersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ServicePrincipalOwnersCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalOwnersCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalOwnersCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Policies returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) Policies() *ServicePrincipalPoliciesCollectionRequestBuilder {
	bb := &ServicePrincipalPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policies"
	return bb
}

// ServicePrincipalPoliciesCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalPoliciesCollectionRequestBuilder) Request() *ServicePrincipalPoliciesCollectionRequest {
	return &ServicePrincipalPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalPoliciesCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalPoliciesCollectionRequest is request for DirectoryObject collection
type ServicePrincipalPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ServicePrincipalPoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalPoliciesCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalPoliciesCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Synchronization is navigation property
func (b *ServicePrincipalRequestBuilder) Synchronization() *SynchronizationRequestBuilder {
	bb := &SynchronizationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/synchronization"
	return bb
}

// TransitiveMemberOf returns request builder for DirectoryObject collection
func (b *ServicePrincipalRequestBuilder) TransitiveMemberOf() *ServicePrincipalTransitiveMemberOfCollectionRequestBuilder {
	bb := &ServicePrincipalTransitiveMemberOfCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/transitiveMemberOf"
	return bb
}

// ServicePrincipalTransitiveMemberOfCollectionRequestBuilder is request builder for DirectoryObject collection
type ServicePrincipalTransitiveMemberOfCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ServicePrincipalTransitiveMemberOfCollectionRequestBuilder) Request() *ServicePrincipalTransitiveMemberOfCollectionRequest {
	return &ServicePrincipalTransitiveMemberOfCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ServicePrincipalTransitiveMemberOfCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ServicePrincipalTransitiveMemberOfCollectionRequest is request for DirectoryObject collection
type ServicePrincipalTransitiveMemberOfCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ServicePrincipalTransitiveMemberOfCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ServicePrincipalTransitiveMemberOfCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ServicePrincipalTransitiveMemberOfCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ServicePrincipalTransitiveMemberOfCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

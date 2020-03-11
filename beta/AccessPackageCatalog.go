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

// AccessPackageCatalog undocumented
type AccessPackageCatalog struct {
	// Entity is the base model of AccessPackageCatalog
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// CatalogType undocumented
	CatalogType *string `json:"catalogType,omitempty"`
	// CatalogStatus undocumented
	CatalogStatus *string `json:"catalogStatus,omitempty"`
	// IsExternallyVisible undocumented
	IsExternallyVisible *bool `json:"isExternallyVisible,omitempty"`
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// AccessPackageResources undocumented
	AccessPackageResources []AccessPackageResource `json:"accessPackageResources,omitempty"`
	// AccessPackageResourceRoles undocumented
	AccessPackageResourceRoles []AccessPackageResourceRole `json:"accessPackageResourceRoles,omitempty"`
	// AccessPackageResourceScopes undocumented
	AccessPackageResourceScopes []AccessPackageResourceScope `json:"accessPackageResourceScopes,omitempty"`
	// AccessPackages undocumented
	AccessPackages []AccessPackage `json:"accessPackages,omitempty"`
}

// AccessPackageCatalogRequestBuilder is request builder for AccessPackageCatalog
type AccessPackageCatalogRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageCatalogRequest
func (b *AccessPackageCatalogRequestBuilder) Request() *AccessPackageCatalogRequest {
	return &AccessPackageCatalogRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageCatalogRequest is request for AccessPackageCatalog
type AccessPackageCatalogRequest struct{ BaseRequest }

// Get performs GET request for AccessPackageCatalog
func (r *AccessPackageCatalogRequest) Get(ctx context.Context) (resObj *AccessPackageCatalog, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessPackageCatalog
func (r *AccessPackageCatalogRequest) Update(ctx context.Context, reqObj *AccessPackageCatalog) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessPackageCatalog
func (r *AccessPackageCatalogRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccessPackageResourceRoles returns request builder for AccessPackageResourceRole collection
func (b *AccessPackageCatalogRequestBuilder) AccessPackageResourceRoles() *AccessPackageCatalogAccessPackageResourceRolesCollectionRequestBuilder {
	bb := &AccessPackageCatalogAccessPackageResourceRolesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceRoles"
	return bb
}

// AccessPackageCatalogAccessPackageResourceRolesCollectionRequestBuilder is request builder for AccessPackageResourceRole collection
type AccessPackageCatalogAccessPackageResourceRolesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResourceRole collection
func (b *AccessPackageCatalogAccessPackageResourceRolesCollectionRequestBuilder) Request() *AccessPackageCatalogAccessPackageResourceRolesCollectionRequest {
	return &AccessPackageCatalogAccessPackageResourceRolesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResourceRole item
func (b *AccessPackageCatalogAccessPackageResourceRolesCollectionRequestBuilder) ID(id string) *AccessPackageResourceRoleRequestBuilder {
	bb := &AccessPackageResourceRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageCatalogAccessPackageResourceRolesCollectionRequest is request for AccessPackageResourceRole collection
type AccessPackageCatalogAccessPackageResourceRolesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageResourceRole collection
func (r *AccessPackageCatalogAccessPackageResourceRolesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageResourceRole, error) {
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
	var values []AccessPackageResourceRole
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
			value  []AccessPackageResourceRole
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

// Get performs GET request for AccessPackageResourceRole collection
func (r *AccessPackageCatalogAccessPackageResourceRolesCollectionRequest) Get(ctx context.Context) ([]AccessPackageResourceRole, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageResourceRole collection
func (r *AccessPackageCatalogAccessPackageResourceRolesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageResourceRole) (resObj *AccessPackageResourceRole, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageResourceScopes returns request builder for AccessPackageResourceScope collection
func (b *AccessPackageCatalogRequestBuilder) AccessPackageResourceScopes() *AccessPackageCatalogAccessPackageResourceScopesCollectionRequestBuilder {
	bb := &AccessPackageCatalogAccessPackageResourceScopesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceScopes"
	return bb
}

// AccessPackageCatalogAccessPackageResourceScopesCollectionRequestBuilder is request builder for AccessPackageResourceScope collection
type AccessPackageCatalogAccessPackageResourceScopesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResourceScope collection
func (b *AccessPackageCatalogAccessPackageResourceScopesCollectionRequestBuilder) Request() *AccessPackageCatalogAccessPackageResourceScopesCollectionRequest {
	return &AccessPackageCatalogAccessPackageResourceScopesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResourceScope item
func (b *AccessPackageCatalogAccessPackageResourceScopesCollectionRequestBuilder) ID(id string) *AccessPackageResourceScopeRequestBuilder {
	bb := &AccessPackageResourceScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageCatalogAccessPackageResourceScopesCollectionRequest is request for AccessPackageResourceScope collection
type AccessPackageCatalogAccessPackageResourceScopesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageResourceScope collection
func (r *AccessPackageCatalogAccessPackageResourceScopesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageResourceScope, error) {
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
	var values []AccessPackageResourceScope
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
			value  []AccessPackageResourceScope
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

// Get performs GET request for AccessPackageResourceScope collection
func (r *AccessPackageCatalogAccessPackageResourceScopesCollectionRequest) Get(ctx context.Context) ([]AccessPackageResourceScope, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageResourceScope collection
func (r *AccessPackageCatalogAccessPackageResourceScopesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageResourceScope) (resObj *AccessPackageResourceScope, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageResources returns request builder for AccessPackageResource collection
func (b *AccessPackageCatalogRequestBuilder) AccessPackageResources() *AccessPackageCatalogAccessPackageResourcesCollectionRequestBuilder {
	bb := &AccessPackageCatalogAccessPackageResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResources"
	return bb
}

// AccessPackageCatalogAccessPackageResourcesCollectionRequestBuilder is request builder for AccessPackageResource collection
type AccessPackageCatalogAccessPackageResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResource collection
func (b *AccessPackageCatalogAccessPackageResourcesCollectionRequestBuilder) Request() *AccessPackageCatalogAccessPackageResourcesCollectionRequest {
	return &AccessPackageCatalogAccessPackageResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResource item
func (b *AccessPackageCatalogAccessPackageResourcesCollectionRequestBuilder) ID(id string) *AccessPackageResourceRequestBuilder {
	bb := &AccessPackageResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageCatalogAccessPackageResourcesCollectionRequest is request for AccessPackageResource collection
type AccessPackageCatalogAccessPackageResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageResource collection
func (r *AccessPackageCatalogAccessPackageResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageResource, error) {
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
	var values []AccessPackageResource
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
			value  []AccessPackageResource
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

// Get performs GET request for AccessPackageResource collection
func (r *AccessPackageCatalogAccessPackageResourcesCollectionRequest) Get(ctx context.Context) ([]AccessPackageResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageResource collection
func (r *AccessPackageCatalogAccessPackageResourcesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageResource) (resObj *AccessPackageResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackages returns request builder for AccessPackage collection
func (b *AccessPackageCatalogRequestBuilder) AccessPackages() *AccessPackageCatalogAccessPackagesCollectionRequestBuilder {
	bb := &AccessPackageCatalogAccessPackagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackages"
	return bb
}

// AccessPackageCatalogAccessPackagesCollectionRequestBuilder is request builder for AccessPackage collection
type AccessPackageCatalogAccessPackagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackage collection
func (b *AccessPackageCatalogAccessPackagesCollectionRequestBuilder) Request() *AccessPackageCatalogAccessPackagesCollectionRequest {
	return &AccessPackageCatalogAccessPackagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackage item
func (b *AccessPackageCatalogAccessPackagesCollectionRequestBuilder) ID(id string) *AccessPackageRequestBuilder {
	bb := &AccessPackageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageCatalogAccessPackagesCollectionRequest is request for AccessPackage collection
type AccessPackageCatalogAccessPackagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackage collection
func (r *AccessPackageCatalogAccessPackagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackage, error) {
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
	var values []AccessPackage
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
			value  []AccessPackage
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

// Get performs GET request for AccessPackage collection
func (r *AccessPackageCatalogAccessPackagesCollectionRequest) Get(ctx context.Context) ([]AccessPackage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackage collection
func (r *AccessPackageCatalogAccessPackagesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackage) (resObj *AccessPackage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

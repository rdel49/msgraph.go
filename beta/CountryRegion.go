// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// CountryRegion undocumented
type CountryRegion struct {
	// Entity is the base model of CountryRegion
	Entity
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// AddressFormat undocumented
	AddressFormat *string `json:"addressFormat,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// CountryRegionRequestBuilder is request builder for CountryRegion
type CountryRegionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CountryRegionRequest
func (b *CountryRegionRequestBuilder) Request() *CountryRegionRequest {
	return &CountryRegionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CountryRegionRequest is request for CountryRegion
type CountryRegionRequest struct{ BaseRequest }

// Get performs GET request for CountryRegion
func (r *CountryRegionRequest) Get(ctx context.Context) (resObj *CountryRegion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CountryRegion
func (r *CountryRegionRequest) Update(ctx context.Context, reqObj *CountryRegion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CountryRegion
func (r *CountryRegionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

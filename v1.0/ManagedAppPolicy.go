// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// ManagedAppPolicy The ManagedAppPolicy resource represents a base type for platform specific policies.
type ManagedAppPolicy struct {
	// Entity is the base model of ManagedAppPolicy
	Entity
	// DisplayName Policy display name.
	DisplayName *string `json:"displayName,omitempty"`
	// Description The policy's description.
	Description *string `json:"description,omitempty"`
	// CreatedDateTime The date and time the policy was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime Last time the policy was modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Version Version of the entity.
	Version *string `json:"version,omitempty"`
}

// ManagedAppPolicyTargetAppsRequestParameter undocumented
type ManagedAppPolicyTargetAppsRequestParameter struct {
	// Apps undocumented
	Apps []ManagedMobileApp `json:"apps,omitempty"`
}

// ManagedAppPolicyRequestBuilder is request builder for ManagedAppPolicy
type ManagedAppPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppPolicyRequest
func (b *ManagedAppPolicyRequestBuilder) Request() *ManagedAppPolicyRequest {
	return &ManagedAppPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppPolicyRequest is request for ManagedAppPolicy
type ManagedAppPolicyRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Get(ctx context.Context) (resObj *ManagedAppPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Update(ctx context.Context, reqObj *ManagedAppPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppPolicy
func (r *ManagedAppPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ManagedAppPolicyTargetAppsRequestBuilder struct{ BaseRequestBuilder }

// TargetApps action undocumented
func (b *ManagedAppPolicyRequestBuilder) TargetApps(reqObj *ManagedAppPolicyTargetAppsRequestParameter) *ManagedAppPolicyTargetAppsRequestBuilder {
	bb := &ManagedAppPolicyTargetAppsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/targetApps"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedAppPolicyTargetAppsRequest struct{ BaseRequest }

//
func (b *ManagedAppPolicyTargetAppsRequestBuilder) Request() *ManagedAppPolicyTargetAppsRequest {
	return &ManagedAppPolicyTargetAppsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedAppPolicyTargetAppsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

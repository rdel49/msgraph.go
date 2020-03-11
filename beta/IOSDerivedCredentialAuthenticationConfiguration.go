// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// IOSDerivedCredentialAuthenticationConfiguration iOS Derived Credential profile.
type IOSDerivedCredentialAuthenticationConfiguration struct {
	// DeviceConfiguration is the base model of IOSDerivedCredentialAuthenticationConfiguration
	DeviceConfiguration
	// DerivedCredentialSettings undocumented
	DerivedCredentialSettings *DeviceManagementDerivedCredentialSettings `json:"derivedCredentialSettings,omitempty"`
}

// IOSDerivedCredentialAuthenticationConfigurationRequestBuilder is request builder for IOSDerivedCredentialAuthenticationConfiguration
type IOSDerivedCredentialAuthenticationConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSDerivedCredentialAuthenticationConfigurationRequest
func (b *IOSDerivedCredentialAuthenticationConfigurationRequestBuilder) Request() *IOSDerivedCredentialAuthenticationConfigurationRequest {
	return &IOSDerivedCredentialAuthenticationConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSDerivedCredentialAuthenticationConfigurationRequest is request for IOSDerivedCredentialAuthenticationConfiguration
type IOSDerivedCredentialAuthenticationConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Get(ctx context.Context) (resObj *IOSDerivedCredentialAuthenticationConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Update(ctx context.Context, reqObj *IOSDerivedCredentialAuthenticationConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DerivedCredentialSettings is navigation property
func (b *IOSDerivedCredentialAuthenticationConfigurationRequestBuilder) DerivedCredentialSettings() *DeviceManagementDerivedCredentialSettingsRequestBuilder {
	bb := &DeviceManagementDerivedCredentialSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/derivedCredentialSettings"
	return bb
}

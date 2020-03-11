// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MacOsVppAppAssignedLicense MacOS Volume Purchase Program license assignment. This class does not support Create, Delete, or Update.
type MacOsVppAppAssignedLicense struct {
	// Entity is the base model of MacOsVppAppAssignedLicense
	Entity
	// UserEmailAddress The user email address.
	UserEmailAddress *string `json:"userEmailAddress,omitempty"`
	// UserID The user ID.
	UserID *string `json:"userId,omitempty"`
	// UserName The user name.
	UserName *string `json:"userName,omitempty"`
	// UserPrincipalName The user principal name.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// MacOsVppAppAssignedLicenseRequestBuilder is request builder for MacOsVppAppAssignedLicense
type MacOsVppAppAssignedLicenseRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOsVppAppAssignedLicenseRequest
func (b *MacOsVppAppAssignedLicenseRequestBuilder) Request() *MacOsVppAppAssignedLicenseRequest {
	return &MacOsVppAppAssignedLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOsVppAppAssignedLicenseRequest is request for MacOsVppAppAssignedLicense
type MacOsVppAppAssignedLicenseRequest struct{ BaseRequest }

// Get performs GET request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Get(ctx context.Context) (resObj *MacOsVppAppAssignedLicense, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Update(ctx context.Context, reqObj *MacOsVppAppAssignedLicense) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

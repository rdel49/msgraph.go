// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// OrganizationSetMobileDeviceManagementAuthorityRequestParameter undocumented
type OrganizationSetMobileDeviceManagementAuthorityRequestParameter struct {
}

//
type OrganizationSetMobileDeviceManagementAuthorityRequestBuilder struct{ BaseRequestBuilder }

// SetMobileDeviceManagementAuthority action undocumented
func (b *OrganizationRequestBuilder) SetMobileDeviceManagementAuthority(reqObj *OrganizationSetMobileDeviceManagementAuthorityRequestParameter) *OrganizationSetMobileDeviceManagementAuthorityRequestBuilder {
	bb := &OrganizationSetMobileDeviceManagementAuthorityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setMobileDeviceManagementAuthority"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type OrganizationSetMobileDeviceManagementAuthorityRequest struct{ BaseRequest }

//
func (b *OrganizationSetMobileDeviceManagementAuthorityRequestBuilder) Request() *OrganizationSetMobileDeviceManagementAuthorityRequest {
	return &OrganizationSetMobileDeviceManagementAuthorityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *OrganizationSetMobileDeviceManagementAuthorityRequest) Post(ctx context.Context) (resObj *int, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

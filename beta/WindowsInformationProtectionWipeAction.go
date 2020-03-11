// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// WindowsInformationProtectionWipeAction Represents wipe requests issued by tenant admin for Bring-Your-Own-Device(BYOD) Windows devices.
type WindowsInformationProtectionWipeAction struct {
	// Entity is the base model of WindowsInformationProtectionWipeAction
	Entity
	// Status Wipe action status.
	Status *ActionState `json:"status,omitempty"`
	// TargetedUserID The UserId being targeted by this wipe action.
	TargetedUserID *string `json:"targetedUserId,omitempty"`
	// TargetedDeviceRegistrationID The DeviceRegistrationId being targeted by this wipe action.
	TargetedDeviceRegistrationID *string `json:"targetedDeviceRegistrationId,omitempty"`
	// TargetedDeviceName Targeted device name.
	TargetedDeviceName *string `json:"targetedDeviceName,omitempty"`
	// TargetedDeviceMacAddress Targeted device Mac address.
	TargetedDeviceMacAddress *string `json:"targetedDeviceMacAddress,omitempty"`
	// LastCheckInDateTime Last checkin time of the device that was targeted by this wipe action.
	LastCheckInDateTime *time.Time `json:"lastCheckInDateTime,omitempty"`
}

// WindowsInformationProtectionWipeActionRequestBuilder is request builder for WindowsInformationProtectionWipeAction
type WindowsInformationProtectionWipeActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsInformationProtectionWipeActionRequest
func (b *WindowsInformationProtectionWipeActionRequestBuilder) Request() *WindowsInformationProtectionWipeActionRequest {
	return &WindowsInformationProtectionWipeActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsInformationProtectionWipeActionRequest is request for WindowsInformationProtectionWipeAction
type WindowsInformationProtectionWipeActionRequest struct{ BaseRequest }

// Get performs GET request for WindowsInformationProtectionWipeAction
func (r *WindowsInformationProtectionWipeActionRequest) Get(ctx context.Context) (resObj *WindowsInformationProtectionWipeAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsInformationProtectionWipeAction
func (r *WindowsInformationProtectionWipeActionRequest) Update(ctx context.Context, reqObj *WindowsInformationProtectionWipeAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsInformationProtectionWipeAction
func (r *WindowsInformationProtectionWipeActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

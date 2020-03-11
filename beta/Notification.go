// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// Notification undocumented
type Notification struct {
	// Entity is the base model of Notification
	Entity
	// TargetHostName undocumented
	TargetHostName *string `json:"targetHostName,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Payload undocumented
	Payload *PayloadTypes `json:"payload,omitempty"`
	// DisplayTimeToLive undocumented
	DisplayTimeToLive *int `json:"displayTimeToLive,omitempty"`
	// Priority undocumented
	Priority *Priority `json:"priority,omitempty"`
	// GroupName undocumented
	GroupName *string `json:"groupName,omitempty"`
	// TargetPolicy undocumented
	TargetPolicy *TargetPolicyEndpoints `json:"targetPolicy,omitempty"`
}

// NotificationRequestBuilder is request builder for Notification
type NotificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns NotificationRequest
func (b *NotificationRequestBuilder) Request() *NotificationRequest {
	return &NotificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NotificationRequest is request for Notification
type NotificationRequest struct{ BaseRequest }

// Get performs GET request for Notification
func (r *NotificationRequest) Get(ctx context.Context) (resObj *Notification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Notification
func (r *NotificationRequest) Update(ctx context.Context, reqObj *Notification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Notification
func (r *NotificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
